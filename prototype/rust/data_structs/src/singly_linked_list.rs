#![allow(dead_code)]

#[derive(Debug)]
pub struct ListItem<T> {
    pub data: Box<T>,
    pub next: Option<Box<ListItem<T>>>,
}

pub struct SinglyLinkedList<T> {
    pub head: ListItem<T>,
}

impl<T> ListItem<T> {
    pub fn new(data: T) -> Self {
        Self { data: Box::new(data), next: None }
    }

    pub fn next(&self) -> Option<&Self> {
        if let Some(item) = &self.next {
            Some(&*item)
        } else {
            None
        }
    }

    pub fn mut_tail(&mut self) -> &mut Self {
        if self.next.is_some() {
            self.next.as_mut().unwrap().mut_tail()
        } else {
            self
        }
    }

    pub fn append(&mut self, data: T) -> &mut Self {
        self.next = Some(Box::new(ListItem::new(data)));
        self.next.as_mut().unwrap()
    }
}

impl<T> SinglyLinkedList<T> {
    pub fn new(data: T) -> Self {
        Self { head: ListItem::new(data) }
    }

    pub fn len(&self) -> usize {
        let (mut count, mut curr) = (1, &self.head);

        while let Some(ref item) = curr.next {
            count += 1;
            curr = item;
        }
        count
    }

    //    pub fn append(&mut self, data: T) -> &mut ListItem<T> {
    //        let mut tail = self.head.mut_tail();
    //        tail.next = Some(Box::new(ListItem::new(data)));
    //        tail.next.as_mut().unwrap()
    //    }
}
