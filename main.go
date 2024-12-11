package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jeroen-plug/advent-of-code-2024/day1"
	"github.com/jeroen-plug/advent-of-code-2024/day10"
	"github.com/jeroen-plug/advent-of-code-2024/day11"
	"github.com/jeroen-plug/advent-of-code-2024/day2"
	"github.com/jeroen-plug/advent-of-code-2024/day3"
	"github.com/jeroen-plug/advent-of-code-2024/day4"
	"github.com/jeroen-plug/advent-of-code-2024/day5"
	"github.com/jeroen-plug/advent-of-code-2024/day6"
	"github.com/jeroen-plug/advent-of-code-2024/day7"
	"github.com/jeroen-plug/advent-of-code-2024/day8"
	"github.com/jeroen-plug/advent-of-code-2024/day9"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [day]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		usage()
	}
	day, err := strconv.Atoi(os.Args[1])
	if err != nil || day <= 0 {
		usage()
	}

	switch day {
	case 1:
		day1.Day1()
	case 2:
		day2.Day2()
	case 3:
		day3.Day3()
	case 4:
		day4.Day4()
	case 5:
		day5.Day5()
	case 6:
		day6.Day6()
	case 7:
		day7.Day7()
	case 8:
		day8.Day8()
	case 9:
		day9.Day9()
	case 10:
		day10.Day10()
	case 11:
		day11.Day11()
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown day %d\n", day)
		usage()
	}
}
