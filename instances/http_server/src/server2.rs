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
    println!("+++ Accepting connection from: {}", addr);

    if let Err(e) = handle_stream2(Arc::new(stream)).await {
        println!("--- {} error: {}", addr, e);
        return;
    }

    println!("--- {} close connection", addr);
}

async fn handle_stream1(stream: Arc<TcpStream>) -> Res<()> {
    let mut stream = &*stream;
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

async fn handle_stream2(stream: Arc<TcpStream>) -> Res<()> {
    let mut stream = &*stream;
    let reader = BufReader::new(stream);
    let mut lines = reader.lines();

    while let Some(line) = lines.next().await {
        let line = line?;
        println!("<-- {}", line);
        stream.write_all(("<-- response: ".to_owned() + &line + "\n").as_bytes()).await?;
    }

    Ok(())
}
