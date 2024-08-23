package day_3

import (
	steveutil "advent_of_code_2023/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Numbers struct {
	lineIndex int
	number    string
	indexs    []int
}

type SymbolChecker struct {
	lineIndex     int
	index         int
	indexsToCheck []int
}

func Puzzle_2() {
	filepath := path.Join("day_3", "input", "puzzle.txt")
	f, err := os.Open(filepath)
	steveutil.Check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	numbers := []Numbers{}
	symbolChecker := []SymbolChecker{}
	lineIndex := 0

	for s.Scan() {
		line := s.Text()
		chars := strings.Split(line, "")
		stringLength := len(chars)

		for i := 0; i < stringLength; i++ {
			_, err := strconv.Atoi(chars[i])
			if err == nil {
				numberToCheck := chars[i]
				indexsForNumber := []int{}
				indexsForNumber = append(indexsForNumber, i)
				stillHaveDigits := true

				for x := i + 1; stillHaveDigits && x < stringLength; x++ {
					_, err := strconv.Atoi(chars[x])
					indexsForNumber = append(indexsForNumber, x)
					if err == nil {
						numberToCheck += chars[x]
					} else {
						stillHaveDigits = false
						i = x
					}
				}

				numbers = append(numbers, Numbers{
					lineIndex: lineIndex,
					number:    numberToCheck,
					indexs:    indexsForNumber,
				})
			} else if chars[i] == "*" {
				indexsToCheck := []int{}
				if i > 0 {
					indexsToCheck = append(indexsToCheck, i-1)
				}

				indexsToCheck = append(indexsToCheck, i)

				if i < stringLength {
					indexsToCheck = append(indexsToCheck, i+1)
				}

				symbolChecker = append(symbolChecker, SymbolChecker{
					index:         i,
					indexsToCheck: indexsToCheck,
					lineIndex:     lineIndex,
				})
			}
		}
		lineIndex++
	}

	fmt.Printf("Total symbols found: %d\n", len(symbolChecker))
	fmt.Printf("Total numbers found: %d\n", len(numbers))

	numbersToSum := 0

	for _, symbol := range symbolChecker {
		foundGearOnThisLine := false

		for _, symbolIndexToCheck := range symbol.indexsToCheck {
			if foundGearOnThisLine {
				break
			}
			numbersFound := 0
			numbersToMultiply := []int{}

			searchAbove := max(0, symbol.lineIndex-1)
			searchBelow := min(lineIndex, symbol.lineIndex+1)

			for _, num := range numbers {
				if num.lineIndex >= searchAbove && num.lineIndex <= searchBelow {
					for _, nIndex := range num.indexs {
						if nIndex == symbolIndexToCheck {
							numberToAdd, err := strconv.Atoi(num.number)
							steveutil.Check(err)
							numbersToMultiply = append(numbersToMultiply, numberToAdd)
							numbersFound++
							break
						}
					}
					if numbersFound > 2 {
						break
					}
				}
			}
			if numbersFound > 2 {
				break
			}
			if numbersFound == 2 {

				foundGearOnThisLine = true

				gearRatio := numbersToMultiply[0] * numbersToMultiply[1]
				numbersToSum += gearRatio
				fmt.Printf("Gear found at line %d: %d * %d = %d\n",
					symbol.lineIndex, numbersToMultiply[0], numbersToMultiply[1], gearRatio)
			}
		}
	}

	// Not finding gears correctly
	fmt.Printf("Final sum: %d\n", numbersToSum)
}
