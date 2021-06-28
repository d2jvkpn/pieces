fn main() {
    let num: f32 = 42.42;
    let n: u32 = unsafe { std::mem::transmute(num) };
    println!("{:b}", n);
}
