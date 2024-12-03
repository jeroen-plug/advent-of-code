package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jeroen-plug/advent-of-code-2024/day1"
	"github.com/jeroen-plug/advent-of-code-2024/day2"
	"github.com/jeroen-plug/advent-of-code-2024/day3"
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
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown day %d\n", day)
		usage()
	}
}
