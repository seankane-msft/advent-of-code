use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::vec::Vec;

const FILE_NAME: &str = "./src/input.txt";

fn main() {
    let v: Vec<i32> = read_input();

    println!("Answer for part 1: {}", part_one(&v));
    println!("Answer for part 2: {}", part_two(&v));
}

// read the input into a single vector that can be shared by both parts (it's read only for both)
fn read_input() -> Vec<i32> {
    let mut vec = Vec::new();
    if let Ok(lines) = read_lines(FILE_NAME) {
        for line in lines {
            if let Ok(ip) = line {
                let int_conv = to_int(ip);
                vec.push(int_conv);
            }
        }
    }
    vec
}

// find the number of increasing windows of size 3
fn part_two(vec: &Vec<i32>) -> i32 {
    let mut count: i32 = 0;

    for i in 4..vec.len()+1 {
        if sum_slice(&vec[i-3..i]) > sum_slice(&vec[i-4..i-1]) {
            count += 1;
        }
    }
    count
}

// sum a slice of a Vec<i32>
fn sum_slice(values: &[i32]) -> i32 {
    let mut sum: i32 = 0;
    for v in values.iter() {
        sum += v;
    }
    sum
}

// Find the number of times the next number was greater than the previous
fn part_one(vec: &Vec<i32>) -> i32 {
    let mut count: i32 = 0;
    for i in 1..vec.len() {
        if vec[i] > vec[i-1] {
            count += 1 ;
        }
    }
    count
}

// convert a string to an integer
fn to_int(s: String) -> i32 {
    s.parse().unwrap_or(0)
}

// read a file
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
