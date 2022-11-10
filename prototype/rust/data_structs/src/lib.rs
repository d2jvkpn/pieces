pub mod doubly_liked_list;
pub mod singly_linked_list;

pub use doubly_liked_list::DoublyLikedList;
pub use singly_linked_list::SinglyLinkedList;

#[cfg(test)]
mod tests {
    use crate::{DoublyLikedList, SinglyLinkedList};

    #[test]
    fn singly_linked_list() {
        let mut list = SinglyLinkedList::new("Alice");

        list.head.append("Bob").append("Rover");
        let mut count = 0;

        let mut item = Some(&list.head);
        while let Some(elem) = item {
            count += 1;
            println!("{:?}", elem.data);
            item = elem.next();
        }

        assert_eq!(count, 3);
    }

    #[test]
    fn doubly_linked_list() {
        let mut list = DoublyLikedList::new("Alice");
        list.append("Bob");
        list.append("Rover");

        println!("{:?}", list.tail().borrow().data);
    }
}
