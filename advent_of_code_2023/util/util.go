package steveutil

import (
	"bufio"
	"os"
)

// ReadInput reads the contents of a file and returns them as a string
func ReadInput(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

// ReadLines reads the lines of a file and returns them as a slice of strings
func ReadLines(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

type Comparable interface {
	Compare(other Comparable) int
}

func BinarySearch[T Comparable](slice []T, target T) int {
	low, high := 0, len(slice)-1

	for low <= high {
		mid := low + (high-low)/2 // Avoid potential overflow

		cmp := target.Compare(slice[mid])
		if cmp < 0 {
			high = mid - 1
		} else if cmp > 0 {
			low = mid + 1
		} else {
			return mid
		}
	}

	return -1
}
