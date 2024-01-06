package main

import (
	"fmt"
	"os"
	"strings"
)

func day1_2() {
	dat, _ := os.ReadFile("./inputs/day1.txt")

	convertedDat := string(dat)
	lines := strings.Split(convertedDat, "\n")
	var sum int

	for _, line := range lines {
		var first int = -1
		var last int = 0
		textFirst, textLast := findNumberOccurrences(line)

		for j := 0; j < len(line); j++ {
			if int(line[j]) >= 47 && int(line[j]) <= 57 {
				if first == -1 {
					first = int(line[j]) - 48
				}
				if j > textLast.Index {
					last = int(line[j]) - 48
				} else {
					last = textLast.Number
				}
			} else if j > textFirst.Index {
				if first == -1 {
					first = textFirst.Number
				}
			}
		}
		sum += first*10 + last
	}
	fmt.Print(sum)
}

type NumberOccurrence struct {
	Number int
	Index  int
}

func findNumberOccurrences(s string) (NumberOccurrence, NumberOccurrence) {
	numbers := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	var occurrences []NumberOccurrence

	for pos, number := range numbers {
		i := 0
		for {
			next := strings.Index(s[i:], number)
			if next == -1 {
				break
			}
			i += next
			occurrences = append(occurrences, NumberOccurrence{Number: pos, Index: i})
			i += len(number)
		}
	}
	var first, last NumberOccurrence
	first.Index = 999999999999999999
	last.Index = -1
	for _, occ := range occurrences {
		if occ.Index < first.Index {
			first.Index = occ.Index
			first.Number = occ.Number
		}

		if occ.Index > last.Index {
			last.Index = occ.Index
			last.Number = occ.Number
		}
	}

	return first, last
}
