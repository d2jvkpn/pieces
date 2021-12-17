use std::io::{self, BufRead, BufReader};
use std::{env, error, fmt, fs};

use clap::{App, Arg};
use flate2::bufread::GzDecoder;

type MyResult<T> = Result<T, Box<dyn error::Error>>;

#[derive(Debug, Default)]
pub struct Config {
    files: Vec<String>,
    lines: bool,
    words: bool,
    chars: bool,
    bytes: bool,
}

#[derive(Debug, Default, PartialEq)]
pub struct FileInfo {
    num_lines: usize,
    num_words: usize,
    num_bytes: usize,
    num_chars: usize,
}

impl fmt::Display for FileInfo {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(
            f,
            "num_lines: {}, num_words: {}, num_bytes: {}, num_chars: {}",
            self.num_lines, self.num_words, self.num_bytes, self.num_chars,
        )
    }
}

impl FileInfo {
    pub fn format_field(&self, config: &Config) -> String {
        fn ff(value: usize, show: bool) -> String {
            // same as fn format_field
            return if show { format!("{:>8}", value) } else { "".to_string() };
        }

        format!(
            "{}{}{}{}",
            ff(self.num_lines, config.lines),
            ff(self.num_words, config.words),
            ff(self.num_bytes, config.bytes),
            ff(self.num_chars, config.chars),
        )
    }
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
                // .default_value("-")
                .min_values(1),
        )
        .arg(
            Arg::with_name("lines")
                .value_name("LINES")
                .help("Show line count")
                .takes_value(false)
                .short("l")
                .long("lines"),
        )
        .arg(
            Arg::with_name("words")
                .value_name("WORDS")
                .help("Show word count")
                .takes_value(false)
                .short("w")
                .long("words"),
        )
        .arg(
            Arg::with_name("chars")
                .value_name("CHARS")
                .help("Show character count")
                .takes_value(false)
                .short("m")
                .long("chars")
                .conflicts_with("bytes"),
        )
        .arg(
            Arg::with_name("bytes")
                .value_name("BYTES")
                .help("Show byte count")
                .takes_value(false)
                .short("c")
                .long("bytes"),
        )
        .get_matches();

    //    Ok(Config {
    //        ..Default::default()
    //    })

    let mut lines = matches.is_present("lines");
    let mut words = matches.is_present("words");
    let mut chars = matches.is_present("chars");
    let mut bytes = matches.is_present("bytes");

    // Iterator::any, all, filter, map, find, position, cmp, min_by, max_by
    if [lines, words, bytes, chars].iter().all(|v| v == &false) {
        lines = true;
        words = true;
        bytes = true;
        chars = false;
    }

    Ok(Config { files: matches.values_of_lossy("files").unwrap(), lines, words, bytes, chars })
}

pub fn run(config: Config) -> MyResult<()> {
    // dbg!(&config); // println!("{:#?}", config);
    let mut n_failed = 0;

    for name in &config.files {
        let result = match open(name) {
            Err(err) => Err(From::from(format!("{}: {}", name, err))),
            Ok(reader) => count(reader),
        };

        let name = if name.as_str() == "-" { "".to_string() } else { format!(" {}", name) };

        match result {
            Err(e) => {
                eprintln!("{}", e);
                n_failed += 1;
            }
            Ok(v) => println!("{}{}", v.format_field(&config), name),
        };
    }

    if n_failed == 0 {
        Ok(())
    } else {
        Err(From::from(n_failed.to_string()))
    }
}

pub fn open(filename: &str) -> MyResult<Box<dyn BufRead>> {
    let file = match filename {
        "-" => return Ok(Box::new(BufReader::new(io::stdin()))),
        _ => Box::new(BufReader::new(fs::File::open(filename)?)),
    };

    if filename.ends_with(".gz") {
        let reader = io::BufReader::new(file);
        Ok(Box::new(io::BufReader::new(GzDecoder::new(reader))))
    } else {
        Ok(file)
    }
}

pub fn count(mut reader: impl BufRead) -> MyResult<FileInfo> {
    let [mut num_lines, mut num_words, mut num_bytes, mut num_chars] = [0, 0, 0, 0];
    let mut line = String::new();

    loop {
        let line_bytes = reader.read_line(&mut line)?;
        if line_bytes == 0 {
            break;
        }

        num_bytes += line_bytes;
        num_lines += 1;
        num_words += line.split_whitespace().count();
        num_chars += line.chars().count();
        line.clear();
    }

    Ok(FileInfo { num_lines, num_words, num_bytes, num_chars })
}

#[cfg(test)]
mod tests {
    use super::{count, format_field, FileInfo};
    use std::io;

    #[test]
    fn test_count() {
        let text = "I don't want the world. I just want your half.\r\n";
        let info = count(io::Cursor::new(text));
        assert!(info.is_ok());

        let expected = FileInfo { num_lines: 1, num_words: 10, num_chars: 48, num_bytes: 48 };
        assert_eq!(info.unwrap(), expected);
    }

    #[test]
    fn test_format_field() {
        assert_eq!(format_field(1, false), "");
        assert_eq!(format_field(3, true), "       3");
        assert_eq!(format_field(10, true), "      10");
    }
}

#[allow(dead_code)]
fn format_field(value: usize, show: bool) -> String {
    return if show { format!("{:>8}", value) } else { "".to_string() };
}
