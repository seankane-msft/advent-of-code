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

### Day 3
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-3-go.exe` | Go | 0.068s | 0.000s | 0.062s |

Improvements:
* There's probably several, I'm copying an entire array twice, I think I could do the first step in parallel but after that it has to be switched out.
* I can at least make this more modular, but for now I'm keeping as is.

Notes:
* Started out in Rust, got part 1 but part 2 I couldn't figure out an elegant way to do it in rust vectors. Go slices worked magically here.
* When I finished there were 51,200 who had completed both, with another 24,136 who had completed part 1.

### Day 4
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-4-go.exe` | Go | 0.079s | 0.000s | 0.062s |

Improvements:
* There's probably a better algorithm for checking the bingo boards, or going through and finding at which value each array has succeeded.
* Overall, I'm pretty happy with the solution I came up with on the first go about. 161 lines, but a lot of nested stuff.

Notes:
* Did this puzzle on day 5, had an important UC Bearcats football game to go watch.
* When I finished there were 59090 who had completed both, with another 3672 who had completed part 1.

### Day 5
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-5-go.exe` | Go | 0.032s | 0.000s | 0.031s |

Improvements:
* Some algorithm I'm not even slightly aware of

Notes:
* When I finished there were 45582 who had completed both, with another 3183 who had completed part 1.
    * 31% of the completions of the day 1 puzzle for completing both
* Had all kinds of silly bugs, like using `<` instead of `<=` in a few spots.