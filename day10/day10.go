package day10

import (
	"slices"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Solution() (any, any) {
	lines := input.Lines(10)
	return day10a(lines), day10b(lines)
}

func day10a(lines []string) int {
	trailheads := findTrailheads(lines)

	sum := 0
	for _, t := range trailheads {
		sum += countTrails(lines, t, true)
	}

	return sum
}

func day10b(lines []string) int {
	trailheads := findTrailheads(lines)

	sum := 0
	for _, t := range trailheads {
		sum += countTrails(lines, t, false)
	}

	return sum
}

func findTrailheads(lines []string) [][2]int {
	var trailheads [][2]int
	for r, row := range lines {
		for c, col := range row {
			if col == '0' {
				trailheads = append(trailheads, [2]int{r, c})
			}
		}
	}
	return trailheads
}

func countTrails(lines []string, trailhead [2]int, findUnique bool) int {
	var visited [][2]int
	peaks := dfs(lines, &visited, trailhead, '0', findUnique)
	return len(peaks)
}

func dfs(lines []string, visited *[][2]int, current [2]int, height byte, findUnique bool) [][2]int {
	if current[0] < 0 || current[0] >= len(lines) ||
		current[1] < 0 || current[1] >= len(lines[0]) ||
		lines[current[0]][current[1]] != height ||
		(findUnique && slices.Contains(*visited, current)) {

		return [][2]int{}
	}

	*visited = append(*visited, current)

	if lines[current[0]][current[1]] == '9' {
		return [][2]int{current}
	}

	var peaks [][2]int
	peaks = append(peaks, dfs(lines, visited, [2]int{current[0] + 1, current[1]}, height+1, findUnique)...)
	peaks = append(peaks, dfs(lines, visited, [2]int{current[0] - 1, current[1]}, height+1, findUnique)...)
	peaks = append(peaks, dfs(lines, visited, [2]int{current[0], current[1] + 1}, height+1, findUnique)...)
	peaks = append(peaks, dfs(lines, visited, [2]int{current[0], current[1] - 1}, height+1, findUnique)...)
	return peaks
}
