mod skip_list;

fn main() {
    let mut list = skip_list::BestTransactionLog::new_empty(20);
    for i in 0..20 {
        list.append(i, format!("INSERT INTO mytable VALUES ({})", i).to_owned());
    }

    println!("{:?}", list);
    println!("{:?}", list.find(10));
}
