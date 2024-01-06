package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var rankJoker = map[byte]int{
	'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8, 'T': 9, 'Q': 10, 'K': 11, 'A': 12,
}

type ByRankJoker []hand

func (a ByRankJoker) Len() int           { return len(a) }
func (a ByRankJoker) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByRankJoker) Less(i, j int) bool { return compareStringsJoker(a[i].cards, a[j].cards) }

func compareStringsJoker(str1, str2 string) bool {
	for k := 0; k < len(str1); k++ {
		if rankJoker[str1[k]] < rankJoker[str2[k]] {
			return true
		} else if rankJoker[str1[k]] > rankJoker[str2[k]] {
			return false
		}
	}
	return false
}

func sortHandsJoker(cartas []hand) []hand {
	sort.Sort(ByRankJoker(cartas))
	return cartas
}

func day7_2() {
	dat, _ := os.ReadFile("./inputs/day7.txt")
	text := strings.Split(string(dat), "\n")

	hands := make([][]hand, 7)

	var sum int

	for id, line := range text {
		line_ := strings.Split(line, " ")

		force := verifyHandJoker(line_[0])
		hands[force-1] = append(hands[force-1], hand{line_[0], id})
	}

	power := 1
	for i := 0; i < 7; i++ {
		for _, mao := range sortHandsJoker(hands[i]) {
			bet, _ := strconv.Atoi(strings.Split(text[mao.id], " ")[1])
			sum += (power * bet)
			power++
		}
	}
	fmt.Println(sum)

}

func verifyHandJoker(hand string) int {
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

	joker := cards[9]

	var games [5]int
	for id, val := range cards {
		if id == 9 {
			continue
		}
		switch val {
		case 5:
			games[4]++
		case 4:
			games[3]++
		case 3:
			games[2]++
		case 2:
			games[1]++
		case 1:
			games[0]++
		}
	}
	// quinta
	if joker > 0 {
		if joker >= 4 {
			return 7
		} else if joker == 3 {
			// quinta
			if games[1] == 1 {
				return 7
				// quadra
			} else {
				return 6
			}
		} else if joker == 2 {
			// quinta
			if games[2] == 1 {
				return 7
				// quadra
			} else if games[1] == 1 {
				return 6
			} else {
				return 4
			}
		} else if joker == 1 {
			if games[3] == 1 {
				return 7
			} else if games[2] == 1 {
				return 6
			} else if games[1] >= 1 {
				if games[1] == 2 {
					return 5
				} else {
					return 4
				}
			} else {
				return 2
			}
		}
	} else {
		if games[4] == 1 {
			return 7
		} else if games[3] == 1 {
			return 6
		} else if games[2] == 1 {
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
	return 0
}
