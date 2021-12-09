package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const fileName = "input.txt"

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

	var input []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		handle(err)
		input = append(input, v)
	}

	// Part 1
	vals := make(map[int]bool)
	for _, v := range input {
		if _, ok := vals[2020-v]; ok {
			fmt.Println("Part 1 answer: ", (2020-v) * v)
		}
		vals[v] = true
	}

	for _, v := range input {
		for key := range vals {
			if _, ok := vals[2020 - v - key]; ok {
				fmt.Println("Part 2 answer: ", key * v * (2020 - v - key))
				return
			}
		}
	}
}
