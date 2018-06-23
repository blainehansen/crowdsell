#[derive(Serialize, Deserialize, Queryable)]
pub struct Project {
	pub id: i32,
	pub title: Option<String>,
}


#[derive(Debug, Serialize, Deserialize, Queryable)]
pub struct User {
	pub id: i32,
	pub hash_id: String,
	pub profile_photo_hash: String,
	pub name: String,
	pub email: String,
	pub slug: String,
	pub password: String,
}
