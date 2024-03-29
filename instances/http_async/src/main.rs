// https://book.async.rs/tutorial/receiving_messages.html
#![allow(unused_imports)]

use std::{error, net, process, result};

use async_std::io::{BufReader, BufWriter};
use async_std::net::{TcpListener, TcpStream, ToSocketAddrs};
use async_std::{prelude::*, task};

type Res<T> = result::Result<T, Box<dyn error::Error + Send + Sync>>;

fn main() {
    if let Err(e) = serve() {
        eprintln!("{}", e);
        process::exit(1);
    }
}

// main
fn serve() -> Res<()> {
    let addr = "127.0.0.1:8080";
    let fut = accept_loop(addr);
    println!(">>> Listening on {}", addr);
    task::block_on(fut)
}

async fn accept_loop(addr: impl ToSocketAddrs) -> Res<()> {
    let listener = TcpListener::bind(addr).await?;

    while let Some(stream) = listener.incoming().next().await {
        let stream = stream?;
        match stream.peer_addr() {
            Err(e) => eprintln!("<-- incoming error: {}", e),
            Ok(addr) => {
                println!("<-- Accepting connection from: {}", addr);
                task::spawn(connection_loop(stream, addr));
            }
        }
    }

    Ok(())
}

async fn connection_loop(mut stream: TcpStream, addr: net::SocketAddr) -> Res<()> {
    // let mut writer = BufWriter::new(&stream);
    let mut lines = BufReader::new(stream.clone()).lines();

    // the first message as username
    let username = match lines.next().await {
        None => Err("peer disconnected immediately")?,
        Some(line) => line?,
    };
    println!("~~~ {} username: {}", addr, username);
    stream.write_all(format!(":) Hello, {}!\n", username).as_bytes()).await?;

    while let Some(line) = lines.next().await {
        let line = match line {
            Ok(line) if line.is_empty() => continue,
            Ok(line) => line,
            Err(e) => Err(e)?,
        };

        let (dest, msg) = match line.find(':') {
            None => {
                println!("~~~ {}: {:?}", username, line);
                stream.write_all(format!(">>> {:?}\n", line).as_bytes()).await?;
                continue;
            }
            Some(idx) => (&line[..idx], line[idx + 1..].trim()),
        };

        let dest: Vec<String> = dest.split(',').map(|v| v.trim().to_string()).collect();
        let msg: String = msg.to_string();

        stream.write_all(format!(">>> {:?}\n", line).as_bytes()).await?;

        println!("~~~ {}: des={:?}, msg={:?}", username, dest, msg);
    }

    println!("--> {} disconnected", username);
    Ok(())
}
