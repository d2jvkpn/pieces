use std::{error, io, marker, net, process, result};

use crate::http::{ParseError, Request, Response, StatusCode};
use crate::server;

use async_std::{
    io::{BufReader, BufWriter},
    net::{TcpListener, TcpStream, ToSocketAddrs},
    prelude::*,
    sync::Arc,
    task,
};

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
        let _handle = task::spawn(handle(stream));
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
    println!("<-- Accepting connection from: {}", addr);

    if let Err(e) = handle_stream(Arc::new(stream)).await {
        println!("{} error: {}", addr, e);
        return;
    }

    println!("{} close connection", addr);
}

async fn handle_stream(stream: Arc<TcpStream>) -> Res<()> {
    let mut stream = &*stream;
    let mut buffer = [0; 1024];
    let mut reader = BufReader::new(stream);

    match reader.read(&mut buffer).await {
        // no need to check v == 0 when client closed the connection
        Ok(v) => { /*println!("{}", v)*/ }
        Err(e) => Err(e)?,
    }

    println!("read message: {}", String::from_utf8_lossy(&buffer));

    //    let req = Request::try_from(&buffer[..])
    //        .map_err(|e| format!("parse request from buffer error: {}", e))?;

    stream.write_all("HTTP/1.1 200 Ok\r\n\r\nHello, world!\n".as_bytes()).await?;
    Ok(())
}
