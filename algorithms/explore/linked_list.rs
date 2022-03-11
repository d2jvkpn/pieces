use std::cell::RefCell;
use std::rc::Rc;

struct Node {
	value: i32,
	next: Option<Rc<RefCell<Node>>>
}
