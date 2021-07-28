fn main() {
    let a = 40;

    unsafe {
        let b = &a as *const i32 as *mut i32;
        *b = 42;
    }

    println!("{}", a);
}
