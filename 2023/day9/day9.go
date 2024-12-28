package day9

import (
	"strings"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

func Solution() (any, any) {
	lines := input.Lines(9)
	return day9a(lines), day9b(lines)
}

func day9a(lines []string) int {
	report := parse(lines)

	sum := 0
	for _, history := range report {
		sequences := recurseDifferences(history)
		next := 0
		for i := len(sequences) - 2; i >= 0; i-- {
			next = next + sequences[i][len(sequences[i])-1]
		}
		sum += next
	}

	return sum
}

func day9b(lines []string) int {
	report := parse(lines)

	sum := 0
	for _, history := range report {
		sequences := recurseDifferences(history)
		next := 0
		for i := len(sequences) - 2; i >= 0; i-- {
			next = sequences[i][0] - next
		}
		sum += next
	}

	return sum
}

func recurseDifferences(in []int) [][]int {
	var out [][]int
	out = append(out, in)
	for {
		diff, zero := differences(out[len(out)-1])
		out = append(out, diff)
		if zero {
			break
		}
	}
	return out
}

func differences(in []int) ([]int, bool) {
	var out []int
	zero := true
	for i, n := range in[1:] {
		diff := n - in[i]
		out = append(out, diff)
		if diff != 0 {
			zero = false
		}
	}
	return out, zero
}

func parse(lines []string) [][]int {
	var report [][]int
	for _, l := range lines {
		var value []int
		for _, f := range strings.Fields(l) {
			value = append(value, input.ParseInt(f))
		}
		report = append(report, value)
	}
	return report
}
