use std::collections::HashMap;
use std::fmt;

#[derive(Debug)]
pub enum Value<'b> {
    Single(&'b str),
    Multiple(Vec<&'b str>),
}

////
// a=1&b=2&c&d=&e===&d=7&d=abc
pub struct QueryString<'b> {
    data: HashMap<&'b str, Value<'b>>,
}

impl<'b> QueryString<'b> {
    pub fn get(&self, key: &str) -> Option<&Value> {
        self.data.get(key);
        unimplemented!()
    }
}

impl<'b> fmt::Debug for QueryString<'b> {
    fn fmt(&self, f: &mut fmt::Formatter) -> fmt::Result {
        write!(f, "data: size= {:?}", self.data) //self.data.keys().len()
    }
}

impl<'b> From<&'b str> for QueryString<'b> {
    fn from(s: &'b str) -> Self {
        let mut data = HashMap::new();

        for sub_str in s.split('&') {
            let mut key = sub_str;
            let mut val = "";

            if let Some(i) = sub_str.find('=') {
                key = &sub_str[..i];
                val = &sub_str[i + 1..];
            }

            data.entry(key)
                .and_modify(|value| match value {
                    Value::Single(v) => {
                        // Vec::new(); vec.push(v); vec.push(val);
                        // let mut vec = vec![value, val]
                        *value = Value::Multiple(vec![v, val]);
                    }
                    Value::Multiple(v) => v.push(val),
                })
                .or_insert(Value::Single(val)); // if key not exists
        }

        QueryString { data: data }
    }
}
