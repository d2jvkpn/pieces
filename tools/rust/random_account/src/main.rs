use std::process;

use random_account::run;

fn main() {
    match run() {
        Ok(_) => process::exit(0),
        Err(msg) => {
            // eprintln!("{:?}", err);
            eprintln!("{:}", msg);
            process::exit(1);
        }
    }
}
