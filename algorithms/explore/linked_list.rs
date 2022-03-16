#![allow(dead_code)]

use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug, Clone)]
struct Node {
    value: i32,
    next: Link,
    prev: Link,
}

type Link = Option<Rc<RefCell<Node>>>;

impl Node {
    fn new(value: i32) -> Rc<RefCell<Node>> {
        Rc::new(RefCell::new(Node { value: value, next: None, prev: None }))
    }
}

#[derive(Debug, Clone)]
pub struct List {
    head: Link,
    tail: Link,
    pub length: u64,
}

impl List {
    fn new() -> List {
        List { head: None, tail: None, length: 0 }
    }

    fn append(&mut self, v: i32) -> &mut List {
        let new = Node::new(v);

        match self.tail.take() {
            Some(v) => {
                v.borrow_mut().next = Some(new.clone());
                new.borrow_mut().prev = Some(v);
            }
            None => {
                // new.borrow_mut().next = Some(new.clone());
                // new.borrow_mut().prev = Some(new.clone());
                self.head = Some(new.clone());
            }
        }

        self.length += 1;
        self.tail = Some(new);
        self
    }
}

pub struct ListIterator {
    current: Link,
}

impl ListIterator {
    fn new(start_at: Link) -> ListIterator {
        ListIterator { current: start_at }
    }
}

impl Iterator for ListIterator {
    type Item = i32;

    fn next(&mut self) -> Option<Self::Item> {
        let current = &self.current;
        let mut result = None;

        self.current = match current {
            Some(ref current) => {
                let current = current.borrow();
                result = Some(current.value.clone());
                current.next.clone()
            }
            None => None,
        };

        result
    }
}

impl DoubleEndedIterator for ListIterator {
    fn next_back(&mut self) -> Option<i32> {
        let current = &self.current;
        let mut result = None;

        self.current = match current {
            Some(ref current) => {
                let current = current.borrow();
                result = Some(current.value.clone());
                current.prev.clone()
            }
            None => None,
        };
        result
    }
}

fn main() {
    let mut list = List::new();
    list.append(1).append(2).append(3).append(4).append(5).append(6).append(7);
    // println!("{:?}", list.head.unwrap().borrow().value);

    let xx = ListIterator::new(list.head);

    //    for i in xx.take(3) {
    //        println!("> {}", i);
    //    }

    // println!("{:?}", xx.map(|v| v*2).collect::<Vec<i32>>());
    // println!("{:?}", xx.collect::<Vec<i32>>());

    // println!("{:?}", xx.reduce(|x: i32, y: i32| x * y));
    println!("{:?}", xx.rev().collect::<Vec<i32>>());
}
