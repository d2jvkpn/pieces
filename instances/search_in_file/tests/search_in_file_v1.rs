use std::{
    env,
    fs::File,
    io::{prelude::*, BufRead, BufReader, ErrorKind},
    str,
};

fn main() {
    let args: Vec<String> = env::args().collect();
    println!("{:?}", args);

    let target = "eee";
    let fp = "a01.txt";

    let file = File::open(fp).unwrap();

    let bts = target.as_bytes();
    const SIZE: usize = 32;
    let (mut t, mut index) = (0_usize, 0_usize);

    t = 2 * bts.len();
    let mut cache: Vec<u8> = Vec::with_capacity(if t > SIZE { t } else { SIZE });
    let mut reader = BufReader::new(file);

    loop {
        if cache.len() + SIZE > cache.capacity() {
            index += cache.len();
            cache.clear();
        }
        let mut ok = true;
        let mut buffer = [0; SIZE];

        match reader.read_exact(&mut buffer) {
            Ok(_) => {}
            Err(ref e) if e.kind() == ErrorKind::UnexpectedEof => ok = false,
            Err(e) => Err(e).unwrap(),
        };

        let slice = buffer.to_vec().into_iter().filter(|v| *v != 0).collect::<Vec<_>>();
        // println!("{:?}", String::from_utf8_lossy(&slice).trim_matches(char::from(0)));
        println!("{:?}", String::from_utf8_lossy(&slice));
        cache.extend_from_slice(&slice);
        if let Some(s) = find_subsequence(&cache, bts) {
            index += s;
            break;
        }
        if !ok {
            break;
        }
    }

    println!("{:?}", str::from_utf8(&cache).unwrap());
    println!("{}", index);
}

fn find_subsequence<T>(haystack: &[T], needle: &[T]) -> Option<usize>
where
    for<'a> &'a [T]: PartialEq,
{
    haystack.windows(needle.len()).position(|window| window == needle)
}
