package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day6_2() {
	dat, _ := os.ReadFile("./inputs/day6.txt")
	input := strings.Fields(string(dat))
	inputTime, _ := strconv.Atoi(strings.Join(input[1:5], ""))
	inputDistance, _ := strconv.Atoi(strings.Join(input[6:10], ""))

	var sum int = getWinningSequencesOptimized(inputTime, inputDistance)

	fmt.Println(sum)
}

func getWinningSequencesOptimized(time int, record int) int {
	var winningSequences int

	delta := math.Sqrt(math.Pow(float64(-time), 2) - 4.0*float64(record))
	var xmin float64 = (float64(time) - delta) / 2
	var xmax float64 = (float64(time) + delta) / 2

	winningSequences = int(math.Ceil(xmax) - math.Ceil(xmin))

	return winningSequences
}
