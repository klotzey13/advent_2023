package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type NumberChecker struct {
	lineIndex     int
	number        string
	indexsToCheck []int
}

type Symbols struct {
	lineIndex int
	indexs    []int
}

func main() {

	f, err := os.Open("./input/example.txt")
	check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	numbersToCheck := []NumberChecker{}
	symbolsFound := []Symbols{}
	lineIndex := 0

	for s.Scan() {
		line := s.Text()
		chars := strings.Split(line, "")
		stringLength := len(chars)
		symbolsForLine := Symbols{
			lineIndex: lineIndex,
			indexs:    []int{},
		}

		for i := 0; i < stringLength; i++ {
			_, err := strconv.Atoi(chars[i])
			if err == nil {
				numberToCheck := chars[i]
				indexsForNumber := []int{}
				indexsForNumber = append(indexsForNumber, i)
				if i != 0 {
					indexsForNumber = append(indexsForNumber, i-1)
				}
				stillHaveDigits := true

				for x := i + 1; stillHaveDigits && x < stringLength; x++ {
					_, err := strconv.Atoi(chars[x])
					indexsForNumber = append(indexsForNumber, x)
					if err == nil {
						numberToCheck += chars[x]
					} else {
						stillHaveDigits = false
						i = x - 1
					}
				}

				numbersToCheck = append(numbersToCheck, NumberChecker{
					lineIndex:     lineIndex,
					number:        numberToCheck,
					indexsToCheck: indexsForNumber,
				})
			} else if chars[i] != "." {
				symbolsForLine.indexs = append(symbolsForLine.indexs, i)
			}
		}
		symbolsFound = append(symbolsFound, symbolsForLine)
		lineIndex++
	}

	numbersToSum := 0

	for _, num := range numbersToCheck {
		if num.lineIndex == 0 {
			searchBelow := 1
			for _, symbol := range symbolsFound {
				if symbol.lineIndex == searchBelow || symbol.lineIndex == num.lineIndex {
					for _, sindex := range symbol.indexs {
						for _, nindex := range num.indexsToCheck {
							if sindex == nindex {
								numberToAdd, err := strconv.Atoi(num.number)
								check(err)
								fmt.Printf("Adding %v\n", numberToAdd)
								numbersToSum += numberToAdd
								break
							}
						}

					}
				}
			}
		} else if num.lineIndex == len(numbersToCheck) {
			searchAbove := num.lineIndex - 1
			for _, symbol := range symbolsFound {
				if symbol.lineIndex == searchAbove || symbol.lineIndex == num.lineIndex {
					for _, sindex := range symbol.indexs {
						for _, nindex := range num.indexsToCheck {
							if sindex == nindex {
								numberToAdd, err := strconv.Atoi(num.number)
								check(err)
								fmt.Printf("Adding %v\n", numberToAdd)
								numbersToSum += numberToAdd
								break
							}
						}

					}
				}
			}
		} else {
			searchBelow := num.lineIndex + 1
			searchAbove := num.lineIndex - 1

			for _, symbol := range symbolsFound {
				if symbol.lineIndex == searchAbove || symbol.lineIndex == searchBelow || symbol.lineIndex == num.lineIndex {
					for _, sindex := range symbol.indexs {
						for _, nindex := range num.indexsToCheck {
							if sindex == nindex {
								numberToAdd, err := strconv.Atoi(num.number)
								check(err)
								fmt.Printf("Adding %v\n", numberToAdd)
								numbersToSum += numberToAdd
								break
							}
						}

					}
				}
			}
		}
	}

	fmt.Println(numbersToCheck)
	fmt.Println(symbolsFound)
	fmt.Println()
	fmt.Println(numbersToSum)
}
