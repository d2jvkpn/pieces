fn main() {
    println!("{}", power_of_two(2));
    println!("{}", power_of_two(3));
}

fn power_of_two(v: usize) -> bool {
    return v & (v - 1) == 0;
}
