fn main() {
    let a = 10;
    let mut b = &a;
    let mut c = &mut b;
    println!("{}", c);

    let b = 9;
    let mut d = &b;
    c = &mut d;
    println!("{}", c);
}
