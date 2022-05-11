// Rust Web Programing 2021, Chapter03
use std::{thread, time};

use async_std::task;
use chrono::{Local, SecondsFormat};
use futures::{executor::block_on, future::join_all, join};

fn now_string() -> String {
    Local::now().to_rfc3339_opts(SecondsFormat::Millis, true)
}

fn run1(number: i8) -> i8 {
    //  DateTime<Local>
    // let now = || chrono::Local::now().to_rfc3339_opts(chrono::SecondsFormat::Millis, true);

    println!(">>> {} Run1 number {} is running", now_string(), number);
    let two_seconds = time::Duration::new(2, 0);
    thread::sleep(two_seconds);
    println!("    {} Run1 number {} is done", now_string(), number);
    return 2;
}

async fn run1_async(number: i8) -> i8 {
    run1(number)
}

async fn run2(number: i8) -> i8 {
    println!(">>> {} Run2 number {} is running", now_string(), number);
    let two_seconds = time::Duration::new(2, 0);
    task::sleep(two_seconds).await;
    println!("    {} Run2 number {} is done", now_string(), number);
    return 2;
}

fn main() {
    // ####
    // defines a future
    let now = time::Instant::now();
    println!("### {} future_1 is created", now_string());
    let future_1 = run1_async(1);
    // holds the program for the result of the first future
    let outcome = block_on(future_1);
    println!("~~~ time elapsed {:?}, result: {}\n", now.elapsed(), outcome);

    let now = time::Instant::now();
    // defines the async block for multiple futures (just like an async function)
    let batch2 = async {
        // defines two futures
        let future_2 = run1_async(2);
        let future_3 = run1_async(3);
        let future_4 = run1_async(4);
        // waits for both futures to complete in sequence
        return join!(future_2, future_3, future_4);
    };
    // holds the program for the result from the async block
    let results = block_on(batch2);
    println!("~~~ time elapsed {:?}, results: {:?}\n", now.elapsed(), results);

    // ####
    let now = time::Instant::now();
    // defines the async block for multiple futures (just like an async function)
    let batch3 = async {
        // defines two futures
        let future_5 = run2(5);
        let future_6 = run2(6);
        let future_7 = run2(7);
        // waits for both futures to complete in sequence
        return join!(future_5, future_6, future_7);
    };
    // holds the program for the result from the async block
    let results = block_on(batch3);
    println!("~~~ time elapsed {:?}, results: {:?}\n", now.elapsed(), results);

    // ####
    let now = time::Instant::now();
    let batch4 = async {
        let futures_vec = vec![run1_async(8), run1_async(9), run1_async(10)];

        // applies the spawn async tasks for all futures and collect them into a vector
        let handles: Vec<task::JoinHandle<i8>> = futures_vec.into_iter().map(task::spawn).collect();
        // futures_vec.into_iter().map(task::spawn).collect::<Vec<_>>();

        let results = join_all(handles).await;
        return results;
    };
    let results = block_on(batch4);
    println!("~~~ time elapsed {:?}, results: {:?}\n", now.elapsed(), results);

    // ####
    // start the timer again
    let now = time::Instant::now();
    // spawn a few functions with the same function
    let batch5: Vec<thread::JoinHandle<i8>> =
        vec![11, 12, 13].into_iter().map(|v| thread::spawn(move || run1(v))).collect();
    // vec![thread::spawn(|| run(12)), thread::spawn(|| run(13)), thread::spawn(|| run(14))];

    let results: Vec<i8> = batch5.into_iter().map(|t| t.join().unwrap()).collect();
    // print the outcomes again from the threads
    println!("~~~~ time elapsed {:?}, results: {:?}\n", now.elapsed(), results);
}
