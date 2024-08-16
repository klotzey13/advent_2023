package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// Converts words to their numerical equivalents
var numberMap = map[string]string{
	"one": "1", "two": "2", "three": "3", "four": "4", "five": "5",
	"six": "6", "seven": "7", "eight": "8", "nine": "9",
}

func main() {
	file, err := os.Open("/challenge_1/example.txt") // Replace "input.txt" with your actual input file
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalSum := 0

	for scanner.Scan() {
		line := scanner.Text()
		firstDigit, lastDigit := findDigits(line)
		value, _ := strconv.Atoi(firstDigit + lastDigit) // Combine and convert to int
		totalSum += value
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	fmt.Println("Total Sum:", totalSum)
}

// Finds the first and last "digit" in a line, converting words to numbers as necessary
func findDigits(line string) (string, string) {
	// Regular expression to match words and digits
	re := regexp.MustCompile(`\b(one|two|three|four|five|six|seven|eight|nine|\d)\b`)
	matches := re.FindAllString(line, -1)

	first, last := matches[0], matches[len(matches)-1]
	first = convertWordToDigit(first)
	last = convertWordToDigit(last)

	return first, last
}

// Converts a word representing a number to its numerical equivalent, if necessary
func convertWordToDigit(word string) string {
	if val, exists := numberMap[word]; exists {
		return val
	}
	return word // Return the word unchanged if it's already a digit
}
