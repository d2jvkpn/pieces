use random_account::run;

use std::process;

fn main() {
    if let Err(msg) = run() {
        eprintln!("{:}", msg);
        process::exit(1);
    }
}
