package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func isDelimiter(r rune) bool {
	return r == ':' || r == '|'
}

type ScratchCardsBundle struct {
	totalScratchCards int
	winningMatches    int
}

func main() {

	f, err := os.Open("../input/puzzle.txt")
	checkError(err)
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
