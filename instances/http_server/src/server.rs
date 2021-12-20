use std::{
    convert::{TryFrom, TryInto},
    error,
    io::{self, Read, Write},
    net::{SocketAddr, TcpListener, TcpStream},
    thread,
};

use crate::http::{ParseError, Request, Response, StatusCode};

////
pub trait Handler {
    fn handle_request(&mut self, request: &Request) -> Response;

    fn handle_bad_request(&mut self, err: &ParseError) -> Response {
        println!("failed to parse request: {}", err);
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
        // listener.set_nonblocking(true)?; // no blocking self.listener.accept()
        Ok(Server { addr, listener })
    }

    // http service
    pub fn http(&self, handler: &mut dyn Handler) {
        println!("HTTP listening on {}", self.addr);

        'outer: loop {
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
            loop {
                if let Err(e) = handle(&mut stream, handler) {
                    eprintln!("client {} {}", addr, e);
                    // continue 'outer;
                    break;
                }
            }
        }
    }

    // chat service
    pub fn chat(&self) {
        println!("Chat listening on {}", self.addr);

        'outer: loop {
            let res = self.listener.accept(); // io.Result<(TcpStream, SocketAddr)>
            if res.is_err() {
                continue;
            }
            let (mut stream, addr) = res.unwrap();
            println!("client connected: {}", addr);

            thread::spawn(move || loop {
                if let Err(e) = chat(&mut stream, addr) {
                    eprintln!("client {} {}", addr, e);
                    // continue 'outer;
                    break;
                };
            });
        }
    }
}

fn chat(stream: &mut TcpStream, addr: SocketAddr) -> Result<(), String> {
    let mut buffer = [0; 1024];

    let size = match stream.read(&mut buffer) {
        Ok(s) => s,
        Err(e) => return Err(format!("stream.read error: {}", e)),
    };

    if size == 1 && buffer[0] as u16 == 10 {
        // return Err("EOF".to_string());
        return Err("EOF".to_string());
    }
    if size == 0 {
        return Err("disconnected".to_string());
    }

    match String::from_utf8((&buffer).to_vec()) {
        Ok(v) => print!("client {} read {} bytes: {}", addr, size, v),
        Err(e) => return Err(format!("invalid utf-8 sequence: {}", e)),
    };

    match stream.write(&buffer[0..size]) {
        Ok(_) => return Ok(()),
        Err(e) => return Err(format!("stream.write error: {}", e)),
    }
}

fn handle(stream: &mut TcpStream, handler: &mut dyn Handler) -> Result<(), String> {
    let mut buffer = [0; 1024];

    let size = match stream.read(&mut buffer) {
        Ok(s) => s,
        Err(e) => return Err(format!("stream.read error: {}", e)),
    };

    if size == 0 {
        return Err("disconnected".to_string());
    }

    // let text = String::from_utf8_lossy(&buffer);
    let req = match Request::try_from(&buffer[..]) {
        Ok(v) => v,
        Err(e) => {
            return Err(format!("parse request from buffer error: {}", e));
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
        return Err(format!("send response error: {}", e));
    }

    Ok(())
}
