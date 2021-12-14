use std::fmt;
use std::io;

use super::status_code::StatusCode;

////
#[derive(Debug)]
pub struct Response {
    status_code: StatusCode,
    body: Option<String>,
}

impl Response {
    pub fn new(status_code: StatusCode, body: Option<String>) -> Self {
        Response { status_code, body }
    }

    pub fn send(&self, w: &mut dyn io::Write) -> io::Result<()> {
        let body = match &self.body {
            Some(v) => v,
            None => "",
        };

        write!(w, "{}\r\n\r\n{}", self, body)
    }
}

impl fmt::Display for Response {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "HTTP/1.1 {} {}", self.status_code, self.status_code.reason_phrase(),)
    }
}
