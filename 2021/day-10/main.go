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

	// values := ParseInputs(input)
	partOne(input)
	partTwo(input)
}

func partOne(values []string) {
	parentheses := 0
	brace := 0
	bracket := 0
	carat := 0

	for _, chunk := range values {
		firstIncorrect := isCorrupt(chunk)
		switch firstIncorrect {
		case ')':
			parentheses += 1
		case ']':
			brace += 1
		case '}':
			bracket += 1
		case '>':
			carat += 1
		}
	}

	result := parentheses * 3
	result += brace * 57
	result += bracket * 1197
	result += carat * 25137
	fmt.Printf("The answer to part 1 is %d\n", result)
}

func isCorrupt(chunk string) rune {
	var lastOpening []rune

	for _, c := range chunk {
		if isOpening(c) {
			lastOpening = append(lastOpening, c)
		} else {
			if matchesOpen(lastOpening[len(lastOpening)-1], c) {
				lastOpening = lastOpening[:len(lastOpening)-1]
			} else {
				return c
			}
		}
	}

	return 'a'
}

func isOpening(r rune) bool {
	return r == '(' || r == '{' || r == '[' || r == '<'
}

func matchesOpen(open, close rune) bool {
	if open == '(' && close == ')' {
		return true
	}
	if open == '{' && close == '}' {
		return true
	}
	if open == '[' && close == ']' {
		return true
	}
	if open == '<' && close == '>' {
		return true
	}
	return false
}

func partTwo(values []string) {
	var incompletes []int

	for _, chunk := range values {
		v := lineValuePartTwo(chunk)
		if v > 0 {
			incompletes = append(incompletes, v)
		}
	}

	sort.Ints(incompletes)
	fmt.Printf("The answer to part 2 is %d\n", incompletes[len(incompletes) / 2])
}

func lineValuePartTwo(chunk string) int {
	var lastOpening []rune

	for _, c := range chunk {
		if isOpening(c) {
			lastOpening = append(lastOpening, c)
		} else {
			if matchesOpen(lastOpening[len(lastOpening)-1], c) {
				lastOpening = lastOpening[:len(lastOpening)-1]
			}else {
				return 0
			}
		}
	}

	score := 0
	for i := len(lastOpening) - 1; i >= 0; i-- {
		score *= 5
		switch lastOpening[i] {
			case '(': score += 1
			case '[': score += 2
			case '{': score += 3
			case '<': score += 4
		}
	}

	return score

}
