#![allow(dead_code)]

use std::{error, fs};

use assert_cmd::Command;
use predicates::prelude::*;
use rand::{distributions::Alphanumeric, Rng};

type TestResult = Result<(), Box<dyn error::Error>>;

const PRG: &str = "catr";
const EMPTY: &str = "tests/inputs/empty.txt";
const FOX: &str = "tests/inputs/fox.txt";
const SPIDERS: &str = "tests/inputs/spiders.txt";
const BUSTLE: &str = "tests/inputs/the-bustle.txt";

fn gen_bad_file() -> String {
    loop {
        let filename: String =
            rand::thread_rng().sample_iter(&Alphanumeric).take(7).map(char::from).collect();

        if fs::metadata(&filename).is_err() {
            return filename;
        }
    }
}

fn run(args: &[&str], expected_file: &str) -> TestResult {
    let expected = fs::read_to_string(expected_file)?;
    Command::cargo_bin(PRG)?.args(args).assert().success().stdout(expected);
    Ok(())
}

fn run_stdin(input_file: &str, args: &[&str], expected_file: &str) -> TestResult {
    let input = fs::read_to_string(input_file)?;
    let expected = fs::read_to_string(expected_file)?;
    Command::cargo_bin(PRG)?.args(args).write_stdin(input).assert().success().stdout(expected);
    Ok(())
}

#[test]
fn skips_bad_file() -> TestResult {
    let bad = gen_bad_file();
    let expected = format!("{}: .* [(]os error 2[)]", bad);

    Command::cargo_bin(PRG)?
        .arg(&bad)
        .assert()
        .success()
        .stderr(predicate::str::is_match(expected)?);

    Ok(())
}

#[test]
fn bustle() -> TestResult {
    run(&[BUSTLE], "tests/inputs/the-bustle.txt")
}

#[test]
fn bustle_stdin() -> TestResult {
    run_stdin(BUSTLE, &["-"], "tests/inputs/the-bustle.txt")
}
