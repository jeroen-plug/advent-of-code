package day15

import (
	"container/list"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2024/grid"
	"github.com/jeroen-plug/advent-of-code/2024/input"
)

type Warehouse [][]rune

const (
	Wall  rune = '#'
	Box   rune = 'O'
	Robot rune = '@'
	Floor rune = '.'

	BigBox rune = '['
)

func Solution() (any, any) {
	lines := input.Lines(15)
	return day15a(lines), day15b(lines)
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
	warehouse, moves := parse(lines)
	warehouse = stretch(warehouse)
	warehouse, robot := separateRobot(warehouse)
	warehouse.moveRobot(robot, moves)

	sum := 0
	for row, r := range warehouse {
		for col, o := range r {
			if o == BigBox {
				sum += 100*row + col
			}
		}
	}

	return sum
}

func (w *Warehouse) moveRobot(robot grid.Position, moves []grid.Direction) grid.Position {
	for _, dir := range moves {
		robot = w.tryMove(robot, dir)
	}
	return robot
}

func (w *Warehouse) tryMove(robot grid.Position, dir grid.Direction) grid.Position {
	newPosition := robot.Move(dir)
	tile := w.At(newPosition)

	if tile == Box {
		if canPush, pushEnd := w.canPushSmall(newPosition, dir); canPush {
			w.Set(pushEnd, Box)
			w.Set(newPosition, Floor)
			return newPosition
		}
	} else if tile == BigBox {
		if canPush := w.tryPushBig(newPosition, dir); canPush {
			return newPosition
		}
	} else if tile == Floor && w.At(newPosition.Move(grid.Left)) == BigBox {
		if canPush := w.tryPushBig(newPosition.Move(grid.Left), dir); canPush {
			return newPosition
		}
	} else if tile == Floor {
		return newPosition
	}

	return robot
}

func (w Warehouse) tryPushBig(box grid.Position, dir grid.Direction) bool {
	// queue
	toCheck := list.New()
	toCheck.PushBack(box)

	// stack
	toMove := list.New()

	for toCheck.Len() > 0 {
		i := toCheck.Front()
		check := i.Value.(grid.Position)

		// Check for walls
		if dir == grid.Left && w.At(check.Move(dir)) == Wall ||
			dir == grid.Right && w.At(check.Move(dir).Move(dir)) == Wall ||
			(dir == grid.Up || dir == grid.Down) && (w.At(check.Move(dir)) == Wall || w.At(check.Move(dir).Move(grid.Right)) == Wall) {
			return false
		}

		// Check for adjacent boxes
		if dir == grid.Left && w.At(check.Move(dir).Move(dir)) == BigBox {
			toCheck.PushBack(check.Move(dir).Move(dir))
		} else if dir == grid.Right && w.At(check.Move(dir).Move(dir)) == BigBox {
			toCheck.PushBack(check.Move(dir).Move(dir))
		} else if dir == grid.Up || dir == grid.Down {
			if w.At(check.Move(dir)) == BigBox {
				toCheck.PushBack(check.Move(dir))
			} else {
				if w.At(check.Move(dir).Move(grid.Left)) == BigBox {
					toCheck.PushBack(check.Move(dir).Move(grid.Left))
				}
				if w.At(check.Move(dir).Move(grid.Right)) == BigBox {
					toCheck.PushBack(check.Move(dir).Move(grid.Right))
				}
			}
		}

		toMove.PushBack(check)
		toCheck.Remove(i)
	}

	for toMove.Len() > 0 {
		i := toMove.Back()
		move := i.Value.(grid.Position)
		w.Set(move, Floor)
		w.Set(move.Move(dir), BigBox)
		toMove.Remove(i)
	}

	return true
}

func (w Warehouse) canPushSmall(box grid.Position, dir grid.Direction) (bool, grid.Position) {
	canPush := false
	pushEnd := grid.Position{}

	newPosition := box.Move(dir)
	switch w.At(newPosition) {
	case Floor:
		canPush = true
		pushEnd = newPosition
	case Box:
		canPush, pushEnd = w.canPushSmall(newPosition, dir)
	}

	return canPush, pushEnd
}

func (w Warehouse) String() string {
	var s strings.Builder
	for _, l := range w {
		s.WriteString(strings.ReplaceAll(string(l), "[.", "[]"))
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

func stretch(warehouse Warehouse) Warehouse {
	var stretched Warehouse
	for _, row := range warehouse {
		var newRow []rune
		for _, c := range row {
			if c == Box {
				newRow = append(newRow, BigBox)
			} else {
				newRow = append(newRow, c)
			}
			if c == Wall {
				newRow = append(newRow, Wall)
			} else {
				newRow = append(newRow, Floor)
			}
		}
		stretched = append(stretched, newRow)
	}
	return stretched
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
