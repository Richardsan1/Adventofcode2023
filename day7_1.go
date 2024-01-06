package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	id    int
}

var rank = map[byte]int{
	'2': 0, '3': 1, '4': 2, '5': 3, '6': 4, '7': 5, '8': 6, '9': 7, 'T': 8, 'J': 9, 'Q': 10, 'K': 11, 'A': 12,
}

type ByRank []hand

func (a ByRank) Len() int           { return len(a) }
func (a ByRank) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRank) Less(i, j int) bool { return comparaStrings(a[i].cards, a[j].cards) }

func comparaStrings(str1, str2 string) bool {
	for k := 0; k < len(str1); k++ {
		if rank[str1[k]] < rank[str2[k]] {
			return true
		} else if rank[str1[k]] > rank[str2[k]] {
			return false
		}
	}
	return false
}

func sortHands(cartas []hand) []hand {

	// Ordena a slice usando a implementação personalizada
	sort.Sort(ByRank(cartas))

	// Imprime o resultado ordenado
	return cartas
}

func day7_1() {
	dat, _ := os.ReadFile("./inputs/day7.txt")
	text := strings.Split(string(dat), "\n")

	hands := make([][]hand, 7)

	var sum int

	for id, line := range text {
		line_ := strings.Split(line, " ")

		force := verifyHand(line_[0])
		hands[force-1] = append(hands[force-1], hand{line_[0], id})
	}

	power := 1
	for i := 0; i < 7; i++ {
		for _, mao := range sortHands(hands[i]) {
			bet, _ := strconv.Atoi(strings.Split(text[mao.id], " ")[1])
			sum += (power * bet)
			power++
		}
	}
	fmt.Println(sum)

}

func verifyHand(hand string) int {
	var cards [13]int

	for _, card := range hand {
		switch card {
		case '2':
			cards[0]++
		case '3':
			cards[1]++
		case '4':
			cards[2]++
		case '5':
			cards[3]++
		case '6':
			cards[4]++
		case '7':
			cards[5]++
		case '8':
			cards[6]++
		case '9':
			cards[7]++
		case 'T':
			cards[8]++
		case 'J':
			cards[9]++
		case 'Q':
			cards[10]++
		case 'K':
			cards[11]++
		case 'A':
			cards[12]++
		}
	}

	var games [4]int
	for _, val := range cards {
		switch val {
		case 1:
			games[0]++
		case 2:
			games[1]++
		case 3:
			games[2]++
		case 4:
			return 6
		case 5:
			return 7
		}
	}

	if games[2] == 1 {
		if games[1] == 1 {
			return 5
		} else {
			return 4
		}
	} else if games[1] >= 1 {
		if games[1] == 2 {
			return 3
		} else {
			return 2
		}
	} else {
		return 1
	}
}
