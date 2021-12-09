package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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
	partTwo(values)
}

func partOne(values [][]int) {
	count := 0
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[0]); j++ {
			if isLowPoint(values, i, j) {
				count += values[i][j] + 1
			}
		}
	}
	fmt.Printf("The answer to part 1 is %d\n", count)
}

func isLowPoint(values [][]int, i, j int) bool {

	if i != 0 {
		if values[i][j] >= values[i-1][j] {
			return false
		}
	}

	if i != len(values)-1 {
		if values[i][j] >= values[i+1][j] {
			return false
		}
	}

	if j != 0 {
		if values[i][j] >= values[i][j-1] {
			return false
		}
	}

	if j != len(values[0])-1 {
		if values[i][j] >= values[i][j+1] {
			return false
		}
	}

	return true
}

// 1. find all my low points
// 2. Create a map of low points to basin size
// 3. Go through the values slice again and find where each point flows to
// 3b. if the value is 9, it doesn't belong in a basin
func partTwo(values [][]int) {
	var basinMap [][]int
	for _, row := range values {
		basinMap = append(basinMap, make([]int, len(row)))
	}

	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[0]); j++ {
			if values[i][j] == 9 {
				basinMap[i][j] = -1
			}
		}
	}

	basinCount := 1
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[0]); j++ {
			if isLowPoint(values, i, j) {
				basinMap[i][j] = basinCount
				basinCount += 1
			}
		}
	}

	for {
		zeroCount := 0
		for i := 0; i < len(values); i++ {
			for j := 0; j < len(values[0]); j++ {
				if basinMap[i][j] == 0 {
					zeroCount++
					b := findBasinNumber(basinMap, i, j)
					if b > 0 {
						basinMap[i][j] = b
					}
				}
			}
		}
		if zeroCount == 0 {
			break
		}
	}

	basinSize := make(map[int]int)
	for i := 0; i < len(values); i++ {
		for j := 0; j < len(values[0]); j++ {
			basinNumber := basinMap[i][j]
			if basinNumber > 0 {
				if val, ok := basinSize[basinNumber]; ok {
					basinSize[basinNumber] = val + 1
				} else {
					basinSize[basinNumber] = 1
				}
			}
		}
	}

	threeLargest := []int{0, 0, 0}
	for _, v := range basinSize {
		if v > threeLargest[0] {
			threeLargest[0] = v
			sort.Ints(threeLargest)
		}
	}

	fmt.Printf("The answer to part two is %d", threeLargest[0]*threeLargest[1]*threeLargest[2])
}

func findBasinNumber(basinMap [][]int, i, j int) int {
	if i != 0 {
		if basinMap[i-1][j] > 0 {
			return basinMap[i-1][j]
		}
	}

	if i != len(basinMap)-1 {
		if basinMap[i+1][j] > 0 {
			return basinMap[i+1][j]
		}
	}

	if j != 0 {
		if basinMap[i][j-1] > 0 {
			return basinMap[i][j-1]
		}
	}

	if j != len(basinMap[0])-1 {
		if basinMap[i][j+1] > 0 {
			return basinMap[i][j+1]
		}
	}
	return -1
}
