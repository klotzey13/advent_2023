package day_2

import (
	steveutil "advent_of_code_2023/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type Cubes struct {
	Colors map[string]int
}

func isDelimiter(r rune) bool {
	return r == ',' || r == ';'
}

func Puzzle_1() {

	filePath := path.Join("day_2", "files", "puzzle.txt")
	f, err := os.Open(filePath)
	steveutil.Check(err)
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
			steveutil.Check(err)
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
