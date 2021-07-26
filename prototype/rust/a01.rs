#![allow(dead_code)]

use std::mem;

fn main() {
    let d1 = new_d1();
    println!(">>> main d1: {:p}", &d1);

    let d2 = d1;
    println!(">>> main d2: {:p}", &d2); // != &d1

    let d3 = d2;
    println!(">>> main d3: {:p}", &d3); // != &d2

    let d4 = &d3;
    let d5 = &d3;
    println!(">>> main d4: {:p}, d5: {:p}", d4, d5); // ==
}

struct D {
    v: i64,
}

fn new_d1() -> D {
    let d = D { v: 1 };
    println!(">>> new_d1: {:p}, size={}", &d, mem::size_of_val(&d));
    return d;
}
