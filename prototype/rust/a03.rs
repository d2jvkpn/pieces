#[derive(Debug)]
struct Person {
    name: String,
}

#[derive(Debug)]
struct Dog<'a> {
    name: String,
    owner: &'a Person,
}

fn main() {
    let p1 = Person {
        name: String::from("John"),
    };

    let mut d1 = Dog {
        name: String::from("Max"),
        owner: &p1,
    };
    println!("{:?}, {:?}", d1, p1);

    //
    let p2 = Person {
        name: String::from("unknown"),
    };

    d1.owner = &p2;
    println!("{:?}, {:?}", d1, p2);
}

/* !!! returns a value referencing data owned by the current function
fn newD<'a>() -> Dog<'a> {
    let p1 = Person {
        name: String::from("John"),
    };

    Dog {
        name: String::from("Max"),
        owner: &p1,
    }
}
*/
