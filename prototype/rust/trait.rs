struct RustDev {
    awesome: bool,
    name: String,
}

struct GoDev {
    awesome: bool,
}

trait Developer {
    fn new(awesome: bool) -> Self;
    fn language(&self) -> &str;
    fn hello(&self) {
        println!("Hello, world!"); // tag::a
    }
}

impl GoDev {
    fn hello(&self) {
        println!("Go is awesome!"); // tag::b
    }
}

impl RustDev {
    fn new(awesome: bool) -> Self {
        RustDev {
            awesome,
            name: "rust developer".to_string(),
        }
    }
}

impl Developer for RustDev {
    fn new(awesome: bool) -> Self {
        RustDev::new(awesome)
    }

    fn language(&self) -> &str {
        &self.name
    }
}

impl Developer for GoDev {
    fn new(awesome: bool) -> Self {
        GoDev { awesome }
    }

    fn language(&self) -> &str {
        "gopher"
    }

    fn hello(&self) {
        println!("Hello, world from Gopher!"); // tag::c
    }
}

fn main() {
    println!(">>> 1");
    let go = GoDev::new(true);
    go.hello(); // self method (tag::b) overrides Developer.hello (tag::a) and impl (tag::c)
    println!("{}", go.language());

    println!(">>> 2");
    let rust = RustDev::new(true);
    rust.hello(); // not implemented by RuestDev, but by Developer -> tag::a
    println!("{}", rust.language());

    println!(">>> 3");
    let dev: RustDev = Developer::new(true);
    dev.hello();
}
