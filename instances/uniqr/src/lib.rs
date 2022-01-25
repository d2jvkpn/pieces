use std::io::{self, BufRead, BufReader, Write};
use std::{env, error, fs};

use clap::{App, Arg};

type MyResult<T> = Result<T, Box<dyn error::Error>>;

#[derive(Debug)]
pub struct Config {
    in_file: String,
    out_file: Option<String>,
    count: bool,
}

// --------------------------------------------------
pub fn get_args() -> MyResult<Config> {
    let matches = App::new(env!("CARGO_PKG_HOMEPAGE"))
        .version(env!("CARGO_PKG_VERSION"))
        .author(env!("CARGO_PKG_AUTHORS"))
        .about(env!("CARGO_PKG_DESCRIPTION"))
        .arg(
            Arg::with_name("in_file")
                .value_name("IN_FILE")
                .help("Input file")
                // .default_value("-")
                .required(true),
        )
        .arg(
            Arg::with_name("out_file")
                .value_name("OUT_FILE")
                .help("Output file"),
        )
        .arg(
            Arg::with_name("count")
                .short("c")
                .long("count")
                .help("Show counts")
                .takes_value(false),
        )
        .get_matches();

    Ok(Config {
        //in_file: matches.value_of_lossy("in_file").unwrap().to_string(),
        //in_file: matches.value_of_lossy("in_file").map(String::from).unwrap(),
        //in_file: matches.value_of_lossy("in_file").map(|v| v.into()).unwrap(),
        in_file: matches.value_of_lossy("in_file").map(Into::into).unwrap(),
        out_file: matches.value_of("out_file").map(|v| v.to_string()),
        count: matches.is_present("count"),
    })
}

// --------------------------------------------------
pub fn run(config: Config) -> MyResult<()> {
    let mut file = open(&config.in_file).map_err(|e| format!("{}: {}", config.in_file, e))?;

    let mut out_file: Box<dyn Write> = match &config.out_file {
        Some(out_name) => Box::new(fs::File::create(out_name)?),
        _ => Box::new(io::stdout()),
    };

    let mut print = |count: u64, text: &str| -> MyResult<()> {
        if count == 0 {
            return Ok(());
        };

        if config.count {
            write!(out_file, "{:>4} {}", count, text)?;
        } else {
            write!(out_file, "{}", text)?;
        }

        Ok(())
    };

    let [mut line, mut previous] = [String::new(), String::new()];
    let mut count: u64 = 0;

    loop {
        let bytes = file.read_line(&mut line)?;
        if bytes == 0 {
            break;
        }

        if line.trim_end() != previous.trim_end() {
            print(count, &previous)?;
            previous = line.clone();
            count = 0;
        }

        count += 1;
        line.clear();
    }
    print(count, &previous)?;

    Ok(())
}

// --------------------------------------------------
fn open(filename: &str) -> MyResult<Box<dyn BufRead>> {
    match filename {
        "-" => Ok(Box::new(BufReader::new(io::stdin()))),
        _ => Ok(Box::new(BufReader::new(fs::File::open(filename)?))),
    }
}
