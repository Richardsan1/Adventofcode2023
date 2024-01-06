package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day3_1() {
	dat, err := os.ReadFile("./inputs/day3.txt")
	if err != nil {
		panic(err)
	}
	text := strings.Split(string(dat), "\n")
	var sum int = 0
	// cada linha tem 140 caracteres

	for lineID, line := range text { // para cada linha
		for charID, char := range line { // para cada caractere da linha
			if char != 46 && !verifyNum(char) { // verifica se é um simbolo
				if charID > 0 { // verifica se não é o primeiro caractere
					sum += verifyNumLeft(charID, []byte(line))
				}
				if charID < 139 { // verifica se não é o último caractere
					sum += verifyNumRight(charID, []byte(line))
				}
				if lineID > 0 {
					sum += verifyNumDiagonal([]byte(text[lineID-1]), charID)
				}
				if lineID < len(text)-1 {
					sum += verifyNumDiagonal([]byte(text[lineID+1]), charID)
				}
			}
		}
	}
	fmt.Println(sum)
}

func verifyNum(test rune) bool {
	test_ := byte(test)

	if test_ > 47 && test_ < 58 {
		return true
	} else {
		return false
	}
}

func verifyNumRight(idActual int, line []byte) int {
	var nums []int
	var val int = 0
	for i := idActual + 1; i < 140; i++ {
		if verifyNum(rune(line[i])) {
			nums = append([]int{int(line[i]) - 48}, nums...)
		} else {
			break
		}
	}
	for index, num := range nums {
		val += num * int(math.Pow(10, float64(index)))
	}
	return val
}

func verifyNumLeft(idActual int, line []byte) int {
	var ActualNum []int
	var val int = 0

	for i := idActual - 1; i >= 0; i-- { // percorre a linha da direita para a esquerda
		if verifyNum(rune(line[i])) { // verifica se o caractere anterior é um número
			ActualNum = append(ActualNum, int(line[i])-48)
		} else {
			break
		}
	}

	for index, num := range ActualNum {
		val += num * int(math.Pow(10, float64(index)))
	}

	return val
}

func verifyNumDiagonal(lineDiagonal []byte, actualId int) int {
	var sum int
	if actualId >= 0 && actualId <= 139 {
		if verifyNum(rune(lineDiagonal[actualId])) {
			val := verifyNumLeft(actualId+1, lineDiagonal)
			if len(strconv.Itoa(val)) > 1 {
				val2 := verifyNumLeft(actualId+2, lineDiagonal)
				if len(strconv.Itoa(val2)) > 1 {
					sum += val2
				} else {
					sum += val
				}
			} else {
				sum += verifyNumRight(actualId-1, lineDiagonal)
			}

		} else {
			sum += verifyNumLeft(actualId, lineDiagonal)
			sum += verifyNumRight(actualId, lineDiagonal)
		}
	}

	return sum
}
