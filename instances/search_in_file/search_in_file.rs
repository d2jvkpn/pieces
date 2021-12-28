use std::{
    env,
    fs::File,
    io::{self, Read},
    // str,
};

fn main() {
    let args: Vec<String> = env::args().collect();
    let (target, fp) = (&args[1], &args[2]);
    eprintln!("target={:?}, fp={}", target, fp);

    let file = File::open(fp).unwrap();
    let bts = target.as_bytes();
    let app_debug = env::var("APP_Debug").unwrap_or("".to_string()) == "true";

    match search_text(bts, file, app_debug) {
        Err(e) => Err(e).unwrap(),
        Ok(v) if v == -1 => println!("NotFound: -1"),
        Ok(v) => println!("Index: {}", v),
    }
}

// Err(e) => io::Error, Ok(-1) => NotFound, Ok(>=0) => Found
fn search_text(bts: &[u8], read: impl io::Read, debug: bool) -> Result<i64, io::Error> {
    const SIZE: usize = 32; // must to be a const for creating an array
    let (mut index, k, mut tag) = (0_i64, bts.len(), 0_i8);
    let mut cache: Vec<u8> = Vec::with_capacity(if 8 * k > 4 * SIZE { 8 * k } else { 4 * SIZE });
    let mut reader = io::BufReader::new(read);
    let mut buffer = [0; SIZE];
    // ?? can't BuffRead read to a vec as I can do in Golang

    if debug {
        eprintln!(
            ">>> bts.len()={}, SIZE={}, cache.len()={}, cache.capacity()={}",
            bts.len(),
            SIZE,
            cache.len(),
            cache.capacity(),
        );
    }

    while tag == 0 {
        if cache.len() + SIZE > cache.capacity() {
            index += cache.len() as i64;
            let tail = cache[(cache.len() - SIZE)..].to_vec(); // left shift
            cache.clear();
            cache.extend_from_slice(&tail);
        }

        if debug {
            eprintln!(
                "~~~ read to cache: [{}:{}], index={}",
                cache.len(),
                cache.len() + SIZE,
                index
            );
        }

        buffer.iter_mut().for_each(|x| *x = 0);
        match reader.read_exact(&mut buffer) {
            Ok(_) => {}
            // don't continue next loop
            Err(ref e) if e.kind() == io::ErrorKind::UnexpectedEof => tag = -1,
            Err(e) => return Err(e),
        };

        // let slice = buffer.to_vec().into_iter().filter(|v| *v != 0).collect::<Vec<_>>();
        // println!("{:?}", String::from_utf8_lossy(&slice).trim_matches(char::from(0)));
        // cache.extend_from_slice(&slice);
        let s = buffer.iter().filter(|&v| *v != 0).count();
        cache.extend_from_slice(&buffer[..s]);

        if debug {
            eprintln!(
                "    save to cache: [{}:{}]\n    cache={:?}",
                cache.len() - s,
                cache.len(),
                String::from_utf8_lossy(&cache),
                // eprintln!("<<< cache={:?}", str::from_utf8(&cache));
            );
        }

        let mut t = cache.len() as i64 - k as i64 - s as i64;
        if t < 0 {
            t = 0;
        }
        if let Some(s) = find_subseq(&cache[t as usize..], bts) {
            index += t + (s as i64);
            tag = 1; // don't continue next loop
        }
    }

    return if tag == 1 { Ok(index) } else { Ok(-1) };
}

fn find_subseq<T>(slice: &[T], sub: &[T]) -> Option<usize>
where
    for<'a> &'a [T]: PartialEq,
{
    slice.windows(sub.len()).position(|window| window == sub)
}
