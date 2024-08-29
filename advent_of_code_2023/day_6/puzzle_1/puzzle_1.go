package day_6

import (
	steveutil "advent_of_code_2023/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

// Each line of file contains time of race + record distance
// User holds buutton for x ms, and boat speed increases by x mm/ms
// x * remaining time = total distance
// For each race, count each value of x where total distance > race record
// Multiply counts for each race to return margin of error

type RaceStats struct {
	time     int
	distance int
}

func Puzzle_1() {
	filepath := path.Join("day_6", "input", "puzzle.txt")

	f, err := os.Open(filepath)
	steveutil.Check(err)
	defer f.Close()

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	races := []RaceStats{}
	index := 0
	for s.Scan() {
		line := s.Text()
		if index == 0 {
			times := parseValues(line)
			for _, time := range times {
				timeInt, err := strconv.Atoi(time)
				steveutil.Check(err)
				timeObj := RaceStats{time: timeInt}
				races = append(races, timeObj)
			}
		}
		if index == 1 {
			distances := parseValues(line)
			for i, distance := range distances {
				distanceInt, err := strconv.Atoi(distance)
				steveutil.Check(err)
				races[i].distance = distanceInt
			}
		}
		index++
	}

	marginError := 0
	for _, race := range races {
		wins := getRaceWins(race)
		if marginError == 0 {
			marginError = wins
		} else {
			marginError *= wins
		}
	}

	fmt.Println(marginError)
}

func Puzzle_2() {
	filepath := path.Join("day_6", "input", "puzzle.txt")

	f, err := os.Open(filepath)
	steveutil.Check(err)
	defer f.Close()
	steveutil.Check(err)

	r := bufio.NewReader(f)
	s := bufio.NewScanner(r)

	var singleRace RaceStats
	index := 0
	for s.Scan() {
		line := s.Text()
		if index == 0 {
			singleRace.time = populateLine(line)
		}
		if index == 1 {
			singleRace.distance = populateLine(line)
		}
		index++
	}
	wins := getRaceWins(singleRace)

	fmt.Println(wins)
}

func parseValues(line string) []string {
	_, after, _ := strings.Cut(line, ": ")
	return strings.Fields(after)
}

func getRaceWins(race RaceStats) int {
	wins := 0
	for i := range race.time {
		multiplier := i + 1
		if multiplier*(race.time-multiplier) > race.distance {
			wins++
		}
	}
	return wins
}

func populateLine(line string) int {
	_, after, _ := strings.Cut(line, ": ")
	trimmed := strings.ReplaceAll(after, " ", "")
	outputInt, err := strconv.Atoi(trimmed)
	steveutil.Check(err)
	return outputInt
}
