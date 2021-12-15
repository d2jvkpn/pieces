use std::process;

use headr::{get_args, run};

fn main() {
    if let Err(e) = get_args().and_then(run) {
        eprintln!("{}", e);
        process::exit(1);
    }
}
