package day4

import (
	"fmt"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Coordinate struct {
	row int
	col int
}

type Direction int

const (
	Left Direction = iota
	UpLeft
	Up
	UpRight
	Right
	DownRight
	Down
	DownLeft
)

func Day4() {
	lines := input.Lines(4)

	fmt.Printf("day 4a: %d\n", day4a(lines))
	fmt.Printf("day 4b: %d\n", day4b(lines))
}

func day4a(lines []string) int {
	var candidates []Coordinate
	for row, l := range lines {
		for col, c := range l {
			if c == 'X' {
				candidates = append(candidates, Coordinate{row, col})
			}
		}
	}

	var count int
	for _, c := range candidates {
		for d := Left; d <= DownLeft; d++ {
			if checkCandidate(lines, c, d) {
				count++
			}
		}
	}
	return count
}

func day4b(lines []string) int {
	var count int
	for row := 1; row < len(lines)-1; row++ {
		for col := 1; col < len(lines[0])-1; col++ {
			if lines[row][col] != 'A' {
				continue
			}

			m := 0
			s := 0
			for _, offset := range [][2]int{{1, 1}, {1, -1}, {-1, 1}, {-1, -1}} {
				if lines[row+offset[0]][col+offset[1]] == 'M' {
					m++
				}
				if lines[row+offset[0]][col+offset[1]] == 'S' {
					s++
				}
			}

			if m == 2 && s == 2 && lines[row+1][col+1] != lines[row-1][col-1] {
				count++
			}
		}
	}
	return count
}

func checkCandidate(lines []string, c Coordinate, dir Direction) bool {
	toCheck := "MAS" // the X is the candidate

	if c.col < len(toCheck) && (dir == Left || dir == DownLeft || dir == UpLeft) {
		return false
	}
	if c.col >= len(lines[0])-len(toCheck) && (dir == Right || dir == DownRight || dir == UpRight) {
		return false
	}
	if c.row < len(toCheck) && (dir == Up || dir == UpLeft || dir == UpRight) {
		return false
	}
	if c.row >= len(lines)-len(toCheck) && (dir == Down || dir == DownLeft || dir == DownRight) {
		return false
	}

	var (
		colStep int
		rowStep int
	)
	switch dir {
	case Left:
		colStep = -1
	case UpLeft:
		colStep = -1
		rowStep = -1
	case Up:
		rowStep = -1
	case UpRight:
		colStep = 1
		rowStep = -1
	case Right:
		colStep = 1
	case DownRight:
		colStep = 1
		rowStep = 1
	case Down:
		rowStep = 1
	case DownLeft:
		colStep = -1
		rowStep = 1
	}

	col := c.col
	row := c.row
	for _, check := range toCheck {
		col += colStep
		row += rowStep

		if lines[row][col] != byte(check) {
			return false
		}
	}

	return true
}
