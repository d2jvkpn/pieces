#![allow(unused_variables)]
#![allow(dead_code)]

use std::process;

mod http;
mod server;
mod website_handler;

use http::query_string::QueryString;
use http::{method::Method, Request}; // http::request::Request
use http_server::hello;
use server::Server;
use website_handler::WebsiteHandler;

fn main() {
    // demo01();

    let addr = "127.0.0.1:8080".to_string();
    dbg!(&addr);

    let public_path = "./static".to_string();
    dbg!(&public_path);

    let server = Server::new(addr);
    let mut handler = WebsiteHandler::new(public_path);

    if let Err(err) = server.run(&mut handler) {
        eprintln!("server.run(handler): {:?}", err);
        process::exit(1);
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
