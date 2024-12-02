package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/jeroen-plug/advent-of-code-2024/day01"
	"github.com/jeroen-plug/advent-of-code-2024/day02"
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
		day01.Day1()
	case 2:
		day02.Day2()
	default:
		fmt.Fprintf(os.Stderr, "Error: Unknown day %d\n", day)
		usage()
	}
}
