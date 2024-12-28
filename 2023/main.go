package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"

	"github.com/fatih/color"
	"github.com/jeroen-plug/advent-of-code/2023/day1"
	"github.com/jeroen-plug/advent-of-code/2023/day10"
	"github.com/jeroen-plug/advent-of-code/2023/day11"
	"github.com/jeroen-plug/advent-of-code/2023/day12"
	"github.com/jeroen-plug/advent-of-code/2023/day13"
	"github.com/jeroen-plug/advent-of-code/2023/day14"
	"github.com/jeroen-plug/advent-of-code/2023/day15"
	"github.com/jeroen-plug/advent-of-code/2023/day16"
	"github.com/jeroen-plug/advent-of-code/2023/day17"
	"github.com/jeroen-plug/advent-of-code/2023/day2"
	"github.com/jeroen-plug/advent-of-code/2023/day3"
	"github.com/jeroen-plug/advent-of-code/2023/day4"
	"github.com/jeroen-plug/advent-of-code/2023/day5"
	"github.com/jeroen-plug/advent-of-code/2023/day6"
	"github.com/jeroen-plug/advent-of-code/2023/day7"
	"github.com/jeroen-plug/advent-of-code/2023/day8"
	"github.com/jeroen-plug/advent-of-code/2023/day9"
)

type Day struct {
	Title    string
	Solution func() (any, any)
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s [day]\n", os.Args[0])
	os.Exit(1)
}

func main() {
	day := 0
	all := true

	if len(os.Args) >= 2 {
		var err error
		day, err = strconv.Atoi(os.Args[1])
		if err != nil || day <= 0 || day > 25 {
			usage()
		}
		all = false
	}

	days := map[int]Day{
		1:  {"Trebuchet?!", day1.Solution},
		2:  {"Cube Conundrum", day2.Solution},
		3:  {"Gear Ratios", day3.Solution},
		4:  {"Scratchcards", day4.Solution},
		5:  {"If You Give A Seed A Fertilizer", day5.Solution},
		6:  {"Wait For It", day6.Solution},
		7:  {"Camel Cards", day7.Solution},
		8:  {"Haunted Wasteland", day8.Solution},
		9:  {"Mirage Maintenance", day9.Solution},
		10: {"Pipe Maze", day10.Solution},
		11: {"Cosmic Expansion", day11.Solution},
		12: {"Hot Springs", day12.Solution},
		13: {"Point of Incidence", day13.Solution},
		14: {"Parabolic Reflector Dish", day14.Solution},
		15: {"Lens Library", day15.Solution},
		16: {"The Floor Will Be Lava", day16.Solution},
		17: {"Clumsy Crucible", day17.Solution},
	}

	if all {
		keys := make([]int, 0, len(days))
		for k := range maps.Keys(days) {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		for _, day := range keys {
			runDay(day, days[day])
		}
	} else {
		if solution, exists := days[day]; exists {
			runDay(day, solution)
		} else {
			fmt.Printf("Day %d not implemented.\n", day)
		}
	}
}

func runDay(n int, day Day) {
	c := color.New(color.FgYellow).Add(color.Bold)
	part1, part2 := day.Solution()
	c.Printf("2023 Day %d: %s\n", n, color.New(color.ResetBold).Sprint(day.Title))
	fmt.Printf("    Part 1: %v\n", part1)
	fmt.Printf("    Part 2: %v\n", part2)
}
