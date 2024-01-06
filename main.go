package main

import (
	"fmt"
	"os"
)

func main() {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		fmt.Println("Recovered from panic:", r)
	// 	}
	// }()
	if len(os.Args) < 3 {
		panic("Insufficient arguments!")
	}

	switch os.Args[1] {
	case "1":
		if os.Args[2] == "1" {
			day1_1()
		} else if os.Args[2] == "2" {
			day1_2()
		}
	case "2":
		if os.Args[2] == "1" {
			day2_1()
		} else if os.Args[2] == "2" {
			day2_2()
		}
	case "3":
		if os.Args[2] == "1" {
			day3_1()
		} else if os.Args[2] == "2" {
			day3_2()
		}
	case "4":
		if os.Args[2] == "1" {
			day4_1()
		} else if os.Args[2] == "2" {
			day4_2()
		}
	case "5":
		if os.Args[2] == "1" {
			day5_1()
		} else if os.Args[2] == "2" {
			day5_2()
		}
	case "6":
		if os.Args[2] == "1" {
			day6_1()
		} else if os.Args[2] == "2" {
			day6_2()
		}
	case "7":
		if os.Args[2] == "1" {
			day7_1()
		} else if os.Args[2] == "2" {
			day7_2()
		}
	default:
		fmt.Println("Day not found")
	}
}
