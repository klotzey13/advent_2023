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

type Cubes struct {
	Colors map[string]int
}

func isDelimiter(r rune) bool {
	return r == ',' || r == ';'
}

func main() {

	filePath := "../files/puzzle.txt"
	f, err := os.Open(filePath)
	check(err)
	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	redLimit := 12
	greenLimit := 13
	blueLimit := 14

	sumOfIds := 0

	puzzleIndex := 1
	for s.Scan() {
		gameCubeCount := Cubes{Colors: map[string]int{"red": redLimit, "blue": blueLimit, "green": greenLimit}}
		line := s.Text()
		game := strings.Split(line, ":")
		cubes := game[1]
		fmt.Println(cubes)

		cubesToSum := strings.FieldsFunc(cubes, isDelimiter)
		shouldAddIndex := true
		fmt.Println(cubesToSum)

		for _, cube := range cubesToSum {
			cubeFields := strings.Fields(cube)
			numberOfCubes, err := strconv.Atoi(cubeFields[0])
			check(err)
			colorOfCube := cubeFields[1]
			if numberOfCubes > gameCubeCount.Colors[colorOfCube] {
				fmt.Println("failed")
				shouldAddIndex = false
				break
			}
		}
		if shouldAddIndex {
			fmt.Println(puzzleIndex)
			sumOfIds += puzzleIndex
		}

		fmt.Println(puzzleIndex)
		puzzleIndex++
	}

	fmt.Println(sumOfIds)

}
