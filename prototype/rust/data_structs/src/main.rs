// mod sll;
// use sll::SinglyLinkedList;

// declared in lib.rs
// use data_structs::SinglyLinkedList;

fn main() {
    println!("Hello, world!");
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn len() {
        let mut list = SinglyLinkedList::new("Alice");

        list.head.append("Bob").append("Rover");

        assert_eq!(list.len(), 3);
    }
}
