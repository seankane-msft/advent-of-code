use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::vec::Vec;

const FILE_NAME: &str = "./input.txt";

// Find the product of the horizontal position and the depth.
fn part_one(vec: &Vec<String>) -> i32 {
    let mut horizontal: i32 = 0;
    let mut depth: i32 = 0;

    for v in vec.iter() {
        let split_vec: Vec<&str> = v.split(" ").collect();
        let orientation = split_vec[0];
        let magnitude: i32 = to_int(split_vec[1].to_string());
        match orientation {
            "forward" => horizontal += magnitude,
            "down" => depth += magnitude,
            "up" => depth -= magnitude,
            _ => println!("Expected forward, down, or up. Got {}", orientation),
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
        match orientation {
            "forward" => {
                horizontal += magnitude;
                depth += aim * magnitude;
            },
            "down" => aim += magnitude,
            "up" => aim -= magnitude,
            _ => println!("Expected forward, down, or up. Got {}", orientation),
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
