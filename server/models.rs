#[derive(Serialize, Deserialize, Queryable)]
pub struct Project {
	pub id: i32,
	pub title: Option<String>,
}


#[derive(Debug, Serialize, Deserialize, Queryable)]
pub struct User {
	pub id: i32,
	pub name: String,
	pub email: String,
	pub password: String,
}
