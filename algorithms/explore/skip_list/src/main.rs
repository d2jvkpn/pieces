mod skip_list;

fn main() {
    let mut list = skip_list::BestTransactionLog::new_empty(5);
    for i in 0..30 {
        list.append(i, format!("INSERT INTO mytable VALUES ({})", i).to_owned());
    }

    println!("{:?}", list);
    println!("{}", list.describe());
    println!("{:?}", list.find(10));
}
