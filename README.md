# advent-of-code
Source for the Advent of Code 2021

I'm using both Go and Rust, mostly Go first then trying to reimplement in Rust as I try and pick up the language. I'm also going to include notes about improvements. I'm using the bash `time` functionality to time the solutions for my own curiosity.

## Day 1
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `time go run .` | Go | 1.061s | 0.016s | 0.031s |
| `go build . && time ./day-1` | Go | 0.025s | 0.000s | 0.031s |
| `time cargo run` | Rust | 0.729s | 0.000s | 0.015s |
| `cargo build && time target/debug/day-1-rust.exe` | Rust | 0.048s | 0.000s | 0.031s |

Improvements:
* Rust reads the same input file twice, probably easier to read it once and pass the `Vec<int32>` to each program. (Implemented later)