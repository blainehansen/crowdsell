#![feature(plugin)]
#![plugin(rocket_codegen)]

extern crate rocket;
#[macro_use] extern crate diesel;

use std::ops::Deref;
use rocket::http::Status;
use rocket::request::{self, FromRequest};
use rocket::{Request, State, Outcome};

use diesel::pg::PgConnection;
use diesel::r2d2::{ConnectionManager, Pool, PooledConnection};

// An alias to the type for a pool of Diesel Postgres connections.
type PostgresPool = Pool<ConnectionManager<PgConnection>>;
// Connection request guard type: a wrapper around an r2d2 pooled connection.
pub struct DbConn(pub PooledConnection<ConnectionManager<PgConnection>>);

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

pub mod schema;
pub mod models;


// The URL to the database, set via the `DATABASE_URL` environment variable.
static DATABASE_URL: &'static str = env!("DATABASE_URL");

/// Initializes a database pool.
fn init_pool() -> Pool {
	let manager = ConnectionManager::<PgConnection>::new(DATABASE_URL);
	Pool::new(manager).expect("db pool")
}



#[get("/projects")]
fn projects(conn: DbConn) -> Vec<Project> {
	projects::table.load::<Project>(&*conn)
}

fn main() {
	rocket::ignite()
		.manage(init_pool())
		.mount("/", routes![projects])
		.launch();
}
