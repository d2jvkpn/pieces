use actix_web::{
    get,
    web::{self, Query},
    HttpRequest, Responder,
};

use std::collections::HashMap;

pub async fn utils_one() -> impl Responder {
    "Utils one reached\n"
}

pub async fn health() -> impl Responder {
    "All good\n"
}

#[get("/hello/{name}")]
pub async fn hello(name: web::Path<String>) -> impl Responder {
    format!("Hello, {name}!\n")
}

/// This is a basic async function that represents a view for the server.
///
/// # Arguments
/// * req (HttpRequest): the http request body that is passed into the view
///
/// # Returns
/// * (Responder): object that has implements the Responder trait
pub async fn greet(req: HttpRequest, query: Query<HashMap<String, String>>) -> impl Responder {
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
