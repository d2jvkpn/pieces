type Tree = Option<Box<Node>>;

#[derive(Debug)]
struct Node {
    pub value: u64,
    left: Tree,
    right: Tree,
}

#[derive(Debug)]
pub struct BinarySearchTree {
    root: Tree,
    pub length: u64,
}

impl Node {
    fn new(value: u64) -> Self {
        Node { value, left: None, right: None }
    }
}

impl BinarySearchTree {
    fn new(value: u64) -> Self {
        BinarySearchTree { root: Some(Box::new(Node::new(value))), length: 1 }
    }

    fn add(&self, value: u64) {}
}

fn main() {
    let bst = BinarySearchTree::new(1);
    println!("{:?}", bst);
}
