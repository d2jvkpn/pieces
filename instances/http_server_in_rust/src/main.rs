#![allow(unused_variables)]
#![allow(dead_code)]
#![allow(unused_imports)]
#![allow(unused_labels)]

mod http;
mod server;
mod simple_handler;

use std::{env, process};

use http::query_string::QueryString;
use http::{method::Method, Request}; // http::request::Request
use http_server::hello;
use server::Server;
use simple_handler::SimpleHandler;

fn main() {
    // demo01();

    let addr = "127.0.0.1:8080".to_string();
    dbg!(&addr);

    let args: Vec<String> = env::args().collect();
    if args.len() == 1 {
        panic!("service required: echo or http");
    }

    let server = Server::new(addr).unwrap();

    match &args[1][..] {
        "echo" => {
            server.echo();
        }
        "http" => {
            let mut handler = SimpleHandler::new("./static").unwrap();
            server.run(&mut handler);
        }
        v => panic!("unknown argument: {}", v),
    };
}

fn demo01() {
    hello();

    // let get = Method::GET("abcd".to_string()); GET(String)
    // let delete = Method::DELETE(100); DELETE(i64)
    let get = Method::GET;
    let delete = Method::DELETE;
    let post = Method::POST;
    let put = Method::PUT;

    let request = Request {
        method: Method::GET,
        path: "/api/open/ping", // String::from("/api/open/ping")
        // query_string: Some("name=rover"), // Some(String::from("name=rover"))
        query_string: Some(QueryString::from("name=rover")),
    };

    dbg!(&request);
}
