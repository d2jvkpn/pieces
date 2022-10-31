use actix_web::{web::Query, HttpRequest, Responder};

use std::collections::HashMap;

/// This function defines a login view.
///
/// # Arguments
/// None
///
/// # Returns
/// (String) message stating that it's the login view
pub async fn login(req: HttpRequest, query: Query<HashMap<String, String>>) -> impl Responder {
    let language: &str = match req.headers().get("X-Language") {
        Some(v) => v.to_str().unwrap_or(""),
        None => "",
    };

    let platform = req.match_info().get("platform").unwrap_or("");

    let timestamp: u64 = match query.get("timestamp") {
        Some(v) => v.parse().unwrap_or(0),
        None => 0,
    };

    println!("<-- platform={:?}, language={:?}, timestamp={}", platform, language, timestamp);

    format!("Login view\n")
}

/// This function defines a logout view.
///
/// # Arguments
/// None
///
/// # Returns
/// (String) message stating that it's the logout view
pub async fn logout() -> String {
    format!("Logout view\n")
}
