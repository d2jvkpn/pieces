// use std::env;

//fn main() {
//    let msg = env::args().nth(1).expect("Missing the message. Usage: catsay < message>");

//    println!(
//        r#"< {msg} >
// \\
//   \\
//     /\\_/\\
//     ( o o )
//     =( I )="#,
//        msg=msg
//    )
//}

// extern crate colored;
// extern crate structopt;
// other crates strfmt, failure

use std::io::{self, Read};
use std::{fs, path};

use colored::*;
use structopt::StructOpt;

#[derive(StructOpt)]
struct Options {
    #[structopt(default_value = "Meow!")]
    /// What does the cat say?
    msg: String,

    #[structopt(short = "d", long = "dead")]
    /// Make the cat appear dead
    dead: bool,

    #[structopt(short = "f", long = "file", parse(from_os_str))]
    /// Load the cat picture from the specified file
    cat_file: Option<path::PathBuf>,

    #[structopt(short = "i", long = "stdin")]
    /// Read the message from STDIN instead of the argument
    stdin: bool,
}

fn main() {
    // -> Result<(), Box<dyn std::error::Error>> {
    let options = Options::from_args();
    // let msg = options.msg;

    let mut msg = String::new();

    if options.stdin {
        io::stdin().read_to_string(&mut msg).expect("can't read from stdin");
    } else {
        msg = options.msg;
    };

    let eye = if options.dead { "x" } else { "0" };

    println!("{}", msg.bright_yellow().underline().on_purple());

    if let Some(v) = &options.cat_file {
        let template = fs::read_to_string(v).expect(&format!("could not read file {:?}", v));

        println!("{}", &template.replace("{eye}", eye).trim()); // TODO: eye.red().bold()
    } else {
        println!(
            r#" \\
   \\
     /\\_/\\
     ( {eye} {eye} )
     =( I )="#,
            eye = eye.red().bold(),
        );
    }
}
