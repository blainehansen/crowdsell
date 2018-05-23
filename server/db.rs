use diesel::pg::PgConnection;
use diesel::r2d2::{ConnectionManager, Pool, PooledConnection};

// The URL to the database, set via the `DATABASE_URL` environment variable.
pub static DATABASE_URL: &'static str = "postgres://user:asdf@localhost/crowdsell";

// An alias to the type for a pool of Diesel Postgres connections.
type PostgresPool = Pool<ConnectionManager<PgConnection>>;

/// Initializes a database pool.
pub fn init_pool() -> PostgresPool {
	let manager = ConnectionManager::<PgConnection>::new(DATABASE_URL);
	PostgresPool::new(manager).expect("db pool")
}


use std::ops::Deref;

use rocket::http::Status;
use rocket::request::{self, FromRequest};
use rocket::{Request, State, Outcome};

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

// For the convenience of using a &DbConn as a &PgConnection.
impl Deref for DbConn {
	type Target = PgConnection;

	fn deref(&self) -> &Self::Target {
		&self.0
	}
}
