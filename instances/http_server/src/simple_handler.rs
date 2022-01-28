use std::{fs, io, path::PathBuf};

use super::server::Handler;
use crate::http::{method::Method, Request, Response, StatusCode};

use path_clean::PathClean;

pub struct SimpleHandler {
    public_path: String,
}

impl SimpleHandler {
    pub fn new(public_path: &str) -> Result<SimpleHandler, String> {
        let pb = PathBuf::from(&public_path);
        // dbg!(&pb);

        let pb = fs::canonicalize(&pb).map_err(|e| format!("fs::canonicalize -> {}", e))?;

        //        let pb = match pb.into_os_string().into_string() {
        //            Ok(v) => v,
        //            Err(e) => return Err(format!("pb into_string -> {:?}", e)),
        //        };

        let pb = pb.display().to_string();

        println!("public_path: {}", pb);
        Ok(SimpleHandler { public_path: pb })
    }

    pub fn read_file(&self, file_path: &str) -> Response {
        let pb = PathBuf::from(&self.public_path);

        let path = match file_path.strip_prefix("/static/") {
            // !!! vulnerable os file path, like ../../../
            Some(v) => pb.join(PathBuf::from(v)).clean(), // maybe /static/../Cargo.toml
            None => {
                return Response::new(StatusCode::NotFound, Some("invalid file path".to_string()))
            }
        };

        // !! fs::canonicalize https://doc.rust-lang.org/std/fs/fn.canonicalize.html
        // return an error of file not exists, not generate full path like realpath or readlink -f
        //        let path = match fs::canonicalize(&path) {
        //            Ok(v) => v,
        //            Err(e) => return None,
        //        };

        // let path = fs::canonicalize(&path).ok()?; // Option<String>

        //        let path = match fs::canonicalize(&path) {
        //            Ok(v) => v,
        //            Err(_) => {
        //                return Response::new(StatusCode::NotFound, Some("file not exits".to_string()))
        //            }
        //        };

        // dbg!(&path);
        // dbg!(&pb);

        if !path.starts_with(pb) {
            eprintln!("!!! vulnerable os file path: {}", path.display().to_string());

            return Response::new(
                StatusCode::StatusForbidden,
                Some("invalid file path".to_string()),
            );
        }

        //        if !path.exists() {
        //            return Response::new(StatusCode::NotFound, Some("file not exits".to_string()));
        //        }

        // fs::read_to_string(path.display().to_string()).ok()
        let err = match fs::read_to_string(path.display().to_string()) {
            Ok(v) => return Response::new(StatusCode::Ok, Some(v.to_string())),
            Err(e) => e,
        };
        dbg!(&err);

        // !! not working
        match err.kind() {
            io::ErrorKind::IsADirectory => {
                Response::new(StatusCode::NotFound, Some("target is a directory".to_string()))
            }
            io::ErrorKind::NotFound => {
                Response::new(StatusCode::NotFound, Some("file not exists".to_string()))
            }
            _ => Response::new(StatusCode::InternalServerError, None),
        }

        //        let kind = format!("{:?}", err.kind());
        //        match &kind[..] {
        //            "IsADirectory" => {
        //                Response::new(StatusCode::NotFound, Some("target is a directory".to_string()))
        //            }
        //            _ => Response::new(StatusCode::InternalServerError, None),
        //        }

        // let Some(e) = err.downcast_ref::<std::io::Error>()
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

        if request.method() != &Method::GET {
            return Response::new(StatusCode::BadRequest, Some("invlid method".to_string()));
        }

        match request.path() {
            "/" => Response::new(StatusCode::Ok, Some("<h4>Welcome</h4>".to_string())),
            "/hello" => Response::new(StatusCode::Ok, Some("<h4>Hello</h4>".to_string())),
            "/ping" => Response::new(StatusCode::Ok, Some("<h4>pong</h4>".to_string())),
            p if p.starts_with("/static/") => self.read_file(p),
            _ => Response::new(StatusCode::BadRequest, Some("invlid path".to_string())),
        }
    }
}
