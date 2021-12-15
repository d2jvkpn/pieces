use std::io::{self, BufRead, BufReader};
use std::{env, error, fs};

use clap::{App, Arg};
use flate2::bufread::GzDecoder;

type MyResult<T> = Result<T, Box<dyn error::Error>>;

#[derive(Debug, Default)]
pub struct Config {
    pub files: Vec<String>,
    pub number_lines: bool,
    pub number_nonblank_lines: bool,
}

#[allow(dead_code)]
impl Config {
    pub fn new() -> Config {
        Config { ..Default::default() }
    }

    pub fn number_lines(&self) -> bool {
        return self.number_lines;
    }
}

#[test]
fn test_config() {
    let config = Config::new();

    assert_eq!(config.files.len(), 0);
    assert_eq!(config.number_lines, false);
    assert_eq!(config.number_lines(), false);
    assert_eq!(config.number_nonblank_lines, false);
}

pub fn get_args() -> MyResult<Config> {
    let matches = App::new(env!("CARGO_PKG_HOMEPAGE"))
        .version(env!("CARGO_PKG_VERSION"))
        .author(env!("CARGO_PKG_AUTHORS"))
        .about(env!("CARGO_PKG_DESCRIPTION"))
        .arg(
            Arg::with_name("files")
                .value_name("FILE")
                .help("Input file(s)")
                .required(true)
                // .default_value("-")
                .min_values(1),
        )
        .arg(
            Arg::with_name("number")
                .help("Number lines")
                .short("n")
                .long("number")
                .takes_value(false)
                .conflicts_with("number_nonblank"),
        )
        .arg(
            Arg::with_name("number_nonblank")
                .help("Number non-blank lines")
                .short("b")
                .long("number-nonblank")
                .takes_value(false),
        )
        .get_matches();

    Ok(Config {
        files: matches.values_of_lossy("files").unwrap(),
        number_lines: matches.is_present("number_lines"),
        number_nonblank_lines: matches.is_present("number_nonblank"),
    })
}

pub fn run(config: Config) -> MyResult<()> {
    // dbg!(&config);
    // println!("{}", config.number_lines());

    for filename in &config.files {
        match open(&filename) {
            Err(err) => eprintln!("Failed to open {}: {}", &filename, err),
            Ok(buf_read) => {
                // eprintln!("Opened {}", filename),
                print_buf_read(&config, buf_read)?
            }
        }
    }

    Ok(())
}

pub fn open(filename: &str) -> MyResult<Box<dyn BufRead>> {
    let file = match filename {
        "-" => return Ok(Box::new(BufReader::new(io::stdin()))),
        _ => Box::new(BufReader::new(fs::File::open(filename)?)),
    };

    match filename {
        filename if filename.ends_with(".gz") => {
            let reader = io::BufReader::new(file);
            Ok(Box::new(io::BufReader::new(GzDecoder::new(reader))))
        }
        _ => Ok(file),
    }
}

pub fn print_buf_read(config: &Config, buf_read: Box<dyn BufRead>) -> MyResult<()> {
    let mut last_num = 0;

    for (index, result) in buf_read.lines().enumerate() {
        let line = result?;

        if config.number_lines {
            println!("{:>6}\t{}", index + 1, line);
        } else if config.number_nonblank_lines {
            if !line.is_empty() {
                last_num += 1;
                println!("{:>6}\t{}", last_num, line);
            } else {
                println!();
            }
        } else {
            println!("{}", line);
        }
    }

    Ok(())
}
