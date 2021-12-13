package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "input.txt"

var folds []string

func ParseInputs(values []string) [][]int {
	var points [][]int

	maxX, maxY := 0, 0

	for _, value := range values {
		if strings.Contains(value, "fold along") {
			folds = append(folds, value)
		}
		parts := strings.Split(value, ",")
		if len(parts) == 2 {
			x, err := strconv.Atoi(parts[0])
			handle(err)
			y, err := strconv.Atoi(parts[1])
			handle(err)
			points = append(points, []int{x, y})
			if x > maxX {
				maxX = x
			}
			if y > maxY {
				maxY = y
			}
		}
	}

	var ret [][]int
	for i := 0; i <= maxY; i++ {
		ret = append(ret, make([]int, maxX+1))
	}

	fmt.Println(maxX, maxY)

	for _, point := range points {
		ret[point[1]][point[0]] = 1
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

	values := ParseInputs(input)
	partOne(values)
	// partTwo(values)
}

func parseFold(s string) (string, int) {
	parts := strings.Split(s, " ")
	line := parts[2]
	parts = strings.Split(line, "=")
	i, err := strconv.Atoi(parts[1])
	handle(err)
	return parts[0], i
}

func partOne(board [][]int) {
	// PrintBoard(board)
	l, i := parseFold(folds[0])
	// fmt.Println(l, i)

	var newBoard [][]int
	if l == "y" {
		newBoard = foldAlongY(board, i)
	} else if l == "x" {
		newBoard = foldAlongX(board, i)
	}
	// newBoard := foldAlongY(board, i)
	// PrintBoard(newBoard)

	count := 0
	for x := 0; x < len(newBoard[0]); x++ {
		for y := 0; y < len(newBoard); y++ {
			if newBoard[y][x] == 1 {
				count++
			}
		}
	}
	fmt.Println("The answer to part 1 is ", count)

	for _, fold := range folds[1:] {
		l, i := parseFold(fold)
		if l == "y" {
			newBoard = foldAlongY(newBoard, i)
		} else if l == "x" {
			newBoard = foldAlongX(newBoard, i)
		}
		// newBoard := foldAlongY(board, i)
		// PrintBoard(newBoard)
	}

	count = 0
	for x := 0; x < len(newBoard[0]); x++ {
		for y := 0; y < len(newBoard); y++ {
			if newBoard[y][x] == 1 {
				count++
			}
		}
	}
	// fmt.Println("The answer to part 2 is ", count)
	PrintBoard(newBoard)
}

func foldAlongY(board [][]int, foldLine int) [][]int {
	var ret [][]int

	for j := 0; j < foldLine; j++ {
		ret = append(ret, make([]int, len(board[0])))
	}

	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[y][x] == 1 {
				if y > foldLine {
					// fold so reflect along value
					newY := foldLine - (y - foldLine)
					ret[newY][x] = 1
				} else {
					ret[y][x] = 1
				}
			}
		}
	}

	return ret
}

func foldAlongX(board [][]int, foldLine int) [][]int {
	var ret [][]int

	for j := 0; j < len(board); j++ {
		ret = append(ret, make([]int, foldLine))
	}

	for x := 0; x < len(board[0]); x++ {
		for y := 0; y < len(board); y++ {
			if board[y][x] == 1 {
				if x > foldLine {
					// fold so reflect along value
					newX := foldLine - (x - foldLine)
					ret[y][newX] = 1
				} else {
					ret[y][x] = 1
				}
			}
		}
	}

	return ret
}

func PrintBoard(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			val := board[i][j]
			if val == 1 {
				fmt.Print("# ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}
}
