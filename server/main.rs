#![feature(plugin)]
#![plugin(rocket_codegen)]

extern crate rocket;

use diesel::sqlite::SqliteConnection;
use diesel::r2d2::{ConnectionManager, Pool, PooledConnection};

// An alias to the type for a pool of Diesel SQLite connections.
type SqlitePool = Pool<ConnectionManager<SqliteConnection>>;

// The URL to the database, set via the `DATABASE_URL` environment variable.
static DATABASE_URL: &'static str = env!("DATABASE_URL");

/// Initializes a database pool.
fn init_pool() -> Pool {
    let manager = ConnectionManager::<SqliteConnection>::new(DATABASE_URL);
    Pool::new(manager).expect("db pool")
}


#[get("/")]
fn index() -> &'static str {
	"Hello, world!"
}

use diesel::r2d2::{ConnectionManager, Pool, PooledConnection};

#[get("/projects")]
fn projects(conn: DbConn) -> Vec<Project> {
	projects::table.load::<Project>(&*conn)
}

fn main() {
	rocket::ignite()
		.manage(init_pool())
		.mount("/", routes![index, projects])
		.launch();
}
