#![allow(dead_code)]

use std::any;
use std::cell::RefCell;
use std::rc::Rc;

#[derive(Debug)]
struct Node {
    value: u64,
    // parent: Tree,
    left: Tree,
    right: Tree,
}

type Tree = Option<Rc<RefCell<Node>>>;

#[derive(Debug)]
struct BinaryTree {
    root: Node,
    size: u64,
}

fn type_name_of<T>(_: &T) -> &str {
    any::type_name::<T>()
}

fn new_tree(value: u64) -> Tree {
    let node = Node::new(value);
    Some(Rc::new(RefCell::new(node)))
}

impl Node {
    fn new(value: u64) -> Node {
        Node { value, left: None, right: None }
    }

    fn add(&mut self, value: u64) {
        if value <= self.value {
            if let Some(node) = self.left.take() {
                println!("    <== walk left ({}, {})", self.value, node.borrow().value);
                // alloc::rc::Rc<core::cell::RefCell<binary_tree::Node>>
                // println!("~~~ node {:?}", type_name_of(&node));
                // let x = node.borrow();
                // println!("x {:?}", type_name_of(&x));
                (*node).borrow_mut().add(value); // !!! not *node.borrow_mut().add(value)
                self.left = Some(node); // must return self.left
            } else {
                println!("    <++ new left {}.left = {}\n", self.value, value);
                let node = Node::new(value);
                self.left = Some(Rc::new(RefCell::new(node)));
                // println!("{} {:?}", self.value, self.left);
            }
        } else {
            if let Some(node) = self.right.take() {
                println!("    ==> walk right ({}, {})", self.value, node.borrow().value);
                (*node).borrow_mut().add(value);
                self.right = Some(node); // must return to self.right
            } else {
                println!("    ++> add right {}.right = {}\n", self.value, value);
                self.right = new_tree(value);
            }
        }
    }
}

impl BinaryTree {
    fn new(value: u64) -> BinaryTree {
        BinaryTree { root: Node::new(value), size: 1 }
    }
    // left.borrow_mut().add(value);
    fn add(&mut self, value: u64) -> &mut Self {
        self.root.add(value);
        self.size += 1;
        self
    }
}

fn main() {
    let mut bt = BinaryTree::new(10);
    println!("{:?}", bt);

    println!("~~~");
    bt.add(5).add(1);

    println!("~~~");
    bt.add(12);

    println!("~~~");
    bt.add(4).add(6);

    println!("{:?}", bt.root);
}
