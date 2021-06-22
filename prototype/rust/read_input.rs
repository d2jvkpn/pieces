use std::io;
use std::thread;
use std::time;

fn main() {
    let mut hunched = String::new();

    println!("Enter:");

    io::stdin()
        .read_line(&mut hunched) // return io::Result
        .ok()
        .expect("Line reading failed");

    println!("{}", hunched);
    thread::sleep(time::Duration::new(1, 0));
    println!("done");

    let arr = vec![5, 1, 4];
    let _y = &arr[1];
    println!("arr: {:?}", arr);

    for (i, v) in arr.iter().enumerate() {
        println!("arr[{}]: {}", i, v);
    }
}
