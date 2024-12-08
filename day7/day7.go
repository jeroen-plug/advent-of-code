package day7

import (
	"fmt"
	"math"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Equation struct {
	result int
	values []int
}

func Day7() {
	lines := input.Lines(7)

	fmt.Printf("day 7a: %d\n", day7a(lines))
	fmt.Printf("day 7b: %d\n", day7b(lines))
}

func day7a(lines []string) int {
	equations := parse(lines)

	sum := 0
	for _, e := range equations {
		if doDfs(e, false) {
			sum += e.result
		}
	}
	return sum
}

func day7b(lines []string) int {
	equations := parse(lines)

	sum := 0
	for _, e := range equations {
		if doDfs(e, true) {
			sum += e.result
		}
	}
	return sum
}

func doDfs(e Equation, allowConcat bool) bool {
	return dfs(e.values[1:], e.values[0], e.result, allowConcat)
}

func dfs(values []int, result, goal int, allowConcat bool) bool {
	if len(values) < 1 {
		return result == goal
	}

	add := dfs(values[1:], result+values[0], goal, allowConcat)
	multiply := dfs(values[1:], result*values[0], goal, allowConcat)
	concat := allowConcat && dfs(values[1:], concat(result, values[0]), goal, allowConcat)

	return add || multiply || concat
}

func concat(left, right int) int {
	lenRight := int(math.Log10(float64(right))) + 1
	return left*int(math.Pow10(lenRight)) + right
}

func parse(lines []string) []Equation {
	var equations []Equation
	for _, l := range lines {
		fields := strings.Split(l, ":")
		var values []int
		for _, v := range strings.Fields(fields[1]) {
			values = append(values, input.ParseInt(v))
		}
		equations = append(equations, Equation{input.ParseInt(fields[0]), values})
	}
	return equations
}
