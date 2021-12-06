use std::fs;

use super::server::Handler;
use crate::http::{Request, Response, StatusCode};

pub struct WebsiteHandler {
    public_path: String,
}

impl WebsiteHandler {
    pub fn new(public_path: String) -> Self {
        Self { public_path }
    }

    pub fn read_file(&self, file_path: &str) -> Option<String> {
        let path = match file_path.strip_prefix("/static/") {
            // !!! vulnerable os file path, like ../../../
            Some(v) => format!("{}/{}", self.public_path, v),
            None => return None,
        };
        dbg!(&path);

        fs::read_to_string(path).ok()
    }
}

impl Handler for WebsiteHandler {
    fn handle_request(&mut self, request: &Request) -> Response {
        //        println!(
        //            "~~~ Request.path: {}, Request.path(): {}",
        //            request.path,
        //            request.path()
        //        );

        match request.path() {
            "/" => Response::new(StatusCode::Ok, Some("<h4>Welcome</h4>".to_string())),
            "/hello" => Response::new(StatusCode::Ok, Some("<h4>Hello</h4>".to_string())),
            "/ping" => Response::new(StatusCode::Ok, Some("<h4>pong</h4>".to_string())),
            p if p.starts_with("/static/") => match self.read_file(p) {
                Some(v) => Response::new(StatusCode::Ok, Some(v.to_string())),
                None => Response::new(StatusCode::NotFound, Some("file not found".to_string())),
            },
            _ => Response::new(StatusCode::NotFound, None),
        }
    }
}
