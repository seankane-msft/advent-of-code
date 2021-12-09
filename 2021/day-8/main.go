package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileName = "input.txt"

func ParseInputs(values []string) [][]string {
	var ret [][]string
	for _, value := range values {
		parts := strings.Split(value, "|")
		ret = append(ret, parts)
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

func partOne(values [][]string) {
	count := 0
	for _, value := range values {
		easyDigits := value[1]
		easyDigits = strings.TrimSpace(easyDigits)

		parts := strings.Split(easyDigits, " ")

		for _, part := range parts {
			if len(part) == 3 || len(part) == 4 || len(part) == 7 || len(part) == 2 {
				count += 1
			}
		}
	}
	fmt.Printf("The answer to part 1 is %d\n", count)
}

func partTwo(values [][]string) {
	sum := 0
	for _, value := range values {
		sum += decodeLine(value)
	}
	fmt.Printf("The answer to part 2 is %d\n", sum)
}

type Display struct {
	a               string
	b               string
	c               string
	d               string
	e               string
	f               string
	g               string
	lettersDiscover string
}

func (d Display) String() string {
	return fmt.Sprintf("{a: %s, b: %s, c: %s, d: %s, e: %s, f: %s, g: %s}\tLetters Discovered: %s", d.a, d.b, d.c, d.d, d.e, d.f, d.g, d.lettersDiscover)
}

func (d *Display) AddA(values []string) {
	one, seven := "", ""
	for _, value := range values {
		if len(value) == 2 {
			one = value
		} else if len(value) == 3 {
			seven = value
		}
	}

	for _, char := range seven {
		if !strings.Contains(one, string(char)) {
			d.a = string(char)
			d.lettersDiscover = fmt.Sprintf("%s%s", d.lettersDiscover, string(char))
		}
	}
}

func (d *Display) AddCF(values []string) {
	var one string
	var sixzero []string

	for _, value := range values {
		if len(value) == 2 {
			one = value
		} else if len(value) == 6 {
			sixzero = append(sixzero, value)
		}
	}

	if one == "" || len(sixzero) == 0 {
		return
	}

	// Find the six value. The missing value between six and one belongs at c position
	for _, sz := range sixzero {
		for _, c := range one {
			if !strings.Contains(sz, string(c)) {
				d.c = string(c)
				d.lettersDiscover = fmt.Sprintf("%s%s", d.lettersDiscover, string(c))
				break
			}
		}
	}

	for _, c := range one {
		if string(c) != d.c {
			d.f = string(c)
			d.lettersDiscover = fmt.Sprintf("%s%s", d.lettersDiscover, string(c))
		}
	}
}

// E can be done by letter frequency
func (d *Display) AddBE(values []string) {
	m := make(map[rune]int)

	for _, value := range values {
		for _, c := range value {
			if v, ok := m[c]; ok {
				m[c] = v + 1
			} else {
				m[c] = 1
			}
		}
	}

	for k, v := range m {
		if v == 4 {
			// Has to be the E spot
			d.e = string(k)
			d.lettersDiscover = fmt.Sprintf("%s%s", d.lettersDiscover, string(k))
		} else if v == 6 {
			if d.a != string(k) {
				// Has to be the B spot
				d.b = string(k)
				d.lettersDiscover = fmt.Sprintf("%s%s", d.lettersDiscover, string(k))
			}
		}
	}
}

// D and G can be done by letter frequency
func (d *Display) AddDG(values []string) {
	var four string
	for _, value := range values {
		if len(value) == 4 {
			four = value
			break
		}
	}

	for _, c := range four {
		if !strings.Contains(d.lettersDiscover, string(c)) {
			d.d = string(c)
			d.lettersDiscover = fmt.Sprintf("%s%s", d.lettersDiscover, string(c))
			break
		}
	}

	for _, c := range "abcdefg" {
		if !strings.Contains(d.lettersDiscover, string(c)) {
			d.g = string(c)
			d.lettersDiscover = fmt.Sprintf("%s%s", d.lettersDiscover, string(c))
			break
		}
	}
}

func (d *Display) Decode(value string) int {
	if len(value) == 2 {
		return 1
	} else if len(value) == 4 {
		return 4
	} else if len(value) == 3 {
		return 7
	} else if len(value) == 7 {
		return 8
	}

	if len(value) == 5 {
		// 2, 3, or 5
		if strings.Contains(value, d.e) {
			return 2
		}

		if strings.Contains(value, d.c) {
			return 3
		}

		return 5
	}

	if len(value) == 6 {
		// 6 or 9 (0 is default)
		if !strings.Contains(value, d.c) {
			return 6
		} else if !strings.Contains(value, d.e) {
			return 9
		}
	}

	return 0
}

func decodeLine(displays []string) int {
	displays[0] = strings.TrimSpace(displays[0])
	parts := strings.Split(displays[0], " ")

	display := &Display{}

	display.AddA(parts)
	display.AddCF(parts)
	display.AddBE(parts)
	display.AddDG(parts)

	val := 0
	for _, value := range strings.Split(displays[1], " ") {
		val = val * 10 + display.Decode(value)
	}

	return val
}
