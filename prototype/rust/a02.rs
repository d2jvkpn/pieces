#![allow(dead_code)]

use std::mem;

fn main() {
    let d1 = new_d1();
    println!(">>> main d1: {:p}", &d1);

    let d2 = d1; // moved
    println!(">>> main d2: {:p}", &d2);

    let d3 = &d2;
    println!(">>> main d3: {:p}, {:p}", d3, &d3);

    let d4 = &d2;
    println!(">>> main d4: {:p}, {:p}", d4, &d4);

    println!("{:?}, {:?}", d3, d4);
}

#[derive(Debug)]
struct D {
    v: i64,
}

fn new_d1() -> D {
    let d = D { v: 1 };
    println!(">>> new_d1: {:p}, size={}", &d, mem::size_of_val(&d));
    return d;
}

/* >>> output
>>> new_d1: 0x7ffddf2d9f90, size=8
>>> main d1: 0x7ffddf2da0a8
>>> main d2: 0x7ffddf2da100
>>> main d3: 0x7ffddf2da100, 0x7ffddf2da158
>>> main d4: 0x7ffddf2da100, 0x7ffddf2da1c8
D { v: 1 }, D { v: 1 }
*/
