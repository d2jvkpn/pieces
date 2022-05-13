use actix_web::web;

mod auth;

/// This function combines the views from other view modules.
///
/// # Arguments
/// * (&mut web::ServiceConfig): reference to the app for configuration
///
/// # Returns
/// None
pub fn load_auth(app: &mut web::ServiceConfig) {
    auth::factory(app);
}
