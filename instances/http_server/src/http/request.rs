use std::convert::TryFrom;
// use std::error::Error;
use std::fmt;
use std::str::{self, Utf8Error};

use super::method::{Method, MethodError};
use super::query_string::QueryString;

pub enum ParseError {
    InvalidRequest,
    InvalidEncoding,
    InvalidProtocol,
    InvalidMethod,
}

impl ParseError {
    fn message(&self) -> &str {
        match self {
            Self::InvalidRequest => "Invalid Request",
            Self::InvalidEncoding => "Invalid Encoding",
            Self::InvalidProtocol => "Invalid Protocol",
            Self::InvalidMethod => "Invalid Method",
        }
    }
}

impl From<Utf8Error> for ParseError {
    fn from(_: Utf8Error) -> Self {
        Self::InvalidEncoding
    }
}

////
impl From<MethodError> for ParseError {
    fn from(_: MethodError) -> Self {
        Self::InvalidMethod
    }
}

// impl Error for ParseError {}
impl fmt::Display for ParseError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.message())
    }
}

impl fmt::Debug for ParseError {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "{}", self.message())
    }
}

////
//#[derive(Debug)]
//pub struct Request {
//    pub method: Method,
//    pub path: String,
//    pub query_string: Option<String>,
//}

#[derive(Debug)]
pub struct Request<'b> {
    pub method: Method,
    pub path: &'b str,
    pub query_string: Option<QueryString<'b>>,
}

impl Request<'_> {
    fn from_byte_array(buf: &[u8]) -> Result<Self, String> {
        unimplemented!()
    }

    pub fn path(&self) -> &str {
        &self.path
    }

    pub fn method(&self) -> &Method {
        &self.method
    }

    pub fn query_string(&self) -> Option<&QueryString> {
        self.query_string.as_ref()
    }
}

impl<'b> TryFrom<&'b [u8]> for Request<'b> {
    type Error = ParseError;

    // GET /search?name=abc&ort=1 HTTP1.1\r\n...HEADERS...
    // fn try_from(buf: &'b [u8]) -> Result<Request<'b>, Self::Error> {
    fn try_from(buf: &[u8]) -> Result<Request, Self::Error> {
        //		let string = String::from("asd");
        //		string.encrypt();
        //		buf.encrypt();

        //        match String::from_utf8(buf[..]).or(Err(ParseError::InvalidEncoding)) {
        //            Ok(v) => {}
        //            Err(e) => return Err(e),
        //        }

        // let req = String::from_utf8(buf[..]).or(Err(ParseError::InvalidEncoding))?;

        // as impl From<Utf8Error> for ParseError {}
        let text = str::from_utf8(buf)?;

        let (method, text) = get_next_word(text).ok_or(ParseError::InvalidRequest)?;
        let method: Method = method.parse()?; // impl From<MethodError> for ParseError

        let (path, text) = get_next_word(text).ok_or(ParseError::InvalidRequest)?;

        let (path, query_string) = match path.find('?') {
            Some(i) => {
                // println!("~~~ {}, {}", i, path);
                (&path[..i], Some(QueryString::from(&path[i + 1..])))
            }
            None => (path, None),
        };
        //        let mut query_string = None;
        //        if let Some(i) = path.find('?') {
        //        	query_string = Some((&path[i + 1..]).to_string());
        //        	path = &path[..i];
        //        }

        let (protocol, _) = get_next_word(text).ok_or(ParseError::InvalidRequest)?;

        if protocol != "HTTP/1.1" {
            return Err(ParseError::InvalidProtocol);
        }

        //        Ok(Request {
        //            method,
        //            path: path.to_string(),
        //            query_string,
        //        })

        Ok(Request { method, path: path, query_string })
        // unimplemented!()
    }
}

fn get_next_word(request: &str) -> Option<(&str, &str)> {
    //    let mut iter = request.chars();
    //    loop {
    //        let item = iter.next();
    //        match item {
    //            Some(v) => {}
    //            None => break,
    //        }
    //    }
    //
    for (i, v) in request.chars().enumerate() {
        if v == ' ' || v == '\r' {
            return Some((&request[..i], &request[i + 1..]));
        }
    }

    None
    // unimplemented!()
}

////
trait Encrypt {
    fn encrypt(&self) -> Self;
}

impl Encrypt for String {
    fn encrypt(&self) -> Self {
        unimplemented!()
    }
}

impl Encrypt for &[u8] {
    fn encrypt(&self) -> Self {
        unimplemented!()
    }
}
