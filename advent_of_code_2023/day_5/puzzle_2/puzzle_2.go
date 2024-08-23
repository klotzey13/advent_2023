package day_5

import (
	steveutil "advent_of_code_2023/util"
	"bufio"
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
)

type SourceToDestinationMapper struct {
	destinationStart int
	sourceStart      int
	rangeLength      int
}

type MappedSeed struct {
	seed        int
	soil        int
	fertilizer  int
	water       int
	light       int
	tempurature int
	humidity    int
	location    int
}

func Puzzle_2() {
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

	newSeeds := []int{}
	for i := 0; i < len(seeds); i = i + 2 {
		for j := 0; j < seeds[i+1]; j++ {
			newSeeds = append(newSeeds, seeds[i]+j)
		}
	}
	fmt.Println(newSeeds)
	mappedSeeds := mapSeeds(newSeeds, mappings)
	fmt.Println(calculateLowestLocation(mappedSeeds))
}

func mapSeeds(seeds []int, mappings map[string][]SourceToDestinationMapper) []MappedSeed {
	mappedSeeds := []MappedSeed{}

	for i := 0; i < len(seeds); i++ {
		mappedSeed := MappedSeed{seed: seeds[i]}

		seedToSoilMapping := mappings["seed-to-soil"]
		for m := 0; m < len(seedToSoilMapping); m++ {
			calculatedRange := seedToSoilMapping[m].rangeLength + seedToSoilMapping[m].sourceStart

			for s := seedToSoilMapping[m].sourceStart; s < calculatedRange; s++ {
				if seeds[i] == s {
					mappedSeed.soil = seeds[i] + (seedToSoilMapping[m].destinationStart - seedToSoilMapping[m].sourceStart)
					break
				}
			}
			if mappedSeed.soil == 0 {
				mappedSeed.soil = seeds[i]
			}
		}

		soilToFertilizerMapping := mappings["soil-to-fertilizer"]
		for m := 0; m < len(soilToFertilizerMapping); m++ {
			calculatedRange := soilToFertilizerMapping[m].rangeLength + soilToFertilizerMapping[m].sourceStart

			for s := soilToFertilizerMapping[m].sourceStart; s < calculatedRange; s++ {
				if mappedSeed.soil == s {
					mappedSeed.fertilizer = mappedSeed.soil + (soilToFertilizerMapping[m].destinationStart - soilToFertilizerMapping[m].sourceStart)
					break
				}
			}
			if mappedSeed.fertilizer == 0 {
				mappedSeed.fertilizer = mappedSeed.soil
			}
		}

		fertilizerToWaterMapping := mappings["fertilizer-to-water"]
		for m := 0; m < len(fertilizerToWaterMapping); m++ {

			calculatedRange := fertilizerToWaterMapping[m].rangeLength + fertilizerToWaterMapping[m].sourceStart

			for s := fertilizerToWaterMapping[m].sourceStart; s < calculatedRange; s++ {
				if mappedSeed.fertilizer == s {
					mappedSeed.water = mappedSeed.fertilizer + (fertilizerToWaterMapping[m].destinationStart - fertilizerToWaterMapping[m].sourceStart)
					break
				}
			}
			if mappedSeed.water == 0 {
				mappedSeed.water = mappedSeed.fertilizer
			}
		}

		waterToLightMapping := mappings["water-to-light"]
		for m := 0; m < len(waterToLightMapping); m++ {

			calculatedRange := waterToLightMapping[m].rangeLength + waterToLightMapping[m].sourceStart

			for s := waterToLightMapping[m].sourceStart; s < calculatedRange; s++ {
				if mappedSeed.water == s {
					mappedSeed.light = mappedSeed.water + (waterToLightMapping[m].destinationStart - waterToLightMapping[m].sourceStart)
					break
				}
			}
			if mappedSeed.light == 0 {
				mappedSeed.light = mappedSeed.water
			}
		}

		lightToTempuratureMapping := mappings["light-to-temperature"]
		for m := 0; m < len(lightToTempuratureMapping); m++ {
			calculatedRange := lightToTempuratureMapping[m].rangeLength + lightToTempuratureMapping[m].sourceStart
			for s := lightToTempuratureMapping[m].sourceStart; s < calculatedRange; s++ {
				if mappedSeed.light == s {
					mappedSeed.tempurature = mappedSeed.light + (lightToTempuratureMapping[m].destinationStart - lightToTempuratureMapping[m].sourceStart)
					break
				}
			}
			if mappedSeed.tempurature == 0 {
				mappedSeed.tempurature = mappedSeed.light
			}
		}

		temperatureToHumidityMapping := mappings["temperature-to-humidity"]
		for m := 0; m < len(temperatureToHumidityMapping); m++ {
			calculatedRange := temperatureToHumidityMapping[m].rangeLength + temperatureToHumidityMapping[m].sourceStart
			for s := temperatureToHumidityMapping[m].sourceStart; s < calculatedRange; s++ {
				if mappedSeed.tempurature == s {
					mappedSeed.humidity = mappedSeed.tempurature + (temperatureToHumidityMapping[m].destinationStart - temperatureToHumidityMapping[m].sourceStart)
					break
				}
				if mappedSeed.humidity == 0 {
					mappedSeed.humidity = mappedSeed.tempurature
				}
			}

		}

		humidityToLocationMapping := mappings["humidity-to-location"]
		for m := 0; m < len(humidityToLocationMapping); m++ {
			calculatedRange := humidityToLocationMapping[m].rangeLength + humidityToLocationMapping[m].sourceStart
			for s := humidityToLocationMapping[m].sourceStart; s < calculatedRange; s++ {
				if mappedSeed.humidity == s {
					mappedSeed.location = mappedSeed.humidity + (humidityToLocationMapping[m].destinationStart - humidityToLocationMapping[m].sourceStart)
					break
				}
				if mappedSeed.location == 0 {
					mappedSeed.location = mappedSeed.humidity
				}
			}
		}
		fmt.Println(mappedSeed)
		mappedSeeds = append(mappedSeeds, mappedSeed)
	}

	return mappedSeeds
}

func calculateLowestLocation(mappedSeeds []MappedSeed) int {
	lowest := mappedSeeds[0].location
	for i := 1; i < len(mappedSeeds); i++ {
		if mappedSeeds[i].location < lowest {
			lowest = mappedSeeds[i].location
		}
	}

	return lowest
}
