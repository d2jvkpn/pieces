use actix_web::web;

mod path;
mod user;
use path::Path;

/// This function adds the auth views to the web server.
///
/// # Arguments
/// * (&mut web::ServiceConfig): reference to the app for configuration
///
/// # Returns
/// None
pub fn factory(app: &mut web::ServiceConfig) {
    // define the path struct
    let base_path = Path::new("/auth");
    // define the routes for the app
    let app = app.route("/login/{platform}", web::post().to(user::login));

    let p = &base_path.define("/logout");
    app.route(p, web::post().to(user::logout));
}
