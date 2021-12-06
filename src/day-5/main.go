package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type LineSegment struct {
	x1, x2, y1, y2 int
}

func (l LineSegment) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.x1, l.y1, l.x2, l.y2)
}

func ParseInputs(values []string) []LineSegment {
	var ret []LineSegment
	for _, value := range values {
		parts := strings.Split(value, "->")
		if len(parts) != 2 {
			panic("Invalid length, expecting 2")
		}

		xs := strings.Split(parts[0], ",")
		ys := strings.Split(parts[1], ",")
		var l LineSegment

		x1, err := strconv.Atoi(strings.TrimSpace(xs[0]))
		check(err)
		l.x1 = x1

		y1, err := strconv.Atoi(strings.TrimSpace(xs[1]))
		check(err)
		l.y1 = y1

		x2, err := strconv.Atoi(strings.TrimSpace(ys[0]))
		check(err)
		l.x2 = x2

		y2, err := strconv.Atoi(strings.TrimSpace(ys[1]))
		check(err)
		l.y2 = y2

		ret = append(ret, l)
	}
	return ret
}

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/* Part One */
	// First step, read the input file into a []string
	file, err := os.Open(fileName)
	handle(err)
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	lines := ParseInputs(input)

	partOne(lines)
}

func partOne(lines []LineSegment) {
	maxValue := 0
	for _, line := range lines {
		if line.x1 > maxValue {
			maxValue = line.x1
		} else if line.x2 > maxValue {
			maxValue = line.x2
		} else if line.y1 > maxValue {
			maxValue = line.y1
		} else if line.y2 > maxValue {
			maxValue = line.y2
		}
	}

	var board [][]int
	for i := 0; i < maxValue+1; i++ {
		board = append(board, make([]int, maxValue+1))
	}

	for _, line := range lines {
		if line.x1 == line.x2 {
			// horizontal
			for i := min(line.y1, line.y2); i <= max(line.y1, line.y2); i++ {
				board[line.x1][i] += 1
			}
		} else if line.y1 == line.y2 {
			// vertical
			for i := min(line.x1, line.x2); i <= max(line.x1, line.x2); i++ {
				board[i][line.y1] += 1
			}
		} else {
			// diagonal
			if line.x1 < line.x2 {
				// Starts on the left of the board
				if line.y1 < line.y2 {
					// goes up and to the right
					for i := 0; i <= abs(line.y2, line.y1); i++ {
						board[line.x1+i][line.y1+i] += 1
					}
				} else {
					// goes down and to the right
					for i := 0; i <= abs(line.y2, line.y1); i++ {
						board[line.x1+i][line.y1-i] += 1
					}
				}
			} else {
				// line starts on the right side of the board
				if line.y1 < line.y2 {
					// goes up and to the left
					for i := 0; i <= abs(line.y2, line.y1); i++ {
						board[line.x1-i][line.y1+i] += 1
					}
				} else {
					// goes down and to the left
					for i := 0; i <= abs(line.y2, line.y1); i++ {
						board[line.x1-i][line.y1-i] += 1
					}
				}

			}

		}
	}
	// PrintBoard(board)

	// count number of overlaps
	count := 0
	for _, row := range board {
		for _, cell := range row {
			if cell >= 2 {
				count++
			}
		}
	}
	fmt.Printf("Answer for part 2: %d\n", count)
}

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func abs(i, j int) int {
	diff := i - j
	if diff >= 0 {
		return diff
	}
	return -diff
}

func PrintBoard(board [][]int) {
	for i := 0; i < len(board[0]); i ++ {
		for j := 0; j < len(board[0]); j ++ {
			val := board[j][i]
			if val > 0 {
				fmt.Printf("%d ", val)
			} else {
				fmt.Printf(". ")
			}
		}
		fmt.Println()
	}
}
