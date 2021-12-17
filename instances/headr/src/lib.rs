use std::io::{self, BufRead, BufReader, Read};
use std::{env, error, fs};

use clap::{App, Arg};
use flate2::bufread::GzDecoder;

type MyResult<T> = Result<T, Box<dyn error::Error>>;

#[derive(Debug)]
pub struct Config {
    files: Vec<String>,
    lines: usize,
    bytes: Option<usize>,
}

pub fn get_args() -> MyResult<Config> {
    let matches = App::new(env!("CARGO_PKG_HOMEPAGE"))
        .version(env!("CARGO_PKG_VERSION"))
        .author(env!("CARGO_PKG_AUTHORS"))
        .about(env!("CARGO_PKG_DESCRIPTION"))
        .arg(
            Arg::with_name("lines")
                .value_name("LINES")
                .short("n")
                .long("lines")
                .help("Number of lines")
                .default_value("10"), // .required(true)
        )
        .arg(
            Arg::with_name("bytes")
                .value_name("BYTES")
                .short("c")
                .long("bytes")
                .takes_value(true)
                .conflicts_with("lines")
                .help("Number of bytes"),
        )
        .arg(
            Arg::with_name("files")
                .value_name("FILE")
                .help("Input file(s)")
                .required(true)
                // .default_value("-")
                .min_values(1),
        )
        .get_matches();

    let lines: Option<usize> = matches
        .value_of("lines") // Option<&str>
        // to unpack a &str from Some and send it to parse_positive_int,
        // return if then Option is None which stops chain operations
        .map(parse_positive_int)
        .transpose() // convert <Option<Result>> to <Result<Option>>
        // create an informative error message. Use ? to propagate an Err or unpack the Ok value
        .map_err(|e| format!("illegal line count -- {}", e))?;

    // let lines = lines.unwrap();
    // .default_value("10") or .required(true) makes sure -n always privided
    let lines: usize = lines.ok_or("arg '-n' is not provided")?;

    let bytes = matches
        .value_of("bytes")
        .map(parse_positive_int)
        .transpose()
        .map_err(|e| format!("illegal byte count -- {}", e))?;

    dbg!(&lines);

    Ok(Config { files: matches.values_of_lossy("files").unwrap(), lines, bytes })
}

////
fn parse_positive_int(val: &str) -> MyResult<usize> {
    match val.parse() {
        Ok(n) if n > 0 => Ok(n),
        _ => Err(From::from(val)),
    }
}

#[test]
fn test_parse_positive_int() {
    // 3 is an OK integerlet
    let res = parse_positive_int("3");
    assert!(res.is_ok());
    assert_eq!(res.unwrap(), 3);

    // Any string is an error
    let res = parse_positive_int("foo");
    assert!(res.is_err());
    assert_eq!(res.unwrap_err().to_string(), "foo".to_string());

    // A zero is an error
    let res = parse_positive_int("0");
    assert!(res.is_err());
    assert_eq!(res.unwrap_err().to_string(), "0".to_string());
}

///
pub fn run(config: Config) -> MyResult<()> {
    dbg!(&config);

    let mut n_failed = 0;

    for filename in &config.files {
        match open(&filename) {
            Err(err) => {
                eprintln!("==> !!! {}: {} <==", filename, err);
                n_failed += 1;
            }

            Ok(mut reader) => {
                println!("==> {} <==", filename);
                if let Err(e) = process_buf_read(&config, &mut reader) {
                    eprintln!("!!! {}", e);
                    n_failed += 1;
                };
            }
        }
    }

    match n_failed {
        0 => Ok(()),
        n => {
            Err(From::from(format!("!!! headr {} file{} failed", n, if n > 1 { "s" } else { "" })))
        }
    }
}

pub fn open(filename: &str) -> MyResult<Box<dyn BufRead>> {
    let file = match filename {
        "-" => return Ok(Box::new(BufReader::new(io::stdin()))),
        _ => Box::new(BufReader::new(fs::File::open(filename)?)),
    };

    match filename {
        v if v.ends_with(".gz") => {
            let reader = io::BufReader::new(file);
            Ok(Box::new(io::BufReader::new(GzDecoder::new(reader))))
        }
        _ => Ok(file),
    }
}

fn process_buf_read(config: &Config, reader: &mut dyn BufRead) -> MyResult<()> {
    if let Some(num_bytes) = config.bytes {
        let mut buffer = vec![0; num_bytes];
        let mut handle = reader.take(num_bytes as u64);
        let n = handle.read(&mut buffer)?;

        // !! awlways add a new line at the end
        print!("{}\n", String::from_utf8_lossy(&buffer[..n]));
        return Ok(());
    }

    let mut line = String::new();

    for _ in 0..config.lines {
        let bytes = reader.read_line(&mut line)?;
        if bytes == 0 {
            break;
        }
        print!("{}", line);
        line.clear();
    }
    Ok(())
}
