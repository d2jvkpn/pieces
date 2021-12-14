use std::process;

use catr::{get_args, run};

fn main() {
    if let Err(e) = get_args().and_then(run) {
        eprintln!("{}", e);
        process::exit(1);
    }
}
