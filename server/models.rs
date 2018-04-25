// use diesel;
// use diesel::prelude::*;
// use schema::projects;

#[derive(Serialize, Deserialize, Queryable)]
pub struct Project {
	pub id: i32,
	pub title: Option<String>,
}
