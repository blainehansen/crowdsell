use rocket::response::{self, Responder};
use rocket::response::Response;
use rocket::http::{Status, ContentType};
use rocket::http::hyper::header as Headers;
use rocket::{Request};

use diesel::result::Error as DieselError;

use serde::Serialize;

use std;
use std::io::Cursor;

extern crate serde_json;

use yansi::Paint;

use catchers::log_status;

// // use std::hash::{Hash, Hasher};
// use std::collections::hash_map::DefaultHasher;
// use std::hash::{Hash};

// trait ApiValue: Serialize + Hash {}
trait ApiValue: Serialize {}

#[derive(Debug)]
pub enum ApiResponse<T: ApiValue> {
	Success(T),
	SuccessStatus(Status),
	SuccessStatusContent(Status, T),
}

// fn calculate_hash<T: Hash>(t: &T) -> u64 {
// 	let mut s = DefaultHasher::new();
// 	t.hash(&mut s);
// 	s.finish()
// }

pub fn build_json_response<'r, T: ApiValue>(value: T, option_status: Option<Status>) -> response::Result<'r> {
	let response_string = match serde_json::to_string(&value) {
		Err(e) => {
			println!("    => JSON failed to serialize: {:?}", Paint::red(e));
			return Response::build().status(Status::InternalServerError).ok();
		},
		Ok(value) => value,
	};

	// use Headers::{CacheControl, CacheDirective, EntityTag, ETag};

	let res = Response::build()
		.header(ContentType::JSON)
		// .header(CacheControl(vec![CacheDirective::MaxAge(86400u32)]))
		// .header(ETag(EntityTag::new(false, value.to_string())))
		.sized_body(Cursor::new(response_string))
		.finalize();

	if let Some(status) = option_status {
		return Response::build_from(res).status(status).ok();
	}
	Ok(res)
}

impl<'r, T: ApiValue> Responder<'r> for ApiResponse<T> {
	fn respond_to(self, _: &Request) -> response::Result<'r> {
		let (status, response) = match self {
			ApiResponse::Success(value) => (Status::Ok, build_json_response(value, None)),
			ApiResponse::SuccessStatus(status) => (status, Response::build().status(status).ok()),
			ApiResponse::SuccessStatusContent(status, value) => (status, build_json_response(value, Some(status))),
		};

		log_status(status);
		response
	}
}

#[derive(Debug)]
pub enum ApiError {
	ErrorStatus(Status),
	ErrorStatusMessage(Status, String),
	UnknownError,
}

impl std::fmt::Display for ApiError {
	fn fmt(&self, f: &mut std::fmt::Formatter) -> std::fmt::Result {
		write!(f, "{:?}", self)
	}
}
impl std::error::Error for ApiError {
	fn description(&self) -> &str {
		"ApiError"
	}
}

impl<'r> Responder<'r> for ApiError {
	fn respond_to(self, _: &Request) -> response::Result<'r> {
		let (status, response) = match self {
			ApiError::ErrorStatus(status) => (status, Response::build().status(status).ok()),
			ApiError::ErrorStatusMessage(status, value) => (status, build_json_response(value, Some(status))),
			ApiError::UnknownError => (Status::InternalServerError, Response::build().status(Status::InternalServerError).ok())
		};

		log_status(status);
		response
	}
}

fn print_unknown_error<E: std::fmt::Debug>(e: E) {
	println!("    => UnknownError: {:?}", Paint::red(e));
}


impl From<DieselError> for ApiError {
	fn from(diesel_error: DieselError) -> Self {
		match diesel_error {
			// DatabaseError, could be result of constraint error
			DieselError::NotFound => ApiError::ErrorStatus(Status::NotFound),
			unknown_diesel_error => {
				print_unknown_error(unknown_diesel_error);
				ApiError::UnknownError
			},
		}
	}
}


pub type ApiResult<T> = Result<ApiResponse<T>, ApiError>;
// pub type MessageApiResult<T> = Result<ApiResponse<T>, ApiError<String>>;



// this trait is all about early termination
pub trait ApiFailable<T> {
	fn or_fail(self) -> Result<T, ApiError>;

	fn or_status(self, status: Status) -> Result<T, ApiError>;
}

impl<T, E: std::fmt::Debug> ApiFailable<T> for Result<T, E> {
	fn or_fail(self) -> Result<T, ApiError> {
		self.map_err(|unknown_error| {
			print_unknown_error(unknown_error);
			ApiError::UnknownError
		})
	}

	fn or_status(self, status: Status) -> Result<T, ApiError> {
		self.or(Err(ApiError::ErrorStatus(status)))
	}
}

impl<T> ApiFailable<T> for Option<T> {
	fn or_fail(self) -> Result<T, ApiError> {
		self.ok_or(ApiError::ErrorStatus(Status::InternalServerError))
	}

	fn or_status(self, status: Status) -> Result<T, ApiError> {
		self.ok_or(ApiError::ErrorStatus(status))
	}
}


pub trait ApiRespondable<T: ApiValue> {
	fn respond(self) -> ApiResult<T>;
}

impl<T: ApiValue, E: std::fmt::Debug> ApiRespondable<T> for Result<T, E> {
	default fn respond(self) -> ApiResult<T> {
		match self {
			Ok(value) => Ok(ApiResponse::Success(value)),
			Err(unknown_error) => {
				print_unknown_error(unknown_error);
				Err(ApiError::UnknownError)
			},
		}
	}
}

impl<T: ApiValue, E: std::fmt::Debug + Into<ApiError>> ApiRespondable<T> for Result<T, E> {
	fn respond(self) -> ApiResult<T> {
		match self {
			Ok(value) => Ok(ApiResponse::Success(value)),
			Err(discernable_error) => Err(discernable_error.into())
		}
	}
}


impl<T: ApiValue> ApiRespondable<T> for Option<T> {
	fn respond(self) -> ApiResult<T> {
		match self {
			Some(value) => Ok(ApiResponse::Success(value)),
			None => Err(ApiError::ErrorStatus(Status::NotFound)),
		}
	}
}
