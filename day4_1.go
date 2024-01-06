package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

func day4_1() {
	dat, err := os.ReadFile("./inputs/day4.txt")

	if err != nil {
		panic(err)
	}

	text := strings.Split(string(dat), "\n")
	var sum int = 0
	for _, line := range text {
		lists := strings.Split(line, ": ")[1]
		cards := strings.Split(lists, " | ")
		victoryNums := cards[0]
		myNums := cards[1]

		rightNums := -1.0
		for _, num := range strings.Split(myNums, " ") {
			for _, victoryNum := range strings.Split(victoryNums, " ") {
				if num == victoryNum && victoryNum != "" && num != "" {
					rightNums++
				}
			}
		}
		if rightNums != .5 {
			sum += int(math.Pow(2, rightNums))
		}
	}
	fmt.Println(sum)
}
