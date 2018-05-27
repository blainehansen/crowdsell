#![feature(plugin, custom_derive)]
#![feature(specialization)]
#![plugin(rocket_codegen)]

// #[macro_use] extern crate log;

#[macro_use]
extern crate rocket;
use rocket::response::{self, Responder};
// use rocket::http::{Status, ContentType, RawStr};
use rocket::http::{Status, ContentType};
use rocket::response::Response;
// use rocket::request::{self, FromRequest, FromParam};
use rocket::request::{self, FromRequest};

use rocket::{Request, State, Outcome};

#[macro_use] extern crate dotenv_codegen;

extern crate openssl;
// use openssl::pkey::{self, PKey};
use openssl::pkey::{PKey};

extern crate base64;

extern crate yansi;

extern crate rand;
use rand::os::OsRng;
use rand::Rng;

extern crate argon2;
extern crate time;

extern crate rocket_contrib;
use rocket_contrib::Json;

extern crate serde;
use serde::Serialize;

// #[macro_use]
extern crate serde_json;

#[macro_use] extern crate serde_derive;

mod db;
use db::{DATABASE_URL, DbConn, init_pool};

#[macro_use]
extern crate diesel;
use diesel::prelude::*;
use diesel::result::Error as DieselError;

mod schema;
use schema::*;
mod models;
use models::*;

mod auth;
use auth::{AuthToken, PrivateKey, ValidAuthToken, check_token_user};

mod api_result;
use api_result::{ApiResult, ApiResponse, ApiError, ApiFailable, ApiRespondable};

#[macro_use] mod catchers;
use catchers::*;



#[get("/projects")]
fn projects(conn: DbConn) -> ApiResult<Vec<Project>> {
	projects::table
		.load::<Project>(&*conn).respond()
}



#[derive(Debug, Insertable, Serialize, Deserialize)]
#[table_name = "users"]
struct NewUser {
	name: String,
	email: String,
	password: String,
}

#[derive(Debug, Serialize)]
struct DisplayUser {
	name: String,
	email: String,
	slug: String,
}

impl From<User> for DisplayUser {
	fn from(user: User) -> Self {
		DisplayUser { ..user }
	}
}

#[derive(Debug, Serialize)]
struct SignedUser {
	token: String,
	user: DisplayUser,
}

#[post("/create-user", data = "<new_user>")]
fn create_user(new_user: Json<NewUser>, conn: DbConn, signing_key: State<PrivateKey>) -> ApiResult<SignedUser> {
	let mut new_user = new_user.into_inner();
	// hash the password
	new_user.password = auth::hash_password(new_user.password.to_owned()).or_fail()?;

	// insert into the table
	let user_ids: Vec<i32> = diesel::insert_into(users::table)
		.values(&new_user)
		.returning(users::id)
		.get_results(&*conn).or_fail()?;
	let user_id = user_ids[0];

	// sign a token and send it up
	let token_string = auth::issue_auth_token(user_id, "admin".to_string(), &signing_key).or_fail()?;
	Ok(ApiResponse::Success(SignedUser { token: token_string, user: DisplayUser { ..user } }))
}


#[derive(Debug, Serialize, Deserialize)]
struct LoginUser {
	email: String,
	password: String,
}

#[post("/login", data = "<login_user>")]
fn login(login_user: Json<LoginUser>, conn: DbConn, signing_key: State<PrivateKey>) -> ApiResult<SignedUser> {
	let login_user = login_user.into_inner();
	// find one user
	// if nothing found then return error 403 Status::Forbidden
	let user = users::table
		.filter(users::email.eq(&login_user.email))
		.first::<User>(&*conn).or_status(Status::Forbidden)?;
	let user_id = user.id;

	if auth::verify_password(user.password, login_user.password) {
		// otherwise sign a token and send it up
		let token_string = auth::issue_auth_token(user_id, "admin".to_string(), &signing_key).or_fail()?;
		Ok(ApiResponse::Success(SignedUser { token: token_string, user: DisplayUser { ..user } }))
	}
	else {
		Err(ApiError::ErrorStatus(Status::Forbidden))
	}
}


#[derive(Debug, Insertable, Deserialize)]
#[table_name = "projects"]
struct NewProject {
	title: String
}

#[post("/user/<user_id>/projects", data = "<new_project>")]
fn create_project(user_id: i32, new_project: Json<NewProject>, token: ValidAuthToken, conn: DbConn) -> ApiResult<()> {
	check_token_user(user_id, token)?;
	let new_project = new_project.into_inner();

	// use schema::projects::dsl::*;
	let project_ids: Vec<i32> = diesel::insert_into(projects::table)
		.values(&new_project)
		.returning(projects::id)
		.get_results(&*conn).or_fail()?;

	Ok(ApiResponse::SuccessStatus(Status::NoContent))
}

#[get("/user/<user_id>/secret")]
fn secret(user_id: i32, token: ValidAuthToken) -> ApiResult<&'static str> {
	check_token_user(user_id, token)?;
	Ok(ApiResponse::Success("secret"))
}


extern crate rocket_cors;
use rocket_cors::{AllowedOrigins};


fn main() {
	// println!("DATABASE URL: {}", DATABASE_URL);

	// create an hmac key
	let hmac_key: [u8; 32] = OsRng::new().unwrap().gen();
	// println!("PRIVATE KEY: {:?}", hmac_key);
	let key = PKey::hmac(&hmac_key).unwrap();

	let (allowed_origins, failed_origins) = AllowedOrigins::some(&["http://localhost:8080"]);
	assert!(failed_origins.is_empty());

	let cors_options = rocket_cors::Cors {
		allowed_origins: allowed_origins,
		// allow_credentials: true,
		max_age: Some(2592000),
		..Default::default()
	};

	rocket::ignite()
		.manage(init_pool())
		.manage(key)
		.mount("/", routes![projects, login, create_user, secret])
		.attach(cors_options)
		.catch(errors![ handle_400, handle_401, handle_402, handle_403, handle_404, handle_405, handle_406, handle_407, handle_408, handle_409, handle_410, handle_411, handle_412, handle_413, handle_414, handle_415, handle_416, handle_417, handle_418, handle_421, handle_422, handle_426, handle_428, handle_429, handle_431, handle_451, handle_500, handle_501, handle_503, handle_504, handle_510 ])
		.launch();
}


// openssl ecparam -genkey -name secp521r1 -noout -out ec512-key-pair.pem
// openssl ec -in ec512-key-pair.pem -pubout -out ecpubkey.pem
