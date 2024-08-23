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

func Puzzle_2() {

	filePath := path.Join("day_2", "files", "puzzle.txt")
	f, err := os.Open(filePath)
	steveutil.Check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	sumOfPowers := 0

	for s.Scan() {
		line := s.Text()
		game := strings.Split(line, ":")
		cubes := game[1]

		cubesToSum := strings.FieldsFunc(cubes, isDelimiter)
		minimumCubes := Cubes{Colors: map[string]int{"red": 0, "blue": 0, "green": 0}}
		minimumCubesPower := 1
		for _, cube := range cubesToSum {
			cubeFields := strings.Fields(cube)
			numberOfCubes, err := strconv.Atoi(cubeFields[0])
			steveutil.Check(err)
			colorOfCube := cubeFields[1]
			if minimumCubes.Colors[colorOfCube] < numberOfCubes {
				minimumCubes.Colors[colorOfCube] = numberOfCubes
			}
		}

		for _, minCube := range minimumCubes.Colors {
			minimumCubesPower *= minCube
		}
		sumOfPowers += minimumCubesPower
	}

	fmt.Println(sumOfPowers)

}
