package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

const fileName = "input.txt"

var partTwoFuelCost = make(map[int]int)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseInputs(values string) []int {
	var ret []int
	for _, value := range strings.Split(values, ",") {
		t, err := strconv.Atoi(value)
		check(err)
		ret = append(ret, t)
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
	if len(input) != 1 {
		panic(fmt.Errorf("expected input length of 1, got %d", len(input)))
	}

	positions := ParseInputs(input[0])
	partOne(positions)
}

func partOne(positions []int) {
	max := 0
	min := 99999999
	for _, p := range positions {
		if p > max {
			max = p
		} else if p < min {
			min = p
		}
	}

	median := findMedian(positions)
	fmt.Printf("The answer to part 1 is %d\n", calculateFuel(positions, median))

	mean := Mean(positions)
	lowestFuel2 := calculatePartTwoFuel(positions, mean)
	for i := mean-1; i <= mean+1; i++ {
		fuelSpent := calculatePartTwoFuel(positions, i)
		if fuelSpent < lowestFuel2 {
			lowestFuel2 = fuelSpent
		}
	}
	fmt.Printf("The answer to part 2 is %d\n", lowestFuel2)
}

func Mean(positions []int) int {
	sum := 0
	for _, p := range positions {
		sum += p
	}
	return sum / len(positions)
}

func findMedian(positions []int) int {
	sort.Ints(positions)
	middle := len(positions) / 2
	if len(positions) % 2 == 0 {
		return (positions[middle-1] + positions[middle]) / 2
	}
	return positions[middle]
}

func calculateFuel(positions []int, move int) int {
	sum := 0
	for _, p := range positions {
		sum += abs(p - move)
	}
	return sum
}

func calculatePartTwoFuel(positions []int, move int) int {
	sum := 0

	for _, p := range positions {
		diff := abs(p - move)
		if fuelCost, ok := partTwoFuelCost[diff]; ok {
			sum += fuelCost
		} else {
			fuelCost := (diff * diff + diff) / 2
			sum += fuelCost
			partTwoFuelCost[diff] = fuelCost
		}
	}

	return sum
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}
