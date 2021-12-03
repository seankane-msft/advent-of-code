use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::vec::Vec;
use std::convert::TryInto;

const FILE_NAME: &str = "./sample.txt";

fn main() {
    let v: Vec<String> = read_input();

    println!("Answer for part 1: {}", part_one(&v));
    println!("Answer for part 2: {}", part_two(&v));
}

fn part_one(vec: &Vec<String>) -> i32 {
    let mut counts: Vec<i32> = Vec::new();

    for _ in 0..vec[0].len() {
        counts.push(0);
    }

    for v in vec.iter() {
        let mut i: usize = 0;
        for c in v.chars() {
            if c == '1' {
                counts[i] += 1;
            }
            i += 1;
        }
    }

    let middle: i32 = (vec.len() / 2).try_into().unwrap();
    let mut gamma = String::new();
    let mut epsilon = String::new();

    for c in counts.iter() {
        if c > &middle {
            gamma.push_str("1");
            epsilon.push_str("0");
        } else {
            gamma.push_str("0");
            epsilon.push_str("1");
        }
    }

    binary_to_i32(gamma) * binary_to_i32(epsilon)
}

fn part_two(vec: &Vec<String>) -> i32 {
    let mut counts: Vec<i32> = Vec::new();

    for _ in 0..vec[0].len() {
        counts.push(0);
    }

    // Create a copy of vec to start
    let mut remaining: Vec<String> = Vec::new();
    for v in vec.iter() {
        remaining.push(v);
    }


    for i in 0..counts {
        let mut new_remaining: Vec<String> = Vec::new();
        // Keep only that have the most common bit in

        let mut count: i32 = 0;
        // First, find the most common bit at position i
        for v in remaining.iter() {
            if v.chars().nth(i) == '1' {
                count += 1;
            }
        }
        println!("count {}", count)

        if count > (remaining.len() / 2).try_into().unwrap() {
            // Keep all 1s
        } else {
            // Keep all 0s
        }
    }

    0
}

// convert a binary string to an integer
fn binary_to_i32(s: String) -> i32 {
    let mut result: i32 = 0;
    let mut base: i32 = 2;
    base = base.pow((s.len() - 1).try_into().unwrap()).try_into().unwrap();
    for c in s.chars() {
        if c == '1' {
            result += base;
        }
        base /= 2;
    }
    result
}

// read the input into a single vector that can be shared by both parts (it's read only for both)
fn read_input() -> Vec<String> {
    let mut vec = Vec::new();
    if let Ok(lines) = read_lines(FILE_NAME) {
        for line in lines {
            if let Ok(ip) = line {
                vec.push(ip);
            }
        }
    }
    vec
}

// read a file
fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
