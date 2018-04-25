#![feature(plugin)]
#![plugin(rocket_codegen)]


extern crate rocket;
// #[macro_use]
extern crate rocket_contrib;

extern crate serde;
#[macro_use]
extern crate serde_json;
#[macro_use] extern crate serde_derive;

#[macro_use] extern crate diesel;
#[macro_use] extern crate dotenv_codegen;

extern crate frank_jwt;
use frank_jwt::{Algorithm, encode, decode};
extern crate time;

use std::env;



use diesel::pg::PgConnection;
use diesel::r2d2::{ConnectionManager, Pool, PooledConnection};

// An alias to the type for a pool of Diesel Postgres connections.
type PostgresPool = Pool<ConnectionManager<PgConnection>>;
// Connection request guard type: a wrapper around an r2d2 pooled connection.
struct DbConn(pub PooledConnection<ConnectionManager<PgConnection>>);

// The URL to the database, set via the `DATABASE_URL` environment variable.
static DATABASE_URL: &'static str = dotenv!("DATABASE_URL");

/// Initializes a database pool.
fn init_pool() -> PostgresPool {
	let manager = ConnectionManager::<PgConnection>::new(DATABASE_URL);
	PostgresPool::new(manager).expect("db pool")
}



use std::ops::Deref;
use rocket::http::Status;
use rocket::request::{self, FromRequest};
use rocket::{Request, State, Outcome};

/// Attempts to retrieve a single connection from the managed database pool. If
/// no pool is currently managed, fails with an `InternalServerError` status. If
/// no connections are available, fails with a `ServiceUnavailable` status.
impl<'a, 'r> FromRequest<'a, 'r> for DbConn {
	type Error = ();

	fn from_request(request: &'a Request<'r>) -> request::Outcome<Self, Self::Error> {
		let pool = request.guard::<State<PostgresPool>>()?;
		match pool.get() {
			Ok(conn) => Outcome::Success(DbConn(conn)),
			Err(_) => Outcome::Failure((Status::ServiceUnavailable, ()))
		}
	}
}

// For the convenience of using an &DbConn as an &PgConnection.
impl Deref for DbConn {
	type Target = PgConnection;

	fn deref(&self) -> &Self::Target {
		&self.0
	}
}


mod schema;
mod models;
use diesel::prelude::*;
use diesel::result::Error as DieselError;
use models::*;
use schema::*;
use rocket_contrib::Json as RocketJson;
use serde::Serialize;
use rocket::response::{self, Responder};

use std::error;
use std::fmt;


#[derive(Debug)]
struct QueryError(DieselError);

impl fmt::Display for QueryError {
	fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
		write!(f, "{}", self.0)
	}
}
impl error::Error for QueryError {
	fn description(&self) -> &str {
		match self.0 {
			DieselError::NotFound => "NotFound",
			DieselError::InvalidCString(_) => "",
			DieselError::DatabaseError(_, _) => "DatabaseError",
			DieselError::QueryBuilderError(_) => "QueryBuilderError",
			DieselError::DeserializationError(_) => "DeserializationError",
			DieselError::SerializationError(_) => "SerializationError",
			DieselError::RollbackTransaction => "RollbackTransaction",
			DieselError::AlreadyInTransaction => "AlreadyInTransaction",
			_ => "Unknown DieselError",
		}
	}
}

impl<'r> Responder<'r> for QueryError {
	fn respond_to(self, _: &Request) -> response::Result<'r> {
		match self.0 {
			DieselError::NotFound => Err(Status::NotFound),
			_ => Err(Status::InternalServerError),
		}
	}
}


type QueryResponse<T> = Result<RocketJson<T>, QueryError>;

trait QueryRespondable<T: Serialize> {
	fn respond(self) -> QueryResponse<T>;
}

impl<T: Serialize> QueryRespondable<T> for QueryResult<T> {
	fn respond(self) -> QueryResponse<T> {
		match self {
			Ok(value) => Ok(RocketJson(value)),
			Err(err) => Err(QueryError(err))
		}
	}
}


#[get("/projects")]
fn projects(conn: DbConn) -> QueryResponse<Vec<Project>> {
	projects::table
		.load::<Project>(&*conn).respond()
}


fn main() {
	println!("{}", DATABASE_URL);

	// rocket::ignite()
	// 	.manage(init_pool())
	// 	.mount("/", routes![projects])
	// 	.launch();


	use time::{self, Duration};
	let now = time::now_utc();
	let tomorrow = now + Duration::hours(24);

	let mut payload = json!({
		"iss": "crowd-sell",
		"sub": "1",
		"iat": now.to_timespec().sec,
		"nbf": now.to_timespec().sec,
		"exp": tomorrow.to_timespec().sec,
	});
	println!("{:?}", payload);
	let mut header = json!({
	});
	println!("{:?}", header);

	let mut private_keypath = env::current_dir().unwrap();
	private_keypath.push("ecprivate.pem");

	let jwt = encode(header, &private_keypath.to_path_buf(), &payload, Algorithm::ES512).unwrap();
	println!("{:?}", jwt);

	let mut public_keypath = env::current_dir().unwrap();
	public_keypath.push("ecpublic.pem");
	let (header, payload) = decode(&jwt, &public_keypath.to_path_buf(), Algorithm::ES512).unwrap();
	println!("{:?}", header);
	println!("{:?}", payload);
}


// openssl ecparam -genkey -name secp521r1 -noout -out ec512-key-pair.pem
// openssl ec -in ec512-key-pair.pem -pubout -out ecpubkey.pem
