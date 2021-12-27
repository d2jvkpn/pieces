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
    let (mut index, mut t, mut tag) = (0, 2 * bts.len(), 0_i8);
    let mut cache: Vec<u8> = Vec::with_capacity(if t > SIZE { t } else { SIZE });
    let mut reader = io::BufReader::new(read);
    let mut buffer = [0; SIZE];

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
            t = cache.len() - SIZE; // left shift
            index += t as i64;
            let tail = cache[t..].to_vec();
            cache.clear();
            cache.extend_from_slice(&tail);
        }

        if debug {
            eprintln!("~~~ fill [{}:{}]: index={}", cache.len(), cache.len() + SIZE, index);
        }

        buffer.iter_mut().for_each(|x| *x = 0);
        match reader.read_exact(&mut buffer) {
            Ok(_) => {}
            Err(ref e) if e.kind() == io::ErrorKind::UnexpectedEof => tag = -1,
            Err(e) => return Err(e),
        };

        // let slice = buffer.to_vec().into_iter().filter(|v| *v != 0).collect::<Vec<_>>();
        // println!("{:?}", String::from_utf8_lossy(&slice).trim_matches(char::from(0)));
        // cache.extend_from_slice(&slice);
        t = buffer.iter().filter(|&v| *v != 0).count();
        cache.extend_from_slice(&buffer[..t]);

        if debug {
            eprintln!(
                "    slice.len()={}, cache.len()={}\n    cache={:?}",
                t, // slice.len()
                cache.len(),
                String::from_utf8_lossy(&cache)
            );
        }

        if let Some(s) = find_subseq(&cache, bts) {
            index += s as i64;
            tag = 1;
        }
    }

    if debug {
        // eprintln!("<<< cache={:?}", str::from_utf8(&cache));
        eprintln!("<<< cache={:?}", String::from_utf8_lossy(&cache[..]));
    }

    return if tag == 1 { Ok(index) } else { Ok(-1) };
}

fn find_subseq<T>(slice: &[T], sub: &[T]) -> Option<usize>
where
    for<'a> &'a [T]: PartialEq,
{
    slice.windows(sub.len()).position(|window| window == sub)
}
