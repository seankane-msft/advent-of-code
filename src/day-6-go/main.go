package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "sample.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ParseInputs(values string) map[int]int {

	ret := make(map[int]int)
	for _, value := range strings.Split(values, ",") {
		t, err := strconv.Atoi(value)
		check(err)
		if v, ok := ret[t]; ok {
			ret[t] = v + 1
		} else {
			ret[t] = 1
		}
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

	lanternFish := ParseInputs(input[0])
	partOne(lanternFish)
}

func partOne(fish map[int]int) {
	for i := 0; i < 256; i++ {

		if i == 80 {
			sum := 0
			for _, v := range fish {
				sum += v
			}
			fmt.Printf("The answer to part 1 is %d\n", sum)
		}

		newFish := make(map[int]int)
		for j := 8; j > 0; j-- {
			if val, ok := fish[j]; ok {
				newFish[j-1] = val
			} else {
				newFish[j-1] = 0
			}
		}

		newFish[6] += fish[0]
		newFish[8] = fish[0]
		fish = newFish
	}

	sum := 0
	for _, v := range fish {
		sum += v
	}
	fmt.Printf("The answer to part 2 is %d\n", sum)
}
