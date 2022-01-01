use std::io::Write;

use chrono::{Local, SecondsFormat};
use env_logger::Builder;
use log::{self, debug, error, info, log_enabled, warn, Level, LevelFilter};

// $ RUST_LOG=info cargo run
fn main() {
    // env_logger::init();
    Builder::new()
        .format(|buf, record| {
            writeln!(
                buf,
                "{} {:<5} {}",
                // Local::now().format("%Y-%m-%dT%H:%M:%S%:z"),
                // Local::now().to_rfc3339(),
                Local::now().to_rfc3339_opts(SecondsFormat::Millis, true),
                record.level(),
                record.args()
            )
        })
        .filter(None, LevelFilter::Info)
        .init();

    let user = rs_log::User::new("Evol", "klamsasaksmas");
    user.sign_in("123456");

    debug!("Mary has a little lamb");
    error!("{}", "Its fleece was white as snow");
    info!("{:?}", "And every where that Mary went");
    warn!("{:#?}", "The lamb was sure to go");

    debug!("this is a debug {}", "message");
    error!("this is printed by default");

    if log_enabled!(Level::Info) {
        let x = 3 * 4; // expensive computation
        info!("the answer was: {}", x);
    }
}
