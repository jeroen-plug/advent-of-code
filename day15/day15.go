package day15

import (
	"fmt"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/grid"
	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Warehouse [][]rune

const (
	Wall  rune = '#'
	Box   rune = 'O'
	Robot rune = '@'
	Floor rune = '.'

	BigBoxL rune = '['
	BigBoxR rune = ']'
)

func Day15() {
	lines := input.Lines(15)

	fmt.Printf("day 15a: %d\n", day15a(lines))
	fmt.Printf("day 15b: %d\n", day15b(lines))
}

func day15a(lines []string) int {
	warehouse, moves := parse(lines)
	warehouse, robot := separateRobot(warehouse)
	warehouse.moveRobot(robot, moves)

	sum := 0
	for row, r := range warehouse {
		for col, o := range r {
			if o == Box {
				sum += 100*row + col
			}
		}
	}

	return sum
}

func day15b(lines []string) int {
	return 0
}

func (w *Warehouse) moveRobot(robot grid.Position, moves []grid.Direction) {
	for _, dir := range moves {
		canMove, willPush, pushEnd := w.checkMove(robot, dir)
		if !canMove {
			continue
		}
		newPosition := robot.Move(dir)
		if willPush {
			w.Set(pushEnd, Box)
			w.Set(newPosition, Floor)
		}
		robot = newPosition
	}
}

func (w Warehouse) checkMove(robot grid.Position, dir grid.Direction) (bool, bool, grid.Position) {
	canMove := false
	willPush := false
	pushEnd := grid.Position{}

	newPosition := robot.Move(dir)
	tile := w.At(newPosition)

	if tile == Floor {
		canMove = true
	} else if tile == Box || tile == BigBoxL || tile == BigBoxR {
		willPush = true
		canMove, pushEnd = w.canPush(newPosition, dir)
	}

	return canMove, willPush, pushEnd
}

func (w Warehouse) canPush(box grid.Position, dir grid.Direction) (bool, grid.Position) {
	canPush := false
	pushEnd := grid.Position{}

	newPosition := box.Move(dir)
	switch w.At(newPosition) {
	case Floor:
		canPush = true
		pushEnd = newPosition
	case Box:
		canPush, pushEnd = w.canPush(newPosition, dir)
	}

	return canPush, pushEnd
}

func (w Warehouse) String() string {
	var s strings.Builder
	for _, l := range w {
		s.WriteString(string(l))
		s.WriteRune('\n')
	}
	return s.String()
}

func (w Warehouse) At(p grid.Position) rune {
	return w[p.Row][p.Col]
}

func (w *Warehouse) Set(p grid.Position, v rune) {
	(*w)[p.Row][p.Col] = v
}

func separateRobot(warehouse Warehouse) (Warehouse, grid.Position) {
	for row, r := range warehouse {
		for col, c := range r {
			if c == Robot {
				warehouse[row][col] = Floor
				return warehouse, grid.Position{Row: row, Col: col}
			}
		}
	}
	return warehouse, grid.Position{}
}

func parse(lines []string) (Warehouse, []grid.Direction) {
	var warehouse Warehouse
	var moves []grid.Direction

	afterBreak := false
	for _, l := range lines {
		if l == "" {
			afterBreak = true
		} else if afterBreak {
			for _, m := range l {
				switch m {
				case '<':
					moves = append(moves, grid.Left)
				case '^':
					moves = append(moves, grid.Up)
				case '>':
					moves = append(moves, grid.Right)
				case 'v':
					moves = append(moves, grid.Down)
				}
			}
		} else {
			warehouse = append(warehouse, []rune(l))
		}
	}

	return warehouse, moves
}
