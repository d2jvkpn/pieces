#![allow(clippy::needless_return)]

use actix_web::{http::header::ContentType, web, App, HttpResponse, HttpServer};
use futures::future;
use serde_derive::{Deserialize, Serialize};

use std::io;

mod api;
use crate::api::v1;
use multiple_servers::load_auth;

#[actix_web::main]
async fn main() -> io::Result<()> {
    let (addr1, addr2) = ("0.0.0.0:8080", "0.0.0.0:8081");

    // produce future for server
    let s1 = HttpServer::new(move || {
        println!(">>> Http server is listening on {}", addr1);
        let app: App<_> = App::new();

        let scope = web::scope("/api/v1");
        let one = web::get().to(v1::utils_one);
        let greet1 = web::get().to(v1::greet);
        let greet2 = web::get().to(v1::greet);

        let router = scope
            .route("/one", one)
            .route("/greet", greet1)
            .route("/greet/{name}", greet2)
            .route("/index", web::get().to(index))
            .route("/index2", web::post().to(index2));

        return app.configure(load_auth).service(router).service(v1::hello);
    })
    .bind(addr1)?;

    // produce second future for server
    let s2 = HttpServer::new(move || {
        println!(">>> Health server is listening on {}", addr2);
        let app = App::new();

        let h1 = web::get().to(v1::health);
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

async fn index() -> HttpResponse {
    HttpResponse::Ok()
        .content_type(ContentType::plaintext())
        .insert_header(("X-Hdr", "sample"))
        .body("indx-data")
}

#[derive(Deserialize)]
struct Info {
    username: String,
}

#[derive(Deserialize, Serialize)]
struct User {
    name: Option<String>,
}

/// deserialize `Info` from request's body
async fn index2(user: web::Json<User>) -> String {
    let name = match user.name {
        Some(ref v) => v,
        None => "",
    };
    format!("Welcome {}!\n", name)
}
