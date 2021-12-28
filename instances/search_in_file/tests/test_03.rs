fn main() {
    let mut buffer = [0; 10];
    buffer[0] = 65;
    println!(
        "buffer notzeros={}, buffer={:?}",
        buffer.iter().filter(|&v| *v != 0).count(),
        String::from_utf8_lossy(&buffer).trim_matches(char::from(0)),
    );

    let slice = &mut buffer[5..];
    println!("slice.len()={}", slice.len());
    slice[0] = 65;

    println!(
        "buffer notzeros={}, buffer={:?}",
        buffer.iter().filter(|&v| *v != 0).count(),
        String::from_utf8_lossy(&buffer).trim_matches(char::from(0)),
    );

    let v = &mut buffer.to_vec(); // clone of buffer
    v[1] = 66;
    // v=[65, 66, 0, 0, 0, 65, 0, 0, 0, 0], buffer[1]=0
    println!("v={:?}, buffer[1]={}", v, buffer[1]);

    println!(
        "buffer notzeros={}, buffer={:?}",
        buffer.iter().filter(|&v| *v != 0).count(),
        String::from_utf8_lossy(&buffer).trim_matches(char::from(0)),
    );
}
