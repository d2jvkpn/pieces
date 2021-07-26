#![allow(dead_code)]

use std::mem;

fn main() {
    let d1 = new_d1();
    println!(">>> main d1: {:p}", &d1);
    println!("memory size of d1: {}", mem::size_of_val(&d1));

    let d2 = d1;
    println!(">>> main d2: {:p}", &d2);
}

struct D {
    v: i64,
}

fn new_d1() -> D {
    let d = D { v: 1 };
    println!(">>> new_d1: {:p}", &d);
    return d;
}
