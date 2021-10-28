#![allow(dead_code)]
#![allow(unused_variables)]

fn main() {
    let mut s = Box::new(42);

    replace_with_02(&mut s);
    println!("{}", s);

    replace_with_03(&mut s);
    println!("{}", s);
}

//fn replace_with_01(s: &mut Box<i32>) {
//    let was = *s; // not ok
//}

fn replace_with_02(s: &mut Box<i32>) {
    let was = std::mem::take(s); // replace with default value

    // *s = was;
}

fn replace_with_03(s: &mut Box<i32>) {
    let mut r = Box::new(84);
    std::mem::swap(s, &mut r);
    assert_ne!(*r, 84);
}
