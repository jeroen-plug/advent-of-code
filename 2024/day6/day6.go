package day6

import (
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2024/input"
)

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

type PosDir struct {
	Row int
	Col int
	Dir Direction
}

func Solution() (any, any) {
	lines := input.Lines(6)
	return day6a(lines), day6b(lines)
}

func day6a(lines []string) int {
	obstacles, guard := findStuff(lines)
	result := slices.Clone(lines)

	for guardInBounds(lines, guard) {
		result[guard.Row] = result[guard.Row][:guard.Col] + "X" + result[guard.Row][guard.Col+1:]
		guard = moveGuard(guard, obstacles)
	}

	visited := 0
	for _, l := range result {
		visited += strings.Count(l, "X")
		// fmt.Println(withColor(l)
	}

	return visited
}

func day6b(lines []string) int {
	obstacles, guard := findStuff(lines)
	start := guard

	visited := []PosDir{start}
	for guardInBounds(lines, guard) {
		guard = moveGuard(guard, obstacles)
		visited = append(visited, guard)
	}

	permutations := 0
	tried := [][2]int{{start.Row, start.Col}}
	for i, newObstacle := range visited {
		newPos := [2]int{newObstacle.Row, newObstacle.Col}
		if slices.Contains(tried, newPos) {
			continue
		}
		if checkLoop(lines, append(obstacles, newPos), visited[i-1]) {
			permutations++
		}
		tried = append(tried, newPos)
	}

	return permutations
}

// func withColor(line string) string {
// 	var p strings.Builder
// 	for _, c := range line {
// 		switch c {
// 		case '.':
// 			p.WriteString("\033[0;40;90m")
// 		case '#':
// 			p.WriteString("\033[0;101;31m")
// 		case 'X':
// 			p.WriteString("\033[0;46;96m")
// 		}
// 		p.WriteRune(c)
// 		p.WriteString("\033[0m")
// 	}
// 	return p.String()
// }

func checkLoop(lines []string, obstacles [][2]int, start PosDir) bool {
	lookup := make([][][]Direction, len(lines))
	for i := range lookup {
		lookup[i] = make([][]Direction, len(lines[0]))
	}

	guard := moveGuard(start, obstacles)
	for guardInBounds(lines, guard) {
		if slices.Contains(lookup[guard.Row][guard.Col], guard.Dir) {
			return true
		}
		lookup[guard.Row][guard.Col] = append(lookup[guard.Row][guard.Col], guard.Dir)
		guard = moveGuard(guard, obstacles)
	}

	return false
}

func guardInBounds(lines []string, pos PosDir) bool {
	return pos.Row >= 0 && pos.Row < len(lines) && pos.Col >= 0 && pos.Col < len(lines[0])
}

func moveGuard(guard PosDir, obstacles [][2]int) PosDir {
	dRow := 0
	dCol := 0

	if guard.Dir%2 == 0 {
		dCol = int(guard.Dir) - 1
	} else {
		dRow = int(guard.Dir) - 2
	}

	if slices.Contains(obstacles, [2]int{guard.Row + dRow, guard.Col + dCol}) {
		guard.Dir = (guard.Dir + 1) % 4
		return moveGuard(guard, obstacles)
	} else {
		guard.Row += dRow
		guard.Col += dCol
		return guard
	}
}

func findStuff(lines []string) ([][2]int, PosDir) {
	var obstacles [][2]int
	var guard PosDir
	for r, row := range lines {
		for c, col := range row {
			if col == '#' {
				obstacles = append(obstacles, [2]int{r, c})
			} else if col == '<' {
				guard = PosDir{r, c, Left}
			} else if col == '^' {
				guard = PosDir{r, c, Up}
			} else if col == '>' {
				guard = PosDir{r, c, Right}
			} else if col == 'v' {
				guard = PosDir{r, c, Down}
			}
		}
	}
	return obstacles, guard
}
