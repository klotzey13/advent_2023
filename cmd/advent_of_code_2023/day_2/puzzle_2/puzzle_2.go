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
			check(err)
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
