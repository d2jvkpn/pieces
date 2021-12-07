use std::convert::{TryFrom, TryInto};
use std::io::{Read, Write};
use std::net::{SocketAddr, TcpListener, TcpStream};
use std::{error, io};

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
    listener: TcpListener,
}

impl Server {
    pub fn new(addr: String) -> Result<Server, io::Error> {
        let listener = TcpListener::bind(&addr)?;
        Ok(Server { addr, listener })
    }

    pub fn run(&self, handler: &mut dyn Handler) {
        println!("HTTP listening on {}", self.addr);

        loop {
            let (mut stream, addr) = match self.listener.accept() {
                Ok((s, a)) => {
                    println!(">>> New tcp connection: {}", a);
                    (s, a)
                }
                Err(e) => {
                    eprintln!("!!! Failed to establish a connection: {}", e);
                    continue;
                }
            };

            handle(&mut stream, addr, handler);
        }

        // return Ok(());
    }

    pub fn echo(&self) {
        println!("Echo listening on {}", self.addr);

        'outer: loop {
            let res = self.listener.accept(); // io.Result<(TcpStream, SocketAddr)>
            if res.is_err() {
                continue;
            }
            let (mut stream, addr) = res.unwrap();
            println!("client connected: {}", addr);

            loop {
                if let Err(e) = echo(&mut stream, addr) {
                    println!("echo {} error: {}", addr, e);
                    continue 'outer;
                };
            }
        }
    }
}

fn echo(stream: &mut TcpStream, addr: SocketAddr) -> Result<(), String> {
    let mut buffer = [0; 1024];

    let size = match stream.read(&mut buffer) {
        Ok(s) => s,
        Err(e) => {
            return Err(format!("error read from {}: {}", addr, e));
        }
    };

    if size == 1 && buffer[0] as u16 == 10 {
        return Err("EOF".to_string());
    }
    if size == 0 {
        return Err("disconnected".to_string());
    }

    match String::from_utf8((&buffer).to_vec()) {
        Ok(v) => print!("    read {} bytes from {}: {}", size, addr, v),
        Err(e) => {
            return Err(format!("error invalid utf-8 sequence from {}: {}", addr, e));
        }
    };

    match stream.write(&buffer[0..size]) {
        Ok(_) => {}
        Err(e) => return Err(format!("error write to {}: {}", addr, e)),
    }

    Ok(())
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

    let text = String::from_utf8_lossy(&buffer);
    let req = match Request::try_from(&buffer[..]) {
        Ok(v) => v,
        Err(e) => {
            eprintln!("    failed to parse request from buffer {}: {}", addr, e);
            return;
        }
    };
    // let req: &Result<Request, _> = &buffer[..].try_into();

    //    let response = Response::new(StatusCode::Ok, Some("hello, world!\n".to_string()));
    //    dbg!(addr, req, &response);

    //    if let Err(e) = write!(stream, "{}", &response) {
    //        eprintln!("    error write to {}: {}", addr, e);
    //        return;
    //    }

    let response = handler.handle_request(&req);

    if let Err(e) = response.send(stream) {
        eprintln!("    failed to send response to {}: {}", addr, e);
        return;
    }
}
