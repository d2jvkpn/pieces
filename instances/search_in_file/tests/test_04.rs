fn main() {
    let h1 = "Hello";
    println!("{:?}", h1.as_bytes());

    let h2 = "你好";
    println!("{:?}", h2.as_bytes());

    let mut v1 = vec![0, 1, 2, 3, 4, 5];
    // dbg!(&v1);
    println!(">>> v1 = {:?}", v1);
    println!("    v1.len() = {}, v1.capacity()={}", v1.len(), v1.capacity()); // 6, 6

    v1.drain(0..2); // drop the first two elements
    println!(">>> v1 = {:?}", v1);
    println!("    v1.len() = {}, v1.capacity()={}", v1.len(), v1.capacity()); // 4, 6

    v1.truncate(2); // keep the first two elements
    println!(">>> v1 = {:?}", v1);
    println!("    v1.len() = {}, v1.capacity()={}", v1.len(), v1.capacity()); // 2, 6

    v1.extend_from_slice(&vec![6, 7, 8, 9, 10, 11]);
    println!(">>> v1 = {:?}", v1);
    println!("    v1.len() = {}, v1.capacity()={}", v1.len(), v1.capacity()); // 8, 12
}
