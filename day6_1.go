package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day6_1() {
	dat, _ := os.ReadFile("./inputs/day6.txt")
	input := strings.Fields(string(dat))
	inputTime := input[1:5]
	inputDistance := input[6:10]

	var sum int = 1
	for i := 0; i < 4; i++ {
		time_, _ := strconv.Atoi(inputTime[i])
		distance_, _ := strconv.Atoi(inputDistance[i])
		sum *= getWinningSequences(time_, distance_)
	}
	fmt.Println(sum)
}

// f(x) = (x * i) * (t - x)
func getWinningSequences(time int, record int) int {
	var winningSequences int
	var timePressing int = 0
	for i := 0; i < time; i++ {
		val := timePressing * (time - timePressing)
		if val > record {
			winningSequences++
		}
		timePressing++
	}
	return winningSequences
}
