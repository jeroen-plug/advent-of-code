package main

import (
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"

	"github.com/fatih/color"
	"github.com/jeroen-plug/advent-of-code/2024/day1"
	"github.com/jeroen-plug/advent-of-code/2024/day10"
	"github.com/jeroen-plug/advent-of-code/2024/day11"
	"github.com/jeroen-plug/advent-of-code/2024/day12"
	"github.com/jeroen-plug/advent-of-code/2024/day13"
	"github.com/jeroen-plug/advent-of-code/2024/day14"
	"github.com/jeroen-plug/advent-of-code/2024/day15"
	"github.com/jeroen-plug/advent-of-code/2024/day16"
	"github.com/jeroen-plug/advent-of-code/2024/day17"
	"github.com/jeroen-plug/advent-of-code/2024/day18"
	"github.com/jeroen-plug/advent-of-code/2024/day19"
	"github.com/jeroen-plug/advent-of-code/2024/day2"
	"github.com/jeroen-plug/advent-of-code/2024/day20"
	"github.com/jeroen-plug/advent-of-code/2024/day21"
	"github.com/jeroen-plug/advent-of-code/2024/day22"
	"github.com/jeroen-plug/advent-of-code/2024/day23"
	"github.com/jeroen-plug/advent-of-code/2024/day24"
	"github.com/jeroen-plug/advent-of-code/2024/day25"
	"github.com/jeroen-plug/advent-of-code/2024/day3"
	"github.com/jeroen-plug/advent-of-code/2024/day4"
	"github.com/jeroen-plug/advent-of-code/2024/day5"
	"github.com/jeroen-plug/advent-of-code/2024/day6"
	"github.com/jeroen-plug/advent-of-code/2024/day7"
	"github.com/jeroen-plug/advent-of-code/2024/day8"
	"github.com/jeroen-plug/advent-of-code/2024/day9"
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
		1:  {"Historian Hysteria", day1.Solution},
		2:  {"Red-Nosed Reports", day2.Solution},
		3:  {"Mull It Over", day3.Solution},
		4:  {"Ceres Search", day4.Solution},
		5:  {"Print Queue", day5.Solution},
		6:  {"Guard Gallivant", day6.Solution},
		7:  {"Bridge Repair", day7.Solution},
		8:  {"Resonant Collinearity", day8.Solution},
		9:  {"Disk Fragmenter", day9.Solution},
		10: {"Hoof It", day10.Solution},
		11: {"Plutonian Pebbles", day11.Solution},
		12: {"Garden Groups", day12.Solution},
		13: {"Claw Contraption", day13.Solution},
		14: {"Restroom Redoubt", day14.Solution},
		15: {"Warehouse Woes", day15.Solution},
		16: {"Reindeer Maze", day16.Solution},
		17: {"Chronospatial Computer", day17.Solution},
		18: {"RAM Run", day18.Solution},
		19: {"Linen Layout", day19.Solution},
		20: {"Race Condition", day20.Solution},
		21: {"Keypad Conundrum", day21.Solution},
		22: {"Monkey Market", day22.Solution},
		23: {"LAN Party", day23.Solution},
		24: {"Crossed Wires", day24.Solution},
		25: {"Code Chronicle", day25.Solution},
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
	c.Printf("2024 Day %d: %s\n", n, color.New(color.ResetBold).Sprint(day.Title))
	fmt.Printf("    Part 1: %v\n", part1)
	fmt.Printf("    Part 2: %v\n", part2)
}
