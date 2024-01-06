package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day5_1() {
	seeds, seed_to_soil, soil_to_fert, fert_to_water, water_to_light, light_to_temp, temp_to_humi, humi_to_locat := extractData()

	seeds = strings.Split(seeds[0], ": ")
	seeds = strings.Split(seeds[1], " ")

	var seedsConverted []int
	for _, seed := range seeds {
		seed_ := convertUnit(seed, seed_to_soil)
		seed_ = convertUnit(strconv.Itoa(seed_), soil_to_fert)
		seed_ = convertUnit(strconv.Itoa(seed_), fert_to_water)
		seed_ = convertUnit(strconv.Itoa(seed_), water_to_light)
		seed_ = convertUnit(strconv.Itoa(seed_), light_to_temp)
		seed_ = convertUnit(strconv.Itoa(seed_), temp_to_humi)
		seed_ = convertUnit(strconv.Itoa(seed_), humi_to_locat)

		seedsConverted = append(seedsConverted, seed_)
	}

	var minor int
	for _, seed := range seedsConverted {
		if minor == 0 {
			minor = seed
		}
		if seed < minor {
			minor = seed
		}
	}
	fmt.Println(minor)
}

func extractData() ([]string, []string, []string, []string, []string, []string, []string, []string) {
	dat, _ := os.ReadFile("./inputs/day5.txt")

	text := strings.Split(string(dat), "\n")

	var maps [8][]string
	var count int = 0

	for _, line := range text {

		if line == "" {
			count++
			continue
		}

		if line[len(line)-4:] != "map:" {
			maps[count] = append(maps[count], line)
		}
	}

	return maps[0], maps[1], maps[2], maps[3], maps[4], maps[5], maps[6], maps[7]
}

func convertUnit(seed string, mapConverter []string) int {
	seed_, _ := strconv.Atoi(seed)
	for _, line := range mapConverter {
		line_ := strings.Split(line, " ")

		rangeConvert, _ := strconv.Atoi(line_[0])
		rangeStart, _ := strconv.Atoi(line_[1])
		space, _ := strconv.Atoi(line_[2])

		pos := seed_ - rangeStart

		if pos >= 0 && pos < space {
			seed_ = rangeConvert + pos
			break
		} else {
			continue
		}
	}

	return seed_
}
