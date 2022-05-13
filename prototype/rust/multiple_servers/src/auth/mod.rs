use actix_web::web;

mod login;
mod logout;
mod path;
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
    let app = app.route("/login", web::get().to(login::login));

    let p = &base_path.define("/logout");
    app.route(p, web::get().to(logout::logout));
}
