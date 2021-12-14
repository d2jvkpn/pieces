use std::fmt;

///
#[derive(Debug, Copy, Clone)]
pub enum StatusCode {
    Ok = 200,
    BadRequest = 400,
    StatusForbidden = 403,
    NotFound = 404,
    InternalServerError = 500,
}

impl StatusCode {
    pub fn reason_phrase(&self) -> &str {
        match self {
            StatusCode::Ok => "Ok",
            StatusCode::BadRequest => "Bad Request",
            StatusCode::StatusForbidden => "Status Forbidden",
            StatusCode::NotFound => "Not Found",
            StatusCode::InternalServerError => "Internal Server Error",
        }
    }
}

impl fmt::Display for StatusCode {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", *self as u16)
    }
}
