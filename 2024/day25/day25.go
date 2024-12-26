package day25

import (
	"github.com/jeroen-plug/advent-of-code/2024/input"
)

type Tumbler [5]int

func (a Tumbler) Fits(b Tumbler) bool {
	for i := range a {
		if a[i]+b[i] > 5 {
			return false
		}
	}
	return true
}

func Solution() (any, any) {
	lines := input.Lines(25)
	return day25a(lines), nil
}

func day25a(lines []string) int {
	locks, keys := parse(lines)

	sum := 0
	for _, lock := range locks {
		for _, key := range keys {
			if lock.Fits(key) {
				sum++
			}
		}
	}
	return sum
}

func parse(lines []string) ([]Tumbler, []Tumbler) {
	var locks, keys []Tumbler

	var current []string
	for _, l := range lines {
		if l != "" {
			current = append(current, l)
		}
		if len(current) == 7 {
			if tumbler, isKey := parseTumbler(current); isKey {
				keys = append(keys, tumbler)
			} else {
				locks = append(locks, tumbler)
			}
			current = []string{}
		}
	}

	return locks, keys
}

func parseTumbler(lines []string) (Tumbler, bool) {
	isKey := true
	heights := lines[:6]
	if lines[0] == "#####" {
		isKey = false
		heights = lines[1:]
	}

	var tumbler Tumbler
	for _, l := range heights {
		for i := range l {
			if l[i] == '#' {
				tumbler[i]++
			}
		}
	}

	return tumbler, isKey
}
