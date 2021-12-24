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

pub fn http(addr: String) -> Result<(), Box<dyn error::Error>> {
    let fut = accept_loop(&addr);
    println!(">>> Listening on {}", &addr);

    if let Err(e) = task::block_on(fut) {
        return Err(e);
    };
    Ok(())
}

async fn accept_loop(addr: impl ToSocketAddrs) -> Res<()> {
    let listener = TcpListener::bind(addr).await?;

    while let Some(stream) = listener.incoming().next().await {
        let stream = stream?;
        let addr = stream.peer_addr()?;
        println!("<-- Accepting connection from: {}", addr);
        let _handle = task::spawn(handle(Arc::new(stream), addr));
    }

    Ok(())
}

async fn handle(stream: Arc<TcpStream>, addr: net::SocketAddr) {
    if let Err(e) = handle_stream(stream).await {
        println!("{} error: {}", addr, e);
    }
    println!("{} close connection", addr);
}

async fn handle_stream(stream: Arc<TcpStream>) -> Res<()> {
    let mut stream = &*stream;

    let mut reader = BufReader::new(stream);
    let mut buffer = [0; 1024];

    match reader.read(&mut buffer).await {
        Ok(v) => { /*println!("{}", v)*/ }
        Err(e) => Err(e)?,
    }

    let req = Request::try_from(&buffer[..])
        .map_err(|e| format!("parse request from buffer error: {}", e))?;

    stream.write_all("HTTP/1.1 200 Ok\r\n\r\nHello, world!\n".as_bytes()).await?;

    Ok(())
}
