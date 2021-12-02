# advent-of-code
Source for the Advent of Code 2021

I'm using both Go and Rust, most likely trying Rust first to help learn the language. I'm also going to include notes about improvements. I'm using the bash `time` functionality to time the solutions for my own curiosity.

## Day 1
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build . && time ./day-1` | Go | 0.025s | 0.000s | 0.031s |
| `cargo build && time target/debug/day-1-rust.exe` | Rust | 0.048s | 0.000s | 0.031s |

Improvements:
* Rust reads the same input file twice, probably easier to read it once and pass the `Vec<int32>` to each program. (Implemented later)

## Day 2
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `cargo build && time target/debug/day-2-rust.exe` | Rust | 0.051s | 0.000s | 0.031s |

Improvements:
* I use an `if/else` for choosing which path to go down, I can convert this to a match statement which is more Rust idiomatic (I think). Implemented, looks cleaner at least.

Notes:
* Able to reuse a decent amount of Rust code from day 1's exercise. Reading the input, collecting as a vector. The `String` vs `&str` is still confusing to me, but the compiler was helpful in giving good hints that I didn't even have ot reach for SO/Google/etc.