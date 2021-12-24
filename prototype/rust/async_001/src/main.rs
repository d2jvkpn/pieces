use std::{thread, time};

use async_std::task;
use chrono;
use futures::{executor, future, join};
use tokio;

fn main() {
    println!(">>> future01");
    let f01 = xx01();
    executor::block_on(f01);
    /*
    >>> future01
    job1 start
    job1 end
    job2 start
    job2 end
    xx01: r01 = Ok("hello"), r02 = Ok("hello")
    */

    println!(">>> future02");
    let f02 = xx02();
    executor::block_on(f02);
    /*
    >>> future02
    job3 start
    job4 start
    job3 end
    job4 end
    xx02: r01 = "hello", r02 = "hello"
    */

    println!(">>> future03");
    // !! not a future
    xx03();
    /*
    >>> future03
    job5 start
    job6 start
    job5 end
    job6 end
    xx03: r01 = "hello", r02 = "hello"
    */
}

fn now() -> String {
    let now = chrono::Local::now();
    now.to_rfc3339_opts(chrono::SecondsFormat::Millis, true)
}

fn a01(name: &str) -> String {
    println!("{} {} start", now(), name);
    thread::sleep(time::Duration::new(5, 0));
    println!("{} {} end", now(), name);
    String::from("hello")
}

async fn a02(name: &str) -> String {
    println!("{} {} start", now(), name);
    task::sleep(time::Duration::new(5, 0)).await;
    println!("{} {} end", now(), name);
    String::from("hello")
}

async fn a03(name: &str) -> String {
    println!("{} {} start", now(), name);
    tokio::time::sleep(time::Duration::new(5, 0)).await;
    println!("{} {} end", now(), name);
    String::from("hello")
}

async fn xx01() {
    // Here I specify the type of the error as (); otherwise the compiler can't infer it
    // wrapped as a async function
    let future01 = future::ok::<String, ()>(a01("job1"));
    let future02 = future::ok::<String, ()>(a01("job2"));
    // let a = future.await;
    // println!("{:?}", a);
    let (r01, r02) = join!(future01, future02);
    // execute in sequential
    println!("{} xx01: r01 = {:?}, r02 = {:?}", now(), r01, r02);
}

async fn xx02() {
    let future01 = a02("job3");
    let future02 = a02("job4");
    let (r01, r02) = join!(future01, future02);
    println!("{} xx02: r01 = {:?}, r02 = {:?}", now(), r01, r02);
}

#[tokio::main]
async fn xx03() {
    let future01 = a03("job5");
    let future02 = a03("job6");
    let (r01, r02) = future::join(future01, future02).await;
    println!("{} xx03: r01 = {:?}, r02 = {:?}", now(), r01, r02);
}
