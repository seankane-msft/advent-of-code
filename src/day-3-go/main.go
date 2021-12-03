package main

import (
	"bufio"
	"fmt"
	"os"
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

	var input []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	fmt.Printf("Answer to part 1: %d\n", partOne(input))
	fmt.Printf("Answer to part 2: %d", partTwo(input))
}

func partTwo(values []string) int {
	if len(values) == 0 {
		return -1
	}

	var remaining []string
	remaining = append(remaining, values...)

	for i := 0; i < len(values[0]); i++ {
		oneCount := oneCount(remaining, i)
		var removeChar byte
		if oneCount*2 >= len(remaining) && oneCount != 0 {
			removeChar = '1'
		} else {
			removeChar = '0'
		}
		// fmt.Printf("\nONE COUNT: %d\t, Middle %d\t: i: %d\t removeChar: %s\n", oneCount, len(remaining)/2, i, string(removeChar))
		// fmt.Println(remaining)
		for j := 0; j < len(remaining); j++ {
			if remaining[j][i] != removeChar {
				remaining = remove(remaining, j)
				j--
			}
		}
	}
	// convert string to []int
	var ogr []int
	for _, c := range remaining[0] {
		if c == '0' {
			ogr = append(ogr, 0)
		} else {
			ogr = append(ogr, 1)
		}
	}
	oxyGenRating := binaryToInt(ogr)

	// Find the c02 scrubber rating
	remaining = make([]string, 0)
	remaining = append(remaining, values...)

	for i := 0; i < len(values[0]); i++ {
		if len(remaining) == 1 {
			break
		}
		oneCount := oneCount(remaining, i)
		var removeChar byte
		if oneCount*2 < len(remaining) && oneCount != 0 {
			removeChar = '1'
		} else {
			removeChar = '0'
		}
		// fmt.Printf("\nONE COUNT: %d\t, Middle %d\t: i: %d\t removeChar: %s\n", oneCount, len(remaining)/2, i, string(removeChar))
		// fmt.Println(remaining)
		for j := 0; j < len(remaining); j++ {
			if remaining[j][i] != removeChar {
				remaining = remove(remaining, j)
				j--
			}
		}
	}

	// convert string to []int
	var co2 []int
	for _, c := range remaining[0] {
		if c == '0' {
			co2 = append(co2, 0)
		} else {
			co2 = append(co2, 1)
		}
	}
	co2Rating := binaryToInt(co2)

	return oxyGenRating * co2Rating
}

// remove the 's' index from a slice
func remove(slice []string, s int) []string {
	return append(slice[:s], slice[s+1:]...)
}

func oneCount(values []string, c int) int {
	oneCount := 0
	for _, v := range values {
		if v[c] == '1' {
			oneCount++
		}
	}
	return oneCount
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
	for i := len(binary) - 1; i >= 0; i-- {
		ret += binary[i] * base
		base *= 2
	}
	return ret
}
