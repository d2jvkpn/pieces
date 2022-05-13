#![allow(clippy::needless_return)]

use actix_web::web::{self, Query};
use actix_web::{get, App, HttpRequest, HttpServer, Responder};
use futures::future;

use std::collections::HashMap;
use std::io;

use multiple_servers::load_auth;

async fn utils_one() -> impl Responder {
    "Utils one reached\n"
}

async fn health() -> impl Responder {
    "All good\n"
}

#[get("/hello/{name}")]
async fn hello(name: web::Path<String>) -> impl Responder {
    format!("Hello, {name}!\n")
}

/// This is a basic async function that represents a view for the server.
///
/// # Arguments
/// * req (HttpRequest): the http request body that is passed into the view
///
/// # Returns
/// * (Responder): object that has implements the Responder trait
async fn greet(req: HttpRequest, query: Query<HashMap<String, String>>) -> impl Responder {
    // let name = req.match_info().get("name").unwrap_or("world");
    let name: &str = match req.match_info().get("name") {
        Some(v) => v,
        _ => match query.get("name") {
            Some(v) => v,
            None => "world",
        },
    };

    println!("~~~ {:?}: name={}", req.connection_info(), name);
    format!("Hello, {}!\n", name)
}

#[actix_rt::main]
async fn main() -> io::Result<()> {
    let addr1 = "0.0.0.0:8080";
    let addr2 = "0.0.0.0:3306";

    // produce future for server
    let s1 = HttpServer::new(move || {
        println!(">>> Http server is listening on {}", addr1);
        let app: App<_> = App::new();

        let scope = web::scope("/api/v1");
        let one = web::get().to(utils_one);
        let greet1 = web::get().to(greet);
        let greet2 = web::get().to(greet);

        let router =
            scope.route("/one", one).route("/greet", greet1).route("/greet/{name}", greet2);

        return app.configure(load_auth).service(router).service(hello);
    })
    .bind(addr1)?;

    // produce second future for server
    let s2 = HttpServer::new(move || {
        println!(">>> Health server is listening on {}", addr2);
        let app = App::new();

        let h1 = web::get().to(health);
        let entry = web::resource("/health").route(h1);

        app.service(entry)
    })
    .bind(addr2)?;

    // join both server futures and run them
    future::try_join(s1.workers(4).run(), s2.workers(1).run()).await?;
    // future::join(s1, s2).await;
    Ok(())

    // s1.workers(4).run().await
}
// let numbers = Query::<HashMap<String, u32>>::from_query("one=1&two=2").unwrap();
// assert_eq!(numbers.get("one"), Some(&1));
// assert_eq!(numbers.get("two"), Some(&2));
// assert!(numbers.get("three").is_none());
