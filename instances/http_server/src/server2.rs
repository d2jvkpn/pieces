#![allow(unused_mut)]

use std::{error, io, marker, net, process, result};

use crate::http::{method::Method, ParseError, Request, Response, StatusCode};
use crate::server;

use async_std::io::{BufReader, BufWriter};
use async_std::net::{TcpListener, TcpStream, ToSocketAddrs};
use async_std::{prelude::*, sync::Arc, task};

type Res<T> = result::Result<T, Box<dyn error::Error + Send + Sync>>;

pub fn http(addr: &str) -> Result<(), Box<dyn error::Error>> {
    let fut = accept_loop(addr);
    println!(">>> Listening on {}", addr);

    if let Err(e) = task::block_on(fut) {
        return Err(e);
    };
    Ok(())
}

// addr: impl ToSocketAddrs
async fn accept_loop(addr: &str) -> Res<()> {
    let listener = TcpListener::bind(addr).await?;

    while let Some(stream) = listener.incoming().next().await {
        let stream = stream?;
        task::spawn(handle(stream));
    }

    Ok(())
}

// can't reuse a connection
async fn handle(stream: TcpStream) {
    // addr: net::SocketAddr
    let addr = match stream.peer_addr() {
        Ok(v) => v,
        Err(e) => {
            println!("get peer_addr error: {}", e);
            return;
        }
    };
    println!("+++ Accepting connection from: {}", addr);

    if let Err(e) = handle_stream3(stream, addr.to_string()).await {
        println!("--- {} error: {}", addr, e);
        return;
    }

    println!("--- {} close connection", addr);
}

async fn handle_stream1(mut stream: TcpStream, addr: String) -> Res<()> {
    //    let mut buffer = [0; 1024];
    //    let mut reader = BufReader::new(stream);
    //    match reader.read(&mut buffer).await {
    //        // no need to check v == 0 when client closed the connection
    //        Ok(v) => { println!("{}", v) }
    //        Err(e) => Err(e)?,
    //    }

    let mut buffer = vec![0u8; 1024];
    let size = stream.read(&mut buffer).await?;
    println!("<-- read message(size={}): {}", size, String::from_utf8_lossy(&buffer));

    //    let req = Request::try_from(&buffer[..])
    //        .map_err(|e| format!("parse request from buffer error: {}", e))?;

    stream.write_all("HTTP/1.1 200 Ok\r\n\r\nHello, world!\n".as_bytes()).await?;
    Ok(())
}

async fn handle_stream2(mut stream: TcpStream, addr: String) -> Res<()> {
    let mut reader = BufReader::new(stream.clone());

    // let mut buffer = vec![0u8; 1024];
    // while let Ok(v) = reader.read_until(b'\n', &mut buffer).await {}
    // println!("<-- {}", String::from_utf8_lossy(&buffer));

    let mut buffer = String::new();
    while let Ok(v) = reader.read_to_string(&mut buffer).await {
        if v == 0 {
            break;
        }
        println!("size: {}", v);
        println!("<-- {}", buffer);
        stream.write_all("ok".as_bytes()).await?;
    }

    Ok(())
}

async fn handle_stream3(mut stream: TcpStream, addr: String) -> Res<()> {
    let reader = BufReader::new(stream.clone());
    let mut lines = reader.lines();

    let mut s = 0_usize;
    let mut blocks = vec![];
    while let Some(line) = lines.next().await {
        let line = line?;
        if line != "" {
            s += line.len() + 1;
            blocks.push(line);
            continue;
        }

        if blocks.len() == 0 {
            return Ok(());
        }
        blocks[0].push_str("\r\n\r\n");
        println!("<-- {} read: size={}, {:?}", addr, s, blocks);

        // Handle "Keep-Alive: timeout=5, max=1000" or "Connection: close"
        let (path, mut response) = handle_request(&blocks[0]);
        // response.body = Some(path);
        let d = String::from("");
        let body = response.body.as_ref().unwrap_or(&d);
        let res_str = format!("{}\r\n\r\n{}\n", response, body);
        println!("--> {} {} {}", addr, path, response.body.unwrap_or("".to_string()));
        stream.write_all(res_str.as_bytes()).await?;
        s = 0;
        blocks.clear();
    }
    Ok(())
}

fn handle_request(req_str: &str) -> (String, Response) {
    let request = match Request::try_from(req_str.as_bytes()) {
        Ok(v) => v,
        Err(e) => return ("".to_string(), Response::new(StatusCode::BadRequest, None)),
    };

    let path = request.path().to_string();

    if request.method() != &Method::GET {
        let response = Response::new(StatusCode::BadRequest, Some("invlid method".to_string()));
        return (path, response);
    }

    let response = match request.path() {
        "/" => Response::new(StatusCode::Ok, Some("<h4>Welcome</h4>".to_string())),
        "/hello" => Response::new(StatusCode::Ok, Some("<h4>Hello</h4>".to_string())),
        "/ping" => Response::new(StatusCode::Ok, Some("<h4>pong</h4>".to_string())),
        // p if p.starts_with("/static/") => self.read_file(p),
        _ => Response::new(StatusCode::BadRequest, Some("invlid path".to_string())),
    };

    (path, response)
}
