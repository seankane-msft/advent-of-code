package main

import (
	"bufio"
	"fmt"
	"os"
)

const fileName = "sample.txt"

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	/* Part One */
	// First step, read the input file into a []int
	file, err := os.Open(fileName)
	handle(err)
	defer file.Close()

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Printf("Answer to part 1: %d", partOne(input))
}

func partOne(values []string) int {
	if len(values) == 0 {
		return -1
	}
	mostCommonBits := make([]int, len(values[0]))
	middle := len(values) / 2

	for _, val := range values {
		for i, char := range val {
			if char == '1' {
				mostCommonBits[i]++
			}
		}
	}

	gamma := make([]int, len(values[0]))
	epsilon := make([]int, len(values[0]))
	for i, m := range mostCommonBits {
		if m > middle {
			gamma[i] = 1
			epsilon[i] = 0
		} else {
			gamma[i] = 0
			epsilon[i] = 1
		}
	}

	return binaryToInt(gamma) * binaryToInt(epsilon)
}

func binaryToInt(binary []int) int {
	var ret int
	base := 1
	for i := len(binary)-1; i >= 0; i-- {
		ret += binary[i] * base
		base *= 2
	}
	return ret
}