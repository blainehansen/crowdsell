use rocket::http::{Status, StatusClass};
use rocket::response::{self, Response};
use rocket::Request;
use rocket::error::Error;
use yansi::Paint;

pub fn log_status(status: Status) {
	let status_string = status.to_string();
	let painted_code = match status.class() {
		StatusClass::Informational => Paint::white(status_string),
		StatusClass::Success => Paint::green(status_string),
		StatusClass::Redirection => Paint::blue(status_string),
		StatusClass::ClientError => Paint::yellow(status_string),
		StatusClass::ServerError => Paint::red(status_string),
		StatusClass::Unknown => Paint::purple(status_string),
	};

	println!("    => {}", painted_code);
}

fn catch_with_code<'r>(code: u16) -> response::Result<'r> {
	let status = Status::from_code(code).unwrap();
	log_status(status);
	Response::build()
		.status(status)
		.ok()
}


#[error(400)]
pub fn handle_400<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(400) }

#[error(401)]
pub fn handle_401<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(401) }

#[error(402)]
pub fn handle_402<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(402) }

#[error(403)]
pub fn handle_403<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(403) }

#[error(404)]
pub fn handle_404<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(404) }

#[error(405)]
pub fn handle_405<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(405) }

#[error(406)]
pub fn handle_406<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(406) }

#[error(407)]
pub fn handle_407<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(407) }

#[error(408)]
pub fn handle_408<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(408) }

#[error(409)]
pub fn handle_409<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(409) }

#[error(410)]
pub fn handle_410<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(410) }

#[error(411)]
pub fn handle_411<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(411) }

#[error(412)]
pub fn handle_412<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(412) }

#[error(413)]
pub fn handle_413<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(413) }

#[error(414)]
pub fn handle_414<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(414) }

#[error(415)]
pub fn handle_415<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(415) }

#[error(416)]
pub fn handle_416<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(416) }

#[error(417)]
pub fn handle_417<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(417) }

#[error(418)]
pub fn handle_418<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(418) }

#[error(421)]
pub fn handle_421<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(421) }

#[error(422)]
pub fn handle_422<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(422) }

#[error(426)]
pub fn handle_426<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(426) }

#[error(428)]
pub fn handle_428<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(428) }

#[error(429)]
pub fn handle_429<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(429) }

#[error(431)]
pub fn handle_431<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(431) }

#[error(451)]
pub fn handle_451<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(451) }

#[error(500)]
pub fn handle_500<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(500) }

#[error(501)]
pub fn handle_501<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(501) }

#[error(503)]
pub fn handle_503<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(503) }

#[error(504)]
pub fn handle_504<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(504) }

#[error(510)]
pub fn handle_510<'r>(_: Error, _: &'r Request) -> response::Result<'r> { catch_with_code(510) }


// macro_rules! create_catchers {
// 	($($code:expr, $fn_name:ident),+) => (

// 		$(
// 			#[error($code)]
// 			pub fn $fn_name<'r>(_: Error, req: &'r Request) -> response::Result<'r> {
// 				let status = Status::from_code($code).unwrap();
// 				log_status(status);
// 				Response::build()
// 					.status(status)
// 					.ok()
// 			}
// 		)+
// 	)
// }

// create_catchers! {
// 	400, handle_400,
// 	401, handle_401,
// 	402, handle_402,
// 	403, handle_403,
// 	404, handle_404,
// 	405, handle_405,
// 	406, handle_406,
// 	407, handle_407,
// 	408, handle_408,
// 	409, handle_409,
// 	410, handle_410,
// 	411, handle_411,
// 	412, handle_412,
// 	413, handle_413,
// 	414, handle_414,
// 	415, handle_415,
// 	416, handle_416,
// 	417, handle_417,
// 	418, handle_418,
// 	421, handle_421,
// 	422, handle_422,
// 	426, handle_426,
// 	428, handle_428,
// 	429, handle_429,
// 	431, handle_431,
// 	451, handle_451,
// 	500, handle_500,
// 	501, handle_501,
// 	503, handle_503,
// 	504, handle_504,
// 	510, handle_510
// }

// #[macro_export]
// macro_rules! output_catchers {
// 	() => (handle_400, handle_401, handle_402, handle_403, handle_404, handle_405, handle_406, handle_407, handle_408, handle_409, handle_410, handle_411, handle_412, handle_413, handle_414, handle_415, handle_416, handle_417, handle_418, handle_421, handle_422, handle_426, handle_428, handle_429, handle_431, handle_451, handle_500, handle_501, handle_503, handle_504, handle_510)
// }
