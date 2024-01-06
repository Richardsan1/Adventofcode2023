package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day2_1() {
	dat, err := os.ReadFile("inputs/day2.txt")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(dat), "\n")
	var sum int
	for id, line := range lines {
		line_ := strings.Split(line, ":")
		id = id + 1
		games := strings.Split(line_[1], ";")
		isValid := true

		for _, game := range games {
			for _, roll := range strings.Split(game, ",") {
				roll_ := strings.Split(roll, " ")
				num, _ := strconv.Atoi(roll_[1])
				if roll_[2] == "red" && num > 12 {
					isValid = false
					break
				} else if roll_[2] == "green" && num > 13 {
					isValid = false
					break
				} else if roll_[2] == "blue" && num > 14 {
					isValid = false
					break
				}
			}
			if !isValid {
				break
			}
		}
		if isValid {
			sum += id
		}
	}
	fmt.Println(sum)
}
