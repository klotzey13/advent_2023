package day_4

import (
	steveutil "advent_of_code_2023/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
)

func isDelimiter(r rune) bool {
	return r == ':' || r == '|'
}

type ScratchCardsBundle struct {
	totalScratchCards int
	winningMatches    int
}

func Puzzle_2() {
	//Puzzle doesnt work for this puzzle2... Fix later
	filepath := path.Join("day_4", "input", "example.txt")
	f, err := os.Open(filepath)
	steveutil.Check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	scratchCardBundle := []ScratchCardsBundle{}

	for s.Scan() {
		line := strings.FieldsFunc(s.Text(), isDelimiter)
		winningNums := strings.Fields(line[1])
		numbers := strings.Fields(line[2])

		matches := 0
		for _, winNum := range winningNums {
			for _, num := range numbers {
				if winNum == num {
					matches++
				}
			}
		}

		scratchCardBundle = append(scratchCardBundle, ScratchCardsBundle{
			totalScratchCards: 1,
			winningMatches:    matches,
		})

	}

	//Increment for each winning match the total scratch cards
	for i := 0; i < len(scratchCardBundle); i++ {
		for x := 1; x <= scratchCardBundle[i].winningMatches; x++ {
			scratchCardBundle[x+i].totalScratchCards += scratchCardBundle[i].totalScratchCards
		}
	}

	total := 0
	for _, scratchCard := range scratchCardBundle {
		total = total + scratchCard.totalScratchCards
	}

	fmt.Println(scratchCardBundle)
	fmt.Printf("Total: %v", total)
}
