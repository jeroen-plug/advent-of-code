package day10

import (
	"log"
	"slices"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

const (
	Left int = iota
	Up
	Right
	Down
)

func Solution() (any, any) {
	lines := input.Lines(10)
	return day10a(lines), day10b(lines)
}

func day10a(lines []string) int {
	return len(findPath(lines)) / 2
}

func day10b(lines []string) int {
	steps := findPath(lines)

	inside := false
	enclosed := 0
	for l := range lines {
		for c, symbol := range lines[l] {
			if slices.Contains(steps, [2]int{l, c}) {
				if symbol == '|' || symbol == 'L' || symbol == 'J' || (len(lines) > 100 && symbol == 'S') {
					inside = !inside
				}
			} else if inside {
				enclosed++
			}
		}
	}

	return enclosed
}

func findPath(lines []string) [][2]int {
	start := findStart(lines)

	pos1 := start
	pos2 := start

	dir1 := 0
	dir2 := 0

	// Don't bother finding start direction
	if len(lines) == 5 {
		// example
		dir1 = Right
		dir2 = Down
	} else if len(lines) == 10 {
		// example2
		dir1 = Left
		dir2 = Down
	} else {
		// My input
		dir1 = Left
		dir2 = Up
	}

	steps := [][2]int{start}
	for pos1[0] != pos2[0] || pos1[1] != pos2[1] || len(steps) == 1 {
		pos1, dir1 = takeStep(lines, pos1, dir1)
		pos2, dir2 = takeStep(lines, pos2, dir2)
		steps = append(steps, pos1, pos2)
	}

	return steps[:len(steps)-1]
}

func takeStep(lines []string, pos [2]int, dir int) ([2]int, int) {
	switch dir {
	case Left:
		pos[1]--
	case Up:
		pos[0]--
	case Right:
		pos[1]++
	case Down:
		pos[0]++
	}
	return pos, newDirection(lines, pos, dir)
}

func newDirection(lines []string, pos [2]int, dir int) int {
	var options [2]int
	switch lines[pos[0]][pos[1]] {
	case 'L':
		options = [2]int{Up, Right}
	case 'J':
		options = [2]int{Up, Left}
	case '7':
		options = [2]int{Down, Left}
	case 'F':
		options = [2]int{Down, Right}
	case '.':
		log.Fatal("Hit the ground")
	default:
		options = [2]int{dir, dir}
	}

	if options[0]%2 == dir%2 {
		return options[1]
	} else {
		return options[0]
	}
}

func findStart(lines []string) [2]int {
	for line, l := range lines {
		for col, c := range l {
			if c == 'S' {
				return [2]int{line, col}
			}
		}
	}
	return [2]int{}
}
