use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::vec::Vec;

const FILE_NAME: &str = "./input.txt";

fn main() {
    let v: Vec<String> = read_input();

    println!("Answer for part 1: {}", part_one(&v));
    println!("Answer for part 2: {}", part_two(&v));
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

// Find the product of the horizontal position and the depth.
fn part_one(vec: &Vec<String>) -> i32 {
    let mut horizontal: i32 = 0;
    let mut depth: i32 = 0;

    for v in vec.iter() {
        let split_vec: Vec<&str> = v.split(" ").collect();
        let orientation = split_vec[0];
        let magnitude: i32 = to_int(split_vec[1].to_string());
        if orientation == "forward" {
            horizontal += magnitude;
        } else if orientation == "down" {
            depth += magnitude;
        }else if orientation == "up" {
            depth -= magnitude;
        }
    }
    depth * horizontal
}

// find the number of increasing windows of size 3
fn part_two(vec: &Vec<String>) -> i32  {
    let mut horizontal: i32 = 0;
    let mut depth: i32 = 0;
    let mut aim: i32 = 0;

    for v in vec.iter() {
        let split_vec: Vec<&str> = v.split(" ").collect();
        let orientation = split_vec[0];
        let magnitude: i32 = to_int(split_vec[1].to_string());
        if orientation == "forward" {
            horizontal += magnitude;
            depth += aim * magnitude;
        } else if orientation == "down" {
            aim += magnitude;
        }else if orientation == "up" {
            aim -= magnitude;
        }
    }
    depth * horizontal
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
