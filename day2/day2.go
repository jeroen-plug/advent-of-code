package day2

import (
	"fmt"
	"math"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Day2() {
	lines := input.Lines(2)

	fmt.Printf("day 2a: %d\n", day2a(lines))
	fmt.Printf("day 2b: %d\n", day2b(lines))
}

func day2a(lines []string) int {
	var safeReports int

	for _, report := range lines {
		if isSafeLevels(parse(report)) {
			safeReports++
		}
	}

	return safeReports
}

func day2b(lines []string) int {
	var safeReports int

	for _, report := range lines {
		levels := parse(report)
		safe := isSafeLevels(levels)

		if !safe {
			for i := 0; i < len(levels); i++ {
				var levelsDampened []int
				levelsDampened = append(levelsDampened, levels[:i]...)
				levelsDampened = append(levelsDampened, levels[i+1:]...)

				if isSafeLevels(levelsDampened) {
					safe = true
					break
				}
			}
		}

		if safe {
			safeReports++
		}
	}

	return safeReports
}

func isSafeLevels(levels []int) bool {
	incrementing := isIncrementing(levels)

	for i := 1; i < len(levels); i++ {
		if !isSafeLevel(incrementing, levels[i-1], levels[i]) {
			return false
		}
	}

	return true
}

func isSafeLevel(incrementing bool, first int, second int) bool {
	if (incrementing && second <= first) ||
		(!incrementing && second >= first) ||
		math.Abs(float64(first-second)) > 3 {
		return false
	}
	return true
}

func isIncrementing(levels []int) bool {
	inc := 0

	for i := 1; i < len(levels); i++ {
		if levels[i] > levels[i-1] {
			inc++
		} else {
			inc--
		}
	}

	return inc > 0
}

func parse(report string) []int {
	var levels []int

	for _, l := range strings.Fields(report) {
		levels = append(levels, input.ParseInt(l))
	}

	return levels
}
