package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	filePath := "./challenge_1/input.txt"

	f, err := os.Open(filePath)
	check(err)
	reader := bufio.NewReader(f)

	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanLines)

	numbersToAdd := make([]int, 0)

	numberStringFirstChars := []rune{'o', 't', 'f', 's', 'e', 'n'}

	for scanner.Scan() {
		line := scanner.Text()

		var firstNumber rune
		var secondNumber rune

		for i := 0; i < len(line); i++ {
			if unicode.IsDigit(rune(line[i])) {
				if firstNumber == 0 {
					firstNumber = rune(line[i])
				} else {
					secondNumber = rune(line[i])
				}
			} else if runeInArray(rune(line[i]), numberStringFirstChars) {
				r := convertStringNumToRune(spliceString(line, i, len(line)))
				if r != 0 {
					if firstNumber == 0 {
						firstNumber = r
					} else {
						secondNumber = r
					}
				}
			}
		}

		if secondNumber == 0 {
			secondNumber = firstNumber
		}

		firstNumberString := string(firstNumber)
		secondNumberString := string(secondNumber)

		numberString := firstNumberString + secondNumberString
		number, err := strconv.Atoi(numberString)
		fmt.Printf("%v\n", number)

		if err == nil {
			numbersToAdd = append(numbersToAdd, number)
		}
	}

	total := 0

	for _, number := range numbersToAdd {
		total += number
	}

	fmt.Printf("%v", total)

}

func convertStringNumToRune(splicedString string) rune {

	lengthOfString := len(splicedString)

	if lengthOfString < 3 {
		return 0
	}

	if rune(splicedString[0]) == 'o' {
		if splicedString[1] == 'n' && splicedString[2] == 'e' {
			return '1'
		}
	}

	if rune(splicedString[0]) == 't' {
		if lengthOfString >= 5 {
			if splicedString[1] == 'h' && splicedString[2] == 'r' && splicedString[3] == 'e' && splicedString[4] == 'e' {
				return '3'
			}
		}

		if splicedString[1] == 'w' && splicedString[2] == 'o' {
			return '2'
		}

	}

	if rune(splicedString[0]) == 'f' {
		if lengthOfString >= 4 {
			if splicedString[1] == 'o' {
				if splicedString[2] == 'u' && splicedString[3] == 'r' {
					return '4'
				}
			} else if splicedString[1] == 'i' {
				if splicedString[2] == 'v' && splicedString[3] == 'e' {
					return '5'
				}
			}
		}
	}

	if rune(splicedString[0]) == 's' {
		if lengthOfString >= 5 {
			if splicedString[1] == 'e' && splicedString[2] == 'v' && splicedString[3] == 'e' && splicedString[4] == 'n' {
				return '7'
			}
		}
		if splicedString[1] == 'i' && splicedString[2] == 'x' {
			return '6'
		}

	}

	if rune(splicedString[0]) == 'e' {
		if lengthOfString >= 5 {
			if splicedString[1] == 'i' && splicedString[2] == 'g' && splicedString[3] == 'h' && splicedString[4] == 't' {
				return '8'
			}
		}
	}

	if rune(splicedString[0]) == 'n' {
		if lengthOfString >= 4 {
			if splicedString[1] == 'i' && splicedString[2] == 'n' && splicedString[3] == 'e' {
				return '9'
			}
		}
	}

	return 0

}

func spliceString(s string, from int, to int) string {
	runeSlice := []rune(s)
	return string(runeSlice[from:to])
}

func runeInArray(target rune, array []rune) bool {
	for _, r := range array {
		if r == target {
			return true
		}
	}
	return false
}
