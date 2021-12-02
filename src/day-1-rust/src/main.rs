use std::fs::File;
use std::io::{self, BufRead};
use std::path::Path;
use std::vec::Vec;

const FILE_NAME: &str = "./src/input.txt";

fn main() {
    print!("Answer for part 1: ");
    part_one();

    print!("Answer for part 2: ");
    part_two();
}

// find the number of increasing windows of size 3
fn part_two() {
    let mut count: i32 = 0;
    let mut vec = Vec::new();

    if let Ok(lines) = read_lines(FILE_NAME) {
        for line in lines {
            if let Ok(ip) = line {
                let int_conv = to_int(ip);
                vec.push(int_conv);
            }
        }
    }

    for i in 4..vec.len()+1 {
        if sum_slice(&vec[i-3..i]) > sum_slice(&vec[i-4..i-1]) {
            count += 1;
        }
    }

    print!("{}\n", count);
}

// sum a slice of a Vec<i32>
fn sum_slice(values: &[i32]) -> i32 {
    let mut sum: i32 = 0;
    for v in values.iter() {
        sum += v;
    }
    sum
}

fn part_one() {
    let mut prev: i32 = -1;
    let mut count: i32 = 0;
    if let Ok(lines) = read_lines(FILE_NAME) {
        for line in lines {
            if let Ok(ip) = line {
                let int_conv = to_int(ip);
                // println!("{}", to_int(ip));
                if prev != -1 {
                    if int_conv > prev {
                        count += 1;
                    }
                }
                prev = int_conv;
            }
        }
    }
    print!("{}\n", count);
}

fn to_int(s: String) -> i32 {
    s.parse().unwrap_or(0)
}

fn read_lines<P>(filename: P) -> io::Result<io::Lines<io::BufReader<File>>> where P: AsRef<Path>, {
    let file = File::open(filename)?;
    Ok(io::BufReader::new(file).lines())
}
