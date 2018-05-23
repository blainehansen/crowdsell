use openssl;
use openssl::hash::MessageDigest;
use openssl::pkey::{self, PKey};
use openssl::error::ErrorStack as OpenSslError;
// use openssl::rsa::Rsa;
// use openssl::sign::{Signer, Verifier};
use openssl::sign::{Signer};
// use openssl::ec::EcKey;

use serde_json::{self, Error as SerdeError};
use base64::{self, encode_config as b64_enc, decode_config as b64_dec, DecodeError as Base64Error};

use time::{self, Duration};

use rand::{self, Rng};
use argon2;

use rocket::request::{self, FromRequest};
use rocket::http::Status;
use rocket::{Request, State, Outcome};


use api_result::ApiError;


#[derive(Debug, Serialize, Deserialize)]
// id, expires, role
pub struct AuthToken(i32, i64, String);

// id, issuedAt, expires, role
// pub struct AuthToken(i32, i64, i64, String);

impl AuthToken {
	pub fn get_id(&self) -> i32 {
		self.0
	}

	// pub fn get_issued_at(&self) -> i64 {
	// 	self.1
	// }

	pub fn get_expires(&self) -> i64 {
		self.1
	}
}

pub type PrivateKey = PKey<pkey::Private>;


#[derive(Debug)]
pub enum AuthError {
	InvalidToken,
	ExpiredToken,
	Serde(SerdeError),
	OpenSsl(OpenSslError),
	Base64(Base64Error),
}

impl From<SerdeError> for AuthError {
	fn from(err: SerdeError) -> AuthError {
		AuthError::Serde(err)
	}
}

impl From<OpenSslError> for AuthError {
	fn from(err: OpenSslError) -> AuthError {
		AuthError::OpenSsl(err)
	}
}

impl From<Base64Error> for AuthError {
	fn from(err: Base64Error) -> AuthError {
		AuthError::Base64(err)
	}
}


type AuthResult<T> = Result<T, AuthError>;

pub fn hash_password(password: String) -> argon2::Result<String> {
	let password = password.as_bytes();
	let salt: [u8; 8] = rand::thread_rng().gen();
	argon2::hash_encoded(password, &salt, &argon2::Config::default())
}

pub fn verify_password(encoded_password: String, trial_password: String) -> bool {
	argon2::verify_encoded(&encoded_password, trial_password.as_bytes()).unwrap_or(false)
}


pub fn issue_auth_token(user_id: i32, role: String, key: &PrivateKey) -> AuthResult<String> {
	// let now = time::now_utc();
	let tomorrow = time::now_utc() + Duration::hours(24);

	// create a tuple to serialize
	// let token = AuthToken(user_id, now.to_timespec().sec, tomorrow.to_timespec().sec, role);
	let token = AuthToken(user_id, tomorrow.to_timespec().sec, role);

	// serialize that tuple
	// base64 encode it
	// serde_json::to_vec
	let json_payload = &serde_json::to_string(&token)?;
	let encoded_payload = b64_enc(json_payload, base64::URL_SAFE);

	let mut signer = Signer::new(MessageDigest::sha256(), key)?;
	signer.update(encoded_payload.as_bytes())?;
	let signature = signer.sign_to_vec()?;

	// base64 encode the signature
	let encoded_signature = b64_enc(signature.as_slice(), base64::URL_SAFE);

	// return it with a dot separating it from it's signature
	Ok(format!("{}.{}", encoded_payload, encoded_signature))
}

