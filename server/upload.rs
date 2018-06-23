extern crate rusoto_core;
extern crate rusoto_credential;
extern crate rusoto_s3;
use rusoto_core::Region;
use rusoto_core::reactor::RequestDispatcher;
use rusoto_s3::{S3, S3Client, PutObjectRequest};

use rusoto_credential::StaticProvider;

use std::default::Default;

lazy_static! {
	static ref SPACES_CLIENT: S3Client<StaticProvider, RequestDispatcher> = S3Client::new(
		Default::default(),
		StaticProvider::new_minimal(
			"WVTBP6TRFMDIQSRFHNQS".to_string(),
			"ulHKld8JmkSqIoebaKVXyqZQEVc9+bODk4zblJAuZT8".to_string()
		),
		Region::Custom {
			name: "nyc3".to_string(),
			endpoint: "https://nyc3.digitaloceanspaces.com".to_string(),
		}
	);
}

use std::collections::hash_map::DefaultHasher;
use std::hash::Hasher;

pub fn send_image_to_cdn(user_id: i32, data: &Vec<u8>) -> Option<String> {
	let mut hasher = DefaultHasher::new();
	hasher.write(data);
	let hash_id = hasher.finish()

	let obfuscated_user_id = "";

	let input = PutObjectRequest {
		bucket: "profile-photos".to_string(),
		body: Some(data),
		key: format!("{}-{}", obfuscated_user_id, hash_id),
		acl: Some("public-read".to_string()),
		cache_control: Some("public, max-age=31556926".to_string()),
		content_encoding: Some("gzip".to_string()),
		content_disposition: Some("inline".to_string()),
		content_type: Some("text/plain; charset=utf-8".to_string()),
		..Default::default()
	};

	// if this fails, we should probably return a 503 service unavailable
	match SPACES_CLIENT.put_object(&input).sync() {
		Ok(_) => Some(hash_id),
		Err(_) => None,
	}
}
