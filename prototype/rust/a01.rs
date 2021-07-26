fn main() {
    let d = new_d1();

    println!(">>> main: {:p}", &d);
}

struct D {
    v: i64,
}

fn new_d1() -> D {
    let d = D { v: 1 };

    println!(">>> new_d1: {:p}", &d);
    return d;
}
