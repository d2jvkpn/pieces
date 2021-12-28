fn main() {
    let mut buffer = [0; 10];
    buffer[0] = 65;

    println!(
        "notzeros={}, buffer={:?}",
        buffer.iter().filter(|&v| *v != 0).count(),
        String::from_utf8_lossy(&buffer).trim_matches(char::from(0)),
    );

    let slice = &mut buffer[5..];
    println!("slice.len()={}", slice.len());
    slice[0] = 65;

    println!(
        "notzeros={}, buffer={:?}",
        buffer.iter().filter(|&v| *v != 0).count(),
        String::from_utf8_lossy(&buffer).trim_matches(char::from(0)),
    );
}
