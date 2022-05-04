fn main() {
    println!("Hello, world!");
    println!("1/42 = {:.8}, 1/24 = {:.8}", 1.0 / 42.0, 1.0 / 24.0);

    println!(
        "Life, the Universe and Everything: {}, {}, {}",
        0b101010, 0x2a, '*' as u8,
    )
    // ?? handle big number
}
