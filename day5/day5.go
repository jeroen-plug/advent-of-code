package day5

import (
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Printer struct {
	Rules   [][2]int
	Updates [][]int
}

func Solution() (any, any) {
	lines := input.Lines(5)
	return day5a(lines), day5b(lines)
}

func day5a(lines []string) int {
	p := parse(lines)

	sum := 0
	for _, update := range p.Updates {
		if validUpdate(p.Rules, update) {
			sum += update[len(update)/2]
		}
	}

	return sum
}

func day5b(lines []string) int {
	p := parse(lines)

	sum := 0
	for _, update := range p.Updates {
		if !validUpdate(p.Rules, update) {
			slices.SortFunc(update, func(a, b int) int {
				if a == b {
					return 0
				}
				for _, r := range p.Rules {
					if r[0] == a && r[1] == b {
						return 1
					}
				}
				return -1
			})
			sum += update[len(update)/2]
		}
	}

	return sum
}

func validUpdate(rules [][2]int, update []int) bool {
	var seen []int
	for _, u := range update {
		for _, r := range rules {
			if r[0] == u && slices.Contains(seen, r[1]) {
				return false
			}
		}
		seen = append(seen, u)
	}
	return true
}

func parse(lines []string) Printer {
	var p Printer
	beforeBreak := true

	for _, l := range lines {
		if len(l) < 1 {
			beforeBreak = false
			continue
		}

		if beforeBreak {
			fields := strings.Split(l, "|")
			p.Rules = append(p.Rules, [2]int{input.ParseInt(fields[0]), input.ParseInt(fields[1])})
		} else {
			var updates []int
			for _, n := range strings.Split(l, ",") {
				updates = append(updates, input.ParseInt(n))
			}
			p.Updates = append(p.Updates, updates)
		}
	}

	// slices.SortFunc(p.Rules, func(a, b [2]int) int { return cmp.Compare(a[0], b[0]) })

	return p
}
