// use std::str;

fn main() {
    let mut cache: Vec<char> = Vec::with_capacity(16); // Vev<u8>

    cache.push('a');
    cache.push('b');
    cache.push('c');

    println!("cache={:?}, len={}, cap={}", &cache, cache.len(), cache.capacity());

    let slice = vec!['d', 'f', 'g'];
    //    cache = cache[..0].to_vec();
    //    println!("cache={:?}, len={}, cap={}", &cache, cache.len(), cache.capacity()); // cache=[], len=0, cap=0

    cache.clear();
    cache.extend_from_slice(&slice);

    println!("cache={:?}, len={}, cap={}", &cache, cache.len(), cache.capacity());
}
