/**
 * Two Sum
 */
use std::collections::HashMap;

fn main() {
    println!("{:?}", two_sum(&vec![1, 2, 3, 4, 5], 6));
}

fn two_sum(nums: &Vec<usize>, target: usize) -> Result<Vec<usize>, String> {
    if nums.len() < 2 {
        return Err("not found".to_string());
    }

    let mut index: HashMap<usize, usize> = HashMap::with_capacity(nums.len());

    for (i, v) in nums.iter().enumerate() {
        if let Some(p) = index.get(&(target - v)) {
            return Ok(vec![i, *p]);
        }
        index.insert(*v, i);
    }

    return Err("not found".to_string());
}
