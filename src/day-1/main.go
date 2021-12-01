package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func sum(vals []int) int {
	count := 0
	for _, v := range vals {
		count += v
	}
	return count
}

func main() {
	/* Part One */
	// First step, read the input file into a []int
	file, err := os.Open("input.txt")
	handle(err)
	defer file.Close()

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		handle(err)
		input = append(input, v)
	}

	// Count number that increase
	count := 0
	for i := 1; i < len(input); i++ {
		if input[i] > input[i-1] {
			count++
		}
	}
	fmt.Println(count)

	/* Part Two */
	// Part two uses sliding windows of three
	count = 0
	for i := 4; i <= len(input); i++ {
		j := i - 1
		if sum(input[i-3:i]) > sum(input[j-3:j]) {
			count++
		}
	}
	fmt.Println(count)
}
