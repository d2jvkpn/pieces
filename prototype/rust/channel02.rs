use std::sync::mpsc;
use std::thread;
// use std::time::Duration;

fn main() {
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let vals = vec![
            String::from("hi"),
            String::from("from"),
            String::from("the"),
            String::from("thread"),
        ];

        for val in vals {
            println!("-> send {}", val.clone());
            tx.send(val).unwrap();
            // thread::sleep(Duration::from_secs(1));
        }
    });

    for received in rx {
        println!("<- got: {}", received);
    }

    /*
    let received = rx.recv().unwrap();
    println!("Got: {}", received);
    */
}
