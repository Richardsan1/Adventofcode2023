package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day3_2() {
	dat, err := os.ReadFile("./inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	text := strings.Split(string(dat), "\n")
	var sum int = 0
	// cada linha tem 140 caracteres

	for lineID, line := range text { // para cada linha
		for charID, char := range line { // para cada caractere da linha
			if char == 42 { // verifica se é um *
				var qntNums int = 0
				var nums []int
				if charID > 0 { // verifica se não é o primeiro caractere
					value_ := verifyNumLeft(charID, []byte(line))
					if value_ != 0 {
						nums = append(nums, value_)
						qntNums++
					}
				}
				if charID < 139 { // verifica se não é o último caractere
					value_ := verifyNumRight(charID, []byte(line))
					if value_ != 0 {
						nums = append(nums, value_)
						qntNums++
					}
				}
				if lineID > 0 {
					sum_, qntNums_ := verifyNumDiagonal2([]byte(text[lineID-1]), charID)
					if sum_ != 0 {
						nums = append(nums, sum_)
						qntNums += qntNums_
					}
				}
				if lineID < len(text)-1 {
					sum_, qntNums_ := verifyNumDiagonal2([]byte(text[lineID+1]), charID)
					if sum_ != 0 {
						nums = append(nums, sum_)
						qntNums += qntNums_
					}
				}
				if qntNums == 2 {
					product := 1
					for _, num := range nums {
						product *= num
					}
					if product != 1 {
						sum += product
					}
				}

			}
		}
	}
	fmt.Println(sum)
}

func verifyNumDiagonal2(lineDiagonal []byte, actualId int) (int, int) {
	var nums []int
	var product int
	var quantity int

	if actualId >= 0 && actualId <= 139 {
		if verifyNum(rune(lineDiagonal[actualId])) {
			val := verifyNumLeft(actualId+1, lineDiagonal)
			if len(strconv.Itoa(val)) > 1 {
				val2 := verifyNumLeft(actualId+2, lineDiagonal)
				if len(strconv.Itoa(val2)) > 1 {
					nums = append(nums, val2)
					quantity++
				} else {
					nums = append(nums, val)
					quantity++
				}
			} else {
				nums = append(nums, verifyNumRight(actualId-1, lineDiagonal))
				quantity++
			}

		} else {
			valLeft := verifyNumLeft(actualId, lineDiagonal)
			valRight := verifyNumRight(actualId, lineDiagonal)
			if valLeft != 0 {
				nums = append(nums, valLeft)
				quantity++
			}
			if valRight != 0 {
				nums = append(nums, valRight)
				quantity++
			}
		}
	}
	if quantity == 0 {
		product = 0
	} else if quantity == 1 {
		product = nums[0]
	} else if quantity == 2 {
		product = nums[0] * nums[1]
	}

	return product, quantity
}
