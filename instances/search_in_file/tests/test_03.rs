fn main() {
    let mut buffer = [0; 16];
    buffer[0] = 65;

    println!("{}", buffer.iter().filter(|&v| *v != 0).count());
    println!("{:?}", String::from_utf8_lossy(&buffer).trim_matches(char::from(0)));
}
