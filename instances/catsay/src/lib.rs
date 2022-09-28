/* bash
  $ cargo add --dev assert_cmd predicates

  $ cargo test
  $ cargo test tests::run_with_defaults -- --exact
  $ cargo test run_with_defaults
*/

#[cfg(test)]
mod tests {
    use std::{error, process};

    use assert_cmd::prelude::*;
    use predicates::prelude::*;

    #[test]
    fn run_with_defaults() -> Result<(), Box<dyn error::Error>> {
        process::Command::cargo_bin("catsay")
            .expect("binary exists")
            .assert()
            .success()
            .stdout(predicate::str::contains("Meow!"));

        Ok(())
    }

    #[test]
    fn fail_on_non_existing_file() -> Result<(), Box<dyn error::Error>> {
        process::Command::cargo_bin("catsay")
            .expect("binary exists")
            .args(&["-f", "no/such/file.txt"])
            .assert()
            .failure();

        Ok(())
    }
}
