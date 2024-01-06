package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day2_2() {
	dat, err := os.ReadFile("inputs/day2.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	var sum int
	for _, line := range lines {
		line_ := strings.Split(line, ":")
		games := strings.Split(line_[1], ";")

		var r, g, b int
		for _, game := range games {

			for _, roll := range strings.Split(game, ",") {
				roll_ := strings.Split(strings.Trim(roll, " "), " ")
				cubes, _ := strconv.Atoi(roll_[0])

				switch roll_[1] {
				case "red":
					if cubes > r {
						r = cubes
					}
				case "green":
					if cubes > g {
						g = cubes
					}
				case "blue":
					if cubes > b {
						b = cubes
					}
				}
			}

		}
		sum += (r * g * b)

	}
	fmt.Println(sum)
}