pub fn verify_auth_token(token: String, key: &PrivateKey) -> AuthResult<AuthToken> {
	// then split a token by dot
	let segments: Vec<&str> = token.split(".").collect();
	if segments.len() != 2 {
		return Err(AuthError::InvalidToken);
	}

	let proposed_encoded_payload = segments[0];
	let proposed_encoded_signature = segments[1];
	// println!("{:?}", proposed_encoded_payload);
	// println!("{:?}", proposed_encoded_signature);

	// base64 decode the payload and signature
	// we actually should only bother to decode the payload once we've verified it matches
	let proposed_signature = b64_dec(proposed_encoded_signature.as_bytes(), base64::URL_SAFE)?;
	// println!("{:?}", proposed_signature);

	// sign it
	let mut signer = Signer::new(MessageDigest::sha256(), key)?;
	signer.update(proposed_encoded_payload.as_bytes())?;
	let actual_signature_of_proposed = signer.sign_to_vec()?;
	// println!("{:?}", actual_signature_of_proposed);

	// compare the signature with the one you just created
	let valid = openssl::memcmp::eq(&actual_signature_of_proposed, &proposed_signature);
	// println!("{:?}", valid);

	let proposed_payload = b64_dec(proposed_encoded_payload.as_bytes(), base64::URL_SAFE)?;
	// println!("{:?}", proposed_payload);

	// get a valid struct
	let successful_token: AuthToken = serde_json::from_slice(&proposed_payload)?;
	// println!("{:?}", successful_token);

	// TODO
	// now that we know the signature is valid (we created this token)
	// we should check if the times are still within range
	// we check that:
	// the issued at isn't greater than now
	// the expires
	let valid_expiration = (time::now_utc() + Duration::minutes(5)).to_timespec().sec;

	// the token is no longer valid if the expiration time is before now
	// so the time is less than now
	if successful_token.get_expires() < valid_expiration {
		Err(AuthError::ExpiredToken)
	}
	else {
		Ok(successful_token)
	}
}


#[derive(Debug)]
pub struct ValidAuthToken(AuthToken);

impl<'a, 'r> FromRequest<'a, 'r> for ValidAuthToken {
	type Error = ();

	fn from_request(request: &'a Request<'r>) -> request::Outcome<ValidAuthToken, ()> {
		let auth_header = match request.headers().get_one("Authorization") {
			Some(string) => string,
			None => return Outcome::Failure((Status::Unauthorized, ())),
		};

		// parse it to remove the bearer portion
		// I'm bothering to do this just to have a basic idea that at least it's in the right neighborhood
		let segments: Vec<&str> = auth_header.splitn(2, ' ').collect();
		if segments.len() != 2 && segments[0] != "Bearer" {
			return Outcome::Failure((Status::BadRequest, ()));
		}

		let token = segments[1];
		let signing_key = match request.guard::<State<PrivateKey>>().succeeded() {
			Some(key) => key,
			None => return Outcome::Failure((Status::InternalServerError, ())),
		};

		let token = match verify_auth_token(token.to_string(), &signing_key).ok() {
			Some(token) => token,
			None => return Outcome::Failure((Status::Forbidden, ())),
		};

		Outcome::Success(ValidAuthToken(token))
	}
}

// // UserId will call the ValidAuthToken guard to make sure the requester is authorized to make changes for this id
// #[derive(Debug)]
// struct UserId(i32);

// trait StrictOutcomeAble<S> {
// 	fn success_or_error(self) -> Result<S, Status>;
// }

// impl<S, E, F> StrictOutcomeAble<S> for request::Outcome<S, E, F> {
// 	fn success_or_error(self, failure_handler: FnOnce() -> OE, forward_error: Option<FE>) -> Result<S, Status> {
// 		match self {
// 			Outcome::Success(success_value) => Ok(success_value),
// 			Outcome::Forward(_) => Err(Status::InternalServerError),
// 			Outcome::Failure((status, _)) => Err(status),
// 		}
// 	}
// }


// #[derive(Debug)]
// struct ValidUserId;

// impl<'a, 'r> FromRequest<'a, 'r> for ValidUserId {
// 	type Error = ();

// 	fn from_request(request: &'a Request<'r>) -> request::Outcome<ValidUserId, ()> {
// 		let user_id = request.get_param::<i32>(0)
// 			.or(Err((Status::BadRequest, ())))?
// 			// .map_err(|e| {
// 			// 	match e {
// 			// 		BadParse => ,
// 			// 		NoKey =>,
// 			// 	}
// 			// })

// 		let token = match request.guard::<ValidAuthToken>() {
// 			Outcome::Success(token) => Ok(token),
// 			Outcome::Forward(_) => Err(Status::InternalServerError),
// 			Outcome::Failure((status, _)) => Err(status),
// 		}?;

// 		if id != token.get_id() {
// 			Err(Status::Forbidden)
// 		}
// 		else {
// 			Ok(UserId(id))
// 		}
// 	}
// }

pub fn check_token_user(user_id: i32, token: ValidAuthToken) -> Result<(), ApiError> {
	if user_id == token.0.get_id() {
		Ok(())
	}
	else {
		Err(ApiError::ErrorStatus(Status::Forbidden))
	}
}
