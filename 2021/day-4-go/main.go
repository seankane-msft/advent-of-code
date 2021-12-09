package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const fileName = "input.txt"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Space struct {
	Number  int
	Checked bool
}

type BingoBoard struct {
	Spaces     [][]Space
	FoundBingo bool
}

func NewBingoBoard(values []string) *BingoBoard {
	if len(values) != 5 {
		panic("Length of values is not 5")
	}
	b := BingoBoard{
		Spaces:     make([][]Space, 0),
		FoundBingo: false,
	}

	for _, row := range values {
		var bRow []Space
		parts := strings.Split(row, " ")
		for _, p := range parts {
			if p != "" {
				i, err := strconv.Atoi(p)
				check(err)
				bRow = append(bRow, Space{
					Number:  i,
					Checked: false,
				})
			}
		}
		if len(bRow) == 5 {
			b.Spaces = append(b.Spaces, bRow)
		} else {
			panic("DIDN't WORK")
		}
	}

	return &b
}

func (b *BingoBoard) Check(val int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Spaces[i][j].Number == val {
				b.Spaces[i][j].Checked = true
				return
			}
		}
	}
}

func (b *BingoBoard) FinalScore(last int) int {
	summed := 0
	for _, row := range b.Spaces {
		for _, col := range row {
			if !col.Checked {
				summed += col.Number
			}
		}
	}
	return summed * last
}

func (b *BingoBoard) HasBingo() bool {
	for i := 0; i < 5; i++ {
		// Check row i first
		if b.Spaces[i][0].Checked && b.Spaces[i][1].Checked && b.Spaces[i][2].Checked && b.Spaces[i][3].Checked && b.Spaces[i][4].Checked {
			b.FoundBingo = true
			return true
		}

		// Check column i
		if b.Spaces[0][i].Checked && b.Spaces[1][i].Checked && b.Spaces[2][i].Checked && b.Spaces[3][i].Checked && b.Spaces[4][i].Checked {
			b.FoundBingo = true
			return true
		}
	}
	return false
}

func createBingoBoards(values []string) []*BingoBoard {
	var ret []*BingoBoard
	for i := 0; i < len(values); i += 6 {
		ret = append(ret, NewBingoBoard(values[i:i+5]))
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
	bingoNumbers := input[0]

	bingoBoardsText := input[2:]
	bingoBoards := createBingoBoards(bingoBoardsText)

	partOne(bingoNumbers, bingoBoards)
}

func partOne(values string, bingoBoards []*BingoBoard) {
	bingoCount := 0
	for count, val := range strings.Split(values, ",") {
		i, err := strconv.Atoi(val)
		check(err)

		for _, bingoBoard := range bingoBoards {

			if !bingoBoard.FoundBingo {
				bingoBoard.Check(i)

				if count > 5 {
					if bingoBoard.HasBingo() {
						bingoCount += 1
						if bingoCount == 1 {
							fmt.Printf("Solution to part one: %d\n", bingoBoard.FinalScore(i))
						} else if bingoCount == len(bingoBoards) {
							fmt.Printf("Solution to part two: %d\n", bingoBoard.FinalScore(i))
						}
					}
				}
			}

		}
	}
}
