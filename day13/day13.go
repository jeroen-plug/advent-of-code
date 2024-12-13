package day13

import (
	"fmt"
	"math"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type ClawMachine struct {
	ButtonA [2]int
	ButtonB [2]int
	Prize   [2]int
}

func Day13() {
	lines := input.Lines(13)

	fmt.Printf("day 13a: %d\n", day13a(lines))
	fmt.Printf("day 13b: %d\n", day13b(lines))
}

func day13a(lines []string) int {
	cost := 0
	for _, m := range parse(lines) {
		a, b, solved := solve(
			float64(m.ButtonA[0]), float64(m.ButtonB[0]), float64(m.Prize[0]),
			float64(m.ButtonA[1]), float64(m.ButtonB[1]), float64(m.Prize[1]),
		)
		if solved && isInteger(a) && isInteger(b) {
			cost += int(3*a + b)
		}
	}
	return cost
}

func day13b(lines []string) int {
	cost := 0
	for _, m := range parse(lines) {
		a, b, solved := solve(
			float64(m.ButtonA[0]), float64(m.ButtonB[0]), float64(m.Prize[0])+10000000000000,
			float64(m.ButtonA[1]), float64(m.ButtonB[1]), float64(m.Prize[1])+10000000000000,
		)
		if solved && isInteger(a) && isInteger(b) {
			cost += int(3*a + b)
		}
	}
	return cost
}

// Cramer's rule
func solve(a1, b1, c1, a2, b2, c2 float64) (float64, float64, bool) {
	determinant := a1*b2 - b1*a2
	if determinant == 0 {
		return 0, 0, false
	}

	x := (c1*b2 - b1*c2) / determinant
	y := (a1*c2 - c1*a2) / determinant
	return x, y, true
}

func parse(lines []string) []ClawMachine {
	var machines []ClawMachine
	var m ClawMachine
	for _, l := range lines {
		if l == "" {
			continue
		}
		fields := strings.Split(l, ":")
		values := strings.Split(fields[1], ",")
		switch fields[0] {
		case "Button A":
			m.ButtonA = [2]int{parseValue(values[0]), parseValue(values[1])}
		case "Button B":
			m.ButtonB = [2]int{parseValue(values[0]), parseValue(values[1])}
		case "Prize":
			m.Prize = [2]int{parseValue(values[0]), parseValue(values[1])}

			machines = append(machines, m)
			m = ClawMachine{}
		}
	}
	return machines
}

func parseValue(s string) int {
	return input.ParseInt(strings.TrimLeft(s, " XY+="))
}

func isInteger(n float64) bool {
	return math.Trunc(n) == n
}
