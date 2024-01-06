package main

import (
	"fmt"
	"os"
	"strings"
)

func day4_2() {
	cards := verifyManyCards()
	scratchCards := [212]int{}
	for i := range scratchCards {
		scratchCards[i] = 1
	}

	var sum int
	for _, card_ := range cards {
		for j := 0; j < scratchCards[card_.id]; j++ {
			for i := 1; i <= card_.victoryNum; i++ {
				scratchCards[card_.id+i]++
			}
		}
	}
	for _, card_ := range scratchCards {
		sum += card_
	}
	fmt.Println(sum)
}

type card struct {
	id         int
	victoryNum int
}

func verifyManyCards() []card {
	dat, err := os.ReadFile("./inputs/day4.txt")

	if err != nil {
		panic(err)
	}

	text := strings.Split(string(dat), "\n")
	cards := []card{}

	for id, line := range text {
		lists := strings.Split(line, ": ")[1]
		separatedNums := strings.Split(lists, " | ")
		victoryNums := separatedNums[0]
		myNums := separatedNums[1]

		rightNums := 0
		var actualCard card
		for _, num := range strings.Split(myNums, " ") {
			for _, victoryNum := range strings.Split(victoryNums, " ") {
				if num == victoryNum && victoryNum != "" && num != "" {
					rightNums++
				}
			}
		}
		actualCard.id = id
		actualCard.victoryNum = rightNums
		cards = append(cards, actualCard)
	}
	return cards
}
