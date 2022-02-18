/**
 * Two Sum Input Array Is Sorted
 */

fn main() {
    println!("{:?}", two_sum_sorted(&vec![1, 2, 3, 4, 5], 7));
    println!("{:?}", two_sum_sorted(&vec![1, 3, 7, 9, 11], 13));
    println!("{:?}", two_sum_sorted(&vec![1, 7, 8, 9], 16));
}

fn two_sum_sorted(nums: &Vec<usize>, target: usize) -> Result<Vec<usize>, String> {
    if nums.len() < 2 {
        return Err("not found".to_string());
    }

    let (mut i, mut j) = (0, nums.len() - 1);

    while i < j {
        println!("~~~ {}, {}", i, j);
        let sum = nums[i] + nums[j];

        match sum {
            _ if sum == target => {
                return Ok(vec![i, j]);
            }
            _ if (sum < target) => i += 1,
            _ => j -= 1,
        }
    }

    return Err("not found".to_string());
}
