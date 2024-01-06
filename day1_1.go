package main

import (
	"fmt"
	"os"
	"strconv"
)

func day1_1() {
	dat, err := os.ReadFile("./inputs/day1.txt")
	if err != nil {
		panic(err)
	}
	var values []int
	var word []byte
	for i := 0; i < len(dat); i++ {
		if string(dat[i]) != "\n" {
			word = append(word, dat[i])

		} else {
			var first byte = 0
			var last byte
			for j := 0; j < len(word); j++ {
				if word[j] >= 47 && word[j] <= 57 {
					if first == 0 {
						first = word[j]
					}
					last = word[j]
				}
			}
			first_, _ := strconv.Atoi(string(first))
			last_, _ := strconv.Atoi(string(last))
			values = append(values, first_*10+last_)

			first = 0
			last = 0

			word = []byte{}
		}
	}
	var sum int
	for i := 0; i < len(values); i++ {
		sum += values[i]
	}
	fmt.Print(sum)
}
