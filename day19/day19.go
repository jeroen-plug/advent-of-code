package day19

import (
	"fmt"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Day19() {
	lines := input.Lines(19)

	fmt.Printf("day 19a: %d\n", day19a(lines))
	fmt.Printf("day 19b: %d\n", day19b(lines))
}

func day19a(lines []string) int {
	available, toDisplay := parse(lines)
	sum := 0
	for _, d := range toDisplay {
		if dfs(available, d) > 0 {
			sum++
		}
	}
	return sum
}

func day19b(lines []string) int {
	available, toDisplay := parse(lines)
	sum := 0
	for _, d := range toDisplay {
		sum += dfs(available, d)
	}
	return sum
}

func dfs(available []string, toDisplay string) int {
	memo := make(map[string]int)
	var dfsRecurse func(current string) int

	dfsRecurse = func(current string) int {
		if current == "" {
			return 1
		} else if m, ok := memo[current]; ok {
			return m
		}

		sum := 0
		for _, pattern := range available {
			if next, ok := strings.CutPrefix(current, pattern); ok {
				sum += dfsRecurse(next)
			}
		}

		memo[current] = sum
		return sum
	}

	return dfsRecurse(toDisplay)
}

func parse(lines []string) ([]string, []string) {
	available := strings.Split(lines[0], ", ")
	toDisplay := lines[2:]

	return available, toDisplay
}
