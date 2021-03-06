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

### Day 6
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-6-go.exe` | Go | 0.088s | 0.015s | 0.046s |

Improvements:
* Started out with a struct for LanternFish and simulating each one individually. This works fine for the first part but the second part it would take a lot longer. I cut off the program at two minutes and refactored.
* Second solution used a map that kept track of how many fish were in each cycle, this was both simpler and much faster.

Notes:
* When I finished there were 32217 who had completed both, with another 7436 who had completed part 1.
    * 23% of the completions of the day 1 puzzle for completing both. Day 1 is now at 148644
* This one felt much easier than the previous two days.

### Day 7
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-7.exe` | Go | 0.158s | 0.000s | 0.047s |
| `go build && time ./day-7.exe` | Go (using median) | 0.139s | 0.031s | 0.046s |
| `go build && time ./day-7.exe` | Go (using median for part 1, and mean for part 2) | 0.097 | 0.000s | 0.031s |

Improvements:
* First solution can be analytically determined as the median (thanks Reddit), which can be calculated by sorting the positions.
* The second solution uses the arithmetic mean and checks one number above and below.
* Second solution used a map that kept track of how much fuel each difference would cost instead of calculating it every time.

Notes:
* When I finished there were 39833 who had completed both, with another 2601 who had completed part 1.
    * 25% of the completions of the day 1 puzzle for completing both. Day 1 is now at 154620

### Day 8
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-8.exe` | Go | 0.093s | 0.000s | 0.046s |

Improvements:
* First solution was pretty easy
* Second solution was probably the toughest yet. I overcomplicated it a bit, didn't realize all ten digits were guaranteed to be given, didn't think about things like letter frequency but after a few iterations got through my solution pretty easily. Also happy to have it finish in > .1 seconds :)

Notes:
* When I finished there were 24289 who had completed both, with another 15567 who had completed part 1.
    * 15% of the completions of the day 1 puzzle for completing both. Day 1 is now at 159425

### Day 9
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-9.exe` | Go | 0.075s | 0.000s | 0.015s |

Improvements:
* The second solution sounds like something that can happen recursively which might be faster but no doubt more complex.
    * Turns out part two can be done with breadth first search (BFS), maintain a count size.

Notes:
* First solution was pretty easy, however indexing in Go slice of slices is a pain. This is where I wish there were traditional 2-D arrays.
* Second solution I made an assumption that paid off, I do it O(n**2) time by going through the whole map a couple times. I assume that each basin is lined by heights of 9 (which turned out to be true). The first time I give each low point a #. The second time I give each 9 a -1 value. I'm left with a bunch of points (labeled as zeros) that are not in a basin. I look at each zero and see if any of it's neighbors are assigned a basin, if they are I give that neighbor a basin. I repeat this process until each point is either a basin or a -1 (indicating a barrier), this only happened 9 times which is less than I thought would happen. Then I create a map to figure out how large each basin is. Finally, iterate through the map to find the three largest basins.
    * The spec says "all other locations will always be part of exactly one basin", so I guess the 9 barrier was a safe assumption.
* When I finished there were 23249 who had completed both, with another 9989 who had completed part 1.
    * 14% of the completions of the day 1 puzzle for completing both. Day 1 is now at 162756

### Day 10
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-10.exe` | Go | 0.071s | 0.000s | 0.046s |

Improvements:
* The second solution sounds like something that can happen recursively which might be faster but no doubt more complex.
    * Turns out part two can be done with breadth first search (BFS), maintain a count size.

Notes:
* This was almost exactly one of the questions I used when interview practicing so I recognized the solution almost immediately. It's a riff off the [Valid Parentheses](https://leetcode.com/problems/valid-parentheses/) leetcode question. All I had to do was add the scoring.
* Go doesn't have a stack type implementation in the standard library, but this is easy enough with slices.
* Rough timing, this took ~16 minutes from reading to completing both, the leaderboard gold cap was at 00:08:06 so I didn't really have much of a chance getting a star.
* When I finished there were 26754 who had completed both, with another 2307 who had completed part 1.
    * 16% of the completions of the day 1 puzzle for completing both. Day 1 is now at 165712

### Day 11
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-11.exe` | Go | 0.107s | 0.000s | 0.046s |

Improvements:
* I iterate over the board until the size of the map doesn't change. Instead you can check each square after you increment it.

Notes:
* When I finished there were 23178 who had completed both, with another 477 who had completed part 1.
    * 13.8% of the completions of the day 1 puzzle for completing both. Day 1 is now at 168457

### Day 12
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-12.exe` | Go | 0.923s | 0.000s | 0.046s |

Improvements:
* The conditional on my second can definitely be improved (at least in readability). I do a BFS search and storing possible paths from a given spot might save time.

Notes:
* When I finished there were 30160 who had completed both, with another 1914 who had completed part 1.
    * 17.5% of the completions of the day 1 puzzle for completing both. Day 1 is now at 171962
    * I completed this puzzle at 10:00 PM at night (spent the day at the browns game and watching football)

### Day 13
| Method | Language | Real | User | Sys |
| ------ | -------- | ---- | ---- | --- |
| `go build && time ./day-13.exe` | Go | 0.092s | 0.000s | 0.031s |

Improvements:
* There's probably a way to do this without storing the values in a `[][]int`, instead storing the points and manipulating them if they get reflected. Keeping the board was useful for debugging and for part two where you had to print them out.

Notes:
* When I finished there were 23577 who had completed both, with another 1183 who had completed part 1.
    * 13.6% of the completions of the day 1 puzzle for completing both. Day 1 is now at 172904