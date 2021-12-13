package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const fileName = "input.txt"

type Cave struct {
	large bool
	name  string
}

func NewCave(s string) Cave {
	large := false
	if strings.ToUpper(string(s[0])) == string(s[0]) {
		large = true
	}
	return Cave{
		large: large,
		name:  s,
	}
}

func ParseInputs(values []string) map[Cave][]Cave {
	ret := make(map[Cave][]Cave)

	for _, value := range values {
		parts := strings.Split(value, "-")
		if len(parts) != 2 {
			panic(fmt.Sprintf("Expected length of 2, got %d\n", len(parts)))
		}
		if _, ok := ret[NewCave(parts[0])]; ok {
			ret[NewCave(parts[0])] = append(ret[NewCave(parts[0])], NewCave(parts[1]))
		} else {
			ret[NewCave(parts[0])] = []Cave{NewCave(parts[1])}
		}

		if _, ok := ret[NewCave(parts[1])]; ok {
			ret[NewCave(parts[1])] = append(ret[NewCave(parts[1])], NewCave(parts[0]))
		} else {
			ret[NewCave(parts[1])] = []Cave{NewCave(parts[0])}
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

	values := ParseInputs(input)
	partOne(values)
	partTwo(values)
}

func partOne(caveMap map[Cave][]Cave) {
	// cave traversal on the firsty try? I need to go buy lotto tickets
	allPaths := make(map[string]bool)
	startingPoints := []string{"start"}
	findPaths(startingPoints, caveMap, allPaths)
	fmt.Println("The answer to part 1 is ", len(allPaths))
}

func partTwo(caveMap map[Cave][]Cave) {
	// cave traversal on the firsty try? I need to go buy lotto tickets
	allPaths := make(map[string]bool)
	startingPoints := []string{"start"}
	findPathsPartTwo(startingPoints, caveMap, allPaths)
	fmt.Println("The answer to part 2 is ", len(allPaths))
}

func findPaths(curPath []string, caveMap map[Cave][]Cave, allPaths map[string]bool) {
	curPosition := curPath[len(curPath)-1]
	// end condition #1: at the end
	if curPosition == "end" {
		allPaths[strings.Join(curPath, "-")] = true
		return
	}

	// end condition #2: nowhere to go
	placesToGo := caveMap[NewCave(curPosition)]
	if len(placesToGo) == 0 {
		return
	}

	// check if visited small cave, can't go there, otherwise spawn new paths
	for _, cave := range placesToGo {
		if inList(curPath, cave.name) && !cave.large {
			// do nothing
		} else {
			newPath := append(curPath, cave.name)
			findPaths(newPath, caveMap, allPaths)
		}
	}
}

func findPathsPartTwo(curPath []string, caveMap map[Cave][]Cave, allPaths map[string]bool) {
	curPosition := curPath[len(curPath)-1]
	// end condition #1: at the end
	if curPosition == "end" {
		allPaths[strings.Join(curPath, "-")] = true
		return
	}

	startCount := 0
	for _, visited := range curPath {
		if visited == "start" {
			startCount += 1
			if startCount > 1 {
				return
			}
		}
	}

	placesToGo := caveMap[NewCave(curPosition)]

	// check if can visit somewhere, can't go there, otherwise spawn new paths
	for _, cave := range placesToGo {
		if cave.name != "start" {
			if CanVisitPart2(append(curPath, cave.name)) {
				newPath := append(curPath, cave.name)
				findPathsPartTwo(newPath, caveMap, allPaths)
			}
		}
	}
}

func inList(l []string, v string) bool {
	for _, entry := range l {
		if entry == v {
			return true
		}
	}
	return false
}

func CanVisitPart2(cavesVisited []string) bool {
	smallCaves := make(map[string]int)
	doubleVisit := false
	for _, cave := range cavesVisited {
		if strings.ToLower(string(cave[0])) == string(cave[0]) {
			if v, ok := smallCaves[cave]; ok {
				smallCaves[cave] = v + 1
				if v+1 >= 2 {
					if doubleVisit || v+1 > 2 {
						return false
					}
					doubleVisit = true
				}
			} else {
				smallCaves[cave] = 1
			}
		}
	}

	return true
}
