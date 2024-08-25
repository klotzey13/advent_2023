package day_5

import (
	steveutil "advent_of_code_2023/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type SourceToDestinationMapper struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

func Puzzle_2() {
	startTime := time.Now()

	filePath := path.Join("day_5", "input", "puzzle.txt")
	file, err := os.Open(filePath)
	steveutil.Check(err)
	defer file.Close()

	s := bufio.NewScanner(file)

	seeds := []int{}
	mappings := make(map[string][]SourceToDestinationMapper)
	currentMap := ""
	for s.Scan() {
		line := s.Text()

		if strings.HasPrefix(line, "seeds:") {
			seedStrings := strings.Fields(strings.TrimPrefix(line, "seeds:"))
			for _, seedStr := range seedStrings {
				seed, _ := strconv.Atoi(seedStr)
				seeds = append(seeds, seed)
			}
		}

		if strings.HasSuffix(line, "map:") {
			currentMap = strings.TrimSuffix(line, " map:")
			mappings[currentMap] = []SourceToDestinationMapper{}
		}

		// Parse the mapping lines
		if currentMap != "" {
			fields := strings.Fields(line)
			if len(fields) == 3 {
				dstart, _ := strconv.Atoi(fields[0])
				sstart, _ := strconv.Atoi(fields[1])
				rlength, _ := strconv.Atoi(fields[2])
				mappings[currentMap] = append(mappings[currentMap], SourceToDestinationMapper{destinationStart: dstart, sourceStart: sstart, rangeLength: rlength})
			}
		}
	}

	newSeeds := [][]int{}
	// Map Seeds to low/min for each seed
	for i := 0; i < len(seeds); i = i + 2 {
		min := seeds[i]
		max := seeds[i] + (seeds[i+1] - 1)
		newSeeds = append(newSeeds, []int{min, max})
	}

	locationMapper := mappings["humidity-to-location"]
	locations := [][]int{}
	getMinMaxAndSort(&locations, locationMapper)

	tempuratureMapper := mappings["temperature-to-humidity"]
	tempurature := [][]int{}
	getMinMaxAndSort(&tempurature, tempuratureMapper)

	lightMapper := mappings["light-to-temperature"]
	light := [][]int{}
	getMinMaxAndSort(&light, lightMapper)

	waterMapper := mappings["water-to-light"]
	water := [][]int{}
	getMinMaxAndSort(&water, waterMapper)

	fertilizerMapper := mappings["fertilizer-to-water"]
	fertilizer := [][]int{}
	getMinMaxAndSort(&fertilizer, fertilizerMapper)

	soilMapper := mappings["soil-to-fertilizer"]
	soil := [][]int{}
	getMinMaxAndSort(&soil, soilMapper)

	seedMapper := mappings["seed-to-soil"]
	mseeds := [][]int{}
	getMinMaxAndSort(&mseeds, seedMapper)

	sort(&newSeeds)

	lowestFound := false
	for i := 0; !lowestFound; i++ {
		numCrawl := i
		crawlNumber(&numCrawl, &locations)
		crawlNumber(&numCrawl, &tempurature)
		crawlNumber(&numCrawl, &light)
		crawlNumber(&numCrawl, &water)
		crawlNumber(&numCrawl, &fertilizer)
		crawlNumber(&numCrawl, &soil)
		crawlNumber(&numCrawl, &mseeds)

		for _, s := range newSeeds {
			if numCrawl >= s[0] && numCrawl <= s[1] {
				lowestFound = true
				fmt.Println(i)
			}
		}
	}

	elapsedTime := time.Since(startTime)
	fmt.Println("Runtime:", elapsedTime)

}

func getMinMaxAndSort(l *[][]int, m []SourceToDestinationMapper) {
	for i := 0; i < len(m); i++ {
		min := m[i].destinationStart
		max := m[i].destinationStart + (m[i].rangeLength - 1)
		mapNumber := m[i].sourceStart - m[i].destinationStart
		*l = append(*l, []int{min, max, mapNumber})
	}
	sort(l)
}

func sort(l *[][]int) {
	n := len(*l)
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < n-1; i++ {
			if (*l)[i][0] > (*l)[i+1][0] {
				(*l)[i], (*l)[i+1] = (*l)[i+1], (*l)[i]
				swapped = true
			}
		}
		n--
	}
}

func crawlNumber(numCrawl *int, mapper *[][]int) {
	low, high := 0, len(*mapper)-1
	for low <= high {
		mid := low + (high-low)/2
		m := (*mapper)[mid]
		if *numCrawl >= m[0] && *numCrawl <= m[1] {
			*numCrawl = *numCrawl + m[2]
			return
		} else if *numCrawl < m[0] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}
