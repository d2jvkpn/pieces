/// This struct defines the Path for a App route.
///
/// # Attributes
/// * prefix (String): the prefix of the view
pub struct Path {
    prefix: String,
}

impl Path {
    pub fn new(prefix: &str) -> Self {
        let prefix = prefix.trim_end_matches("/").to_owned() + "/";
        Path { prefix: prefix.to_string() }
    }
}

impl Path {
    /// This function defines a full path based on the struct's prefix and the String passed in.
    ///
    /// # Arguments
    /// * following_path (String): the rest of the path to be appended to the self.prefix
    ///
    /// # Use
    /// To use this in a route, we have to reference it:
    ///
    /// ```rust
    /// let path = Path{base: String::from("/base/")};
    /// app.route(&path.define(String::from("tail/path")), web::get().to(login::login))
    /// ```
    pub fn define(&self, following_path: &str) -> String {
        let p = following_path.trim_start_matches("/");
        return self.prefix.to_owned() + p;
    }
}
