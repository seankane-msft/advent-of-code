package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName = "input.txt"

func ParseInputs(values []string) [][]int {
	var ret [][]int
	for _, value := range values {
		var row []int
		for _, c := range value {
			i, err := strconv.Atoi(string(c))
			handle(err)
			row = append(row, i)
		}
		ret = append(ret, row)
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
	// partTwo(input)
}

func partOne(values [][]int) {
	flashes := 0

	allFlashed := false

	for count := 0; count < 100; count++ {
		// fmt.Println()
		// PrintBoard(values)

		numFlashes := simulateStep(values)
		if numFlashes == 100 {
			allFlashed = true
			fmt.Println("The answer to part 2 is ", count)
		}
		flashes += numFlashes
		// fmt.Println(flashes)
		// PrintBoard(values)
		// break
	}

	fmt.Println("The answer to part 1 is ", flashes)

	count := 0
	for !allFlashed {
		count += 1
		numFlashes := simulateStep(values)
		if numFlashes == 100 {
			allFlashed = true
			fmt.Println("The answer to part 2 is ", 100 + count)
		}
		flashes += numFlashes
	}
}

type Point struct {
	x, y int
}

func simulateStep(board [][]int) int {
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board); j++ {
			board[i][j] += 1
		}
	}

	flashedSpots := make(map[Point]bool)

	oldSize := 1
	for oldSize != len(flashedSpots) {
		oldSize = len(flashedSpots)
		for i := 0; i < len(board[0]); i++ {
			for j := 0; j < len(board); j++ {
				if board[i][j] > 9 {
					p := Point{x: i, y: j}
					if _, ok := flashedSpots[p]; !ok {
						// this spot has not been taken care of
						incrementNeighbors(board, i, j)

						flashedSpots[p] = true
					}
				}
			}
		}
	}

	// Do the resets

	flashcount := 0
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] > 9 {
				board[i][j] = 0
				flashcount += 1
			}
		}
	}

	return flashcount
	// return len(flashedSpots)
}

func incrementNeighbors(board [][]int, i, j int) {
	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if !(i+r < 0 || i+r > len(board)-1 || j+c > len(board)-1 || j+c < 0) {
				board[i+r][j+c] += 1
			}
		}
	}
}

func partTwo(values []string) {

}

func PrintBoard(board [][]int) {
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board[0]); j++ {
			val := board[i][j]
			fmt.Printf("%d ", val)
		}
		fmt.Println()
	}
}
