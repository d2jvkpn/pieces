#![allow(dead_code)]

use std::cell::RefCell;
use std::rc::Rc;

pub struct ListItem<T> {
    pub(crate) prev: Option<ItemRef<T>>,
    pub(crate) data: Box<T>,
    pub(crate) next: Option<ItemRef<T>>,
}

pub type ItemRef<T> = Rc<RefCell<ListItem<T>>>;

pub struct DoublyLikedList<T> {
    pub(crate) head: ItemRef<T>,
}

impl<T> ListItem<T> {
    pub fn new(data: T) -> Self {
        Self { prev: None, data: Box::new(data), next: None }
    }

    pub fn item_ref(data: T) -> ItemRef<T> {
        Rc::new(RefCell::new(ListItem::new(data)))
    }
}

impl<T> DoublyLikedList<T> {
    pub fn new(data: T) -> Self {
        Self { head: ListItem::item_ref(data) }
    }

    pub fn append(&mut self, data: T) {
        let tail = Self::find_tail(self.head.clone());
        let item = ListItem::item_ref(data);

        item.borrow_mut().prev = Some(tail.clone());
        tail.borrow_mut().next = Some(item);
    }

    pub fn find_tail(item: ItemRef<T>) -> ItemRef<T> {
        if let Some(next) = &item.borrow().next {
            Self::find_tail(next.clone())
        } else {
            item.clone()
        }
    }

    pub fn head(&self) -> ItemRef<T> {
        self.head.clone()
    }

    pub fn tail(&self) -> ItemRef<T> {
        Self::find_tail(self.head.clone())
    }
}
