use std::path::PathBuf;
use std::{fs, io};

use super::server::Handler;
use crate::http::{Request, Response, StatusCode};

pub struct SimpleHandler {
    public_path: String,
}

impl SimpleHandler {
    pub fn new(public_path: &str) -> Result<SimpleHandler, String> {
        let pb = PathBuf::from(&public_path);
        // dbg!(&pb);

        let pb = match fs::canonicalize(&pb) {
            Ok(v) => v,
            Err(e) => return Err(format!("fs::canonicalize -> {}", e)),
        };

        //        let pb = match pb.into_os_string().into_string() {
        //            Ok(v) => v,
        //            Err(e) => return Err(format!("pb into_string -> {:?}", e)),
        //        };

        let pb = pb.display().to_string();

        println!("public_path: {}", pb);
        Ok(SimpleHandler { public_path: pb })
    }

    pub fn read_file(&self, file_path: &str) -> Option<String> {
        let pb = PathBuf::from(&self.public_path);

        let path = match file_path.strip_prefix("/static/") {
            // !!! vulnerable os file path, like ../../../
            Some(v) => pb.join(PathBuf::from(v)), // maybe /static/../Cargo.toml
            None => return None,
        };

        // !! fs::canonicalize https://doc.rust-lang.org/std/fs/fn.canonicalize.html
        // return an error of file not exists, not generate full path like realpath or readlink -f
        //        let path = match fs::canonicalize(&path) {
        //            Ok(v) => v,
        //            Err(e) => return None,
        //        };

        let path = fs::canonicalize(&path).ok()?;
        // dbg!(&path);
        // dbg!(&pb);

        if !path.starts_with(pb) {
            eprintln!(
                "!!! vulnerable os file path: {}",
                path.display().to_string()
            );
            return None; // should return http.StatusForbidden 403
        }

        fs::read_to_string(path.display().to_string()).ok()
    }
}

impl Handler for SimpleHandler {
    fn handle_request(&mut self, request: &Request) -> Response {
        //        println!(
        //            "~~~ Request.path: {}, Request.path(): {}",
        //            request.path,
        //            request.path()
        //        );
        // dbg!(request.path());

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
