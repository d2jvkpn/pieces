use chrono::prelude::{DateTime, Local, SecondsFormat};
use clap::{App, Arg};
use rand::{distributions::Alphanumeric, thread_rng, Rng}; // Values
use serde::{Deserialize, Serialize};
use toml;

use std::{fmt, fs};

#[derive(Default, Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "camelCase")]
pub struct Account {
    pub site: String,
    pub username: String,
    pub password: String,
    pub time: String,
}

// pub fn run() -> Result<Account, Box<dyn error::Error>> {
pub fn run() -> Result<Account, String> {
    let args = App::new(env!("CARGO_PKG_HOMEPAGE"))
        .about(env!("CARGO_PKG_DESCRIPTION"))
        .author(env!("CARGO_PKG_AUTHORS"))
        .version(env!("CARGO_PKG_VERSION"))
        .set_term_width(100)
        .arg(
            Arg::with_name("site")
                .long("site")
                .short("s")
                .takes_value(true)
                .required(false)
                .default_value("")
                .help("set account site url"),
        )
        .arg(
            Arg::with_name("prefix")
                .long("prefix")
                .short("p")
                .takes_value(true)
                .required(false)
                .default_value("")
                .help("set account username prefix"),
        )
        .arg(
            Arg::with_name("length")
                .long("length")
                .short("l")
                .takes_value(true)
                .required(false)
                .default_value("24")
                .help("set account name length"),
        )
        .arg(
            Arg::with_name("password_length")
                .long("password_length")
                .short("L")
                .takes_value(true)
                .required(false)
                .default_value("32")
                .help("set account password length"),
        )
        .arg(
            Arg::with_name("save")
                .long("save")
                .takes_value(false)
                .required(false)
                .help("save account data to file"),
        )
        .get_matches();

    let site = args.value_of("site").unwrap_or("").to_string();
    let prefix = args.value_of("prefix").unwrap_or("").to_string();

    // let length = matches.value_of("length").unwrap_or("24").parse::<usize>()?;
    let length = match args.value_of("length").unwrap_or("24").parse::<usize>() {
        Ok(v) => v,
        // Err(e) => return Err(From::from(format!("parse arg --length error: {:?}", e))),
        Err(e) => return Err(format!("parse arg --length error: {:?}", e)),
    };

    // let password_len = matches.value_of("password_len").unwrap_or("32").parse::<usize>()?;
    let password_len = match args.value_of("password_length").unwrap_or("32").parse::<usize>() {
        Ok(v) => v,
        // Err(e) => return Err(From::from(format!("parse arg --password_len error: {:?}", e))),
        Err(e) => return Err(format!("parse arg --password_length error: {:?}", e)),
    };

    let save = args.is_present("save");

    ////
    let mut account = Account::new(prefix, length, password_len);
    account.site = site;

    // println!("{}", account.serialize("json"));
    // println!("{}", account.serialize("yaml"));
    // println!("{}", account.serialize("toml"));

    if !save {
        println!("{}", account);
        return Ok(account);
    }

    let output = "account_".to_owned() + &account.username + ".yaml";

    if let Err(err) = fs::write(output.clone(), format!("{}", account)) {
        // return Err(From::from(format!("unable to write file: {:?}", err)));
        return Err(format!("unable to write file: {:?}", err));
    }

    eprintln!("save to {}", output);
    Ok(account)
}

impl Account {
    pub fn new(prefix: String, length: usize, password_len: usize) -> Account {
        let time: DateTime<Local> = Local::now();
        let time = time.to_rfc3339_opts(SecondsFormat::Millis, true);

        let username = if length > prefix.len() {
            prefix.to_string()
                + &thread_rng()
                    .sample_iter(&Alphanumeric)
                    .take(length - prefix.len())
                    .map(char::from)
                    .collect::<String>()
        } else {
            prefix
        };

        // let password: String =
        //    thread_rng().sample_iter(&Alphanumeric).take(password_len).map(char::from).collect();

        const CHARSET: &[u8] =
            b"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789!@#$%^&*()_+-=[]{}|";

        // Minimum password length is 8 characters
        // Include a minimum of three of the following mix of character types: uppercase, lowercase,
        // numbers, and ! @ # $ % ^ & * ( ) _ + - = [ ] { } | '
        // ~ / " , ; < > ? \ `

        let mut rng = rand::thread_rng();

        let password: String = (0..password_len)
            .map(|_| {
                let idx = rng.gen_range(0..CHARSET.len());
                CHARSET[idx] as char
            })
            .collect();

        // Account { time: time, username: username, password: password }
        Account { username, password, time, ..Default::default() }
    }

    pub fn serialize(&self, t: &str) -> String {
        match t {
            "json" => serde_json::to_string(&self).unwrap_or(String::from("")),
            "yaml" => serde_yaml::to_string(&self).unwrap_or(String::from("")),
            "toml" => toml::to_string(&self).unwrap_or(String::from("")),
            _ => String::from(""),
        }
    }
}

impl fmt::Display for Account {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        // let site = if self.site.len() > 0 { self.site.clone() } else { "\"\"".to_string() };

        //        write!(
        //            f,
        //            "- site: {}\n  username: {}\n  password: {}\n  time: {}",
        //            site, self.username, self.password, self.time,
        //        )

        write!(f, "{}", serde_yaml::to_string(&vec![self]).unwrap_or(String::from("")))
    }
}
