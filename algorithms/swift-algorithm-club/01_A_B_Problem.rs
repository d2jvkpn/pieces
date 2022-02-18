/**
 *A + B Problem
 */

fn main() {
    println!("{} + {} = {}", 100, 2, get_sum(100, 2));

    println!("({} + {})/2 = {}", 3, 5, mean(3, 5));

    println!("{}, {}, {}, {}", 1 | 2, 1 & 2, 2 << 1, 1 ^ 2);
}

fn get_sum(mut a: i64, mut b: i64) -> i64 {
    while a != 0 {
        a = (a & b) << 1;
        b = a ^ b; // AND, Left Shift, XOR
    }

    return b;
}

////
fn mean(a: i64, b: i64) -> i64 {
    return a >> 1 + b >> 1 + a & b & 1;
}
