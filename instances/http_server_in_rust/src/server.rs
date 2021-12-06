#![allow(unused_labels)]
#![allow(unused_imports)]

use std::convert::{TryFrom, TryInto};
use std::error;
use std::io::{Read, Write};
use std::net::{SocketAddr, TcpListener, TcpStream};

use crate::http::{ParseError, Request, Response, StatusCode};

////
pub trait Handler {
    fn handle_request(&mut self, request: &Request) -> Response;

    fn handle_bad_request(&mut self, err: &ParseError) -> Response {
        println!("Failed to parse request: {}", err);
        Response::new(StatusCode::BadRequest, None)
    }
}

////
pub struct Server {
    addr: String,
}

impl Server {
    pub fn new(addr: String) -> Server {
        Server { addr }
    }

    pub fn run(&self, handler: &mut dyn Handler) -> Result<(), Box<dyn error::Error>> {
        println!("HTTP Listening on {}", self.addr);

        let listener = match TcpListener::bind(&self.addr) {
            Ok(v) => v,
            Err(e) => return Err(From::from(e)),
        };

        'outer: loop {
            //			let res = listener.accept(); // io.Result<(TcpStream, SocketAddr)>
            //			if res.is_err() {
            //				continue;
            //			}
            //			let (stream, addr) = res.unwrap();

            let (mut stream, addr) = match listener.accept() {
                Ok((s, a)) => {
                    println!(">>> New tcp connection: {}", a);
                    (s, a)
                }
                Err(e) => {
                    eprintln!("!!! Failed to establish a connection: {}", e);
                    continue;
                }
            };

            //			loop{
            //				handle_stream(&mut stream, addr);
            //			}
            handle(&mut stream, addr, handler);
        }

        // return Ok(());
    }
}

fn handle(stream: &mut TcpStream, addr: SocketAddr, handler: &mut dyn Handler) {
    let mut buffer = [0; 1024];

    let size = match stream.read(&mut buffer) {
        Ok(s) => s,
        Err(e) => {
            eprintln!("    error read from {}: {}", addr, e);
            return;
        }
    };

    //        let text = match String::from_utf8(&buffer) {
    //            Ok(v) => v,
    //            Err(e) => {
    //                eprintln!("    error invalid utf-8 sequence from {}: {}", addr, e);
    //                return;
    //            }
    //        };

    let text = String::from_utf8_lossy(&buffer);
    let req = match Request::try_from(&buffer[..]) {
        Ok(v) => v,
        Err(e) => {
            eprintln!("    Request::try_from buffer {}: {}", addr, e);
            return;
        }
    };
    // let req: &Result<Request, _> = &buffer[..].try_into();

    let response = handler.handle_request(&req);

    //    let response = Response::new(StatusCode::Ok, Some("hello, world!\n".to_string()));
    //    dbg!(addr, req, &response);

    //    if let Err(e) = write!(stream, "{}", &response) {
    //        eprintln!("    error write to {}: {}", addr, e);
    //        return;
    //    }

    if let Err(e) = response.send(stream) {
        eprintln!("    error write to {}: {}", addr, e);
        return;
    }

    //    print!("    read {} bytes from {}: {}", size, addr, text);
    //    match stream.write(&buffer[0..size]) {
    //        Ok(s) => {
    //            println!("    write {} bytes to {}", s, addr);
    //        }
    //        Err(e) => {
    //            eprintln!("    error write to {}: {}", addr, e);
    //            return;
    //        }
    //    }
}
