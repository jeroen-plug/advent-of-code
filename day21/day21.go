package day21

import (
	"math"

	"github.com/jeroen-plug/advent-of-code-2024/grid"
	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Solution() (any, any) {
	lines := input.Lines(21)
	return day21a(lines), day21b(lines)
}

func day21a(lines []string) int {
	db := newDpadBots(2)
	sum := 0
	for _, code := range lines {
		sum += db.CountCodeSteps(code) * input.ParseInt(code[:3])
	}
	return sum
}

func day21b(lines []string) int {
	db := newDpadBots(25)
	sum := 0
	for _, code := range lines {
		sum += db.CountCodeSteps(code) * input.ParseInt(code[:3])
	}
	return sum
}

var (
	DpadGap = grid.Position{Row: 0, Col: 0}
	DpadU   = grid.Position{Row: 0, Col: 1}
	DpadA   = grid.Position{Row: 0, Col: 2}

	DpadL = grid.Position{Row: 1, Col: 0}
	DpadD = grid.Position{Row: 1, Col: 1}
	DpadR = grid.Position{Row: 1, Col: 2}

	DpadNone = grid.Position{Row: -1, Col: -1}
)

type Memo struct {
	Current     grid.Position
	Destination grid.Position
	Depth       int
}

type DpadBots struct {
	MaxDepth int
	Memo     map[Memo]int
}

func newDpadBots(maxDepth int) DpadBots {
	return DpadBots{maxDepth, make(map[Memo]int)}
}

func (db DpadBots) CountCodeSteps(code string) int {
	sum := 0
	current := NumA
	for _, key := range code {
		next := parseKey(key)
		sum += db.CountSteps(current, next, 0) + 1
		current = next
	}
	return sum
}

func (db DpadBots) CountSteps(current, destination grid.Position, depth int) int {
	if steps, ok := db.Memo[Memo{current, destination, depth}]; ok {
		return steps
	}
	if depth == db.MaxDepth {
		steps := current.Distance(destination)
		db.Memo[Memo{current, destination, depth}] = steps
		return steps
	}
	depth++

	dx := DpadNone
	if destination.Col < current.Col {
		dx = DpadL
	} else if destination.Col > current.Col {
		dx = DpadR
	}

	dy := DpadNone
	if destination.Row < current.Row {
		dy = DpadU
	} else if destination.Row > current.Row {
		dy = DpadD
	}

	var shortest int
	if dx == DpadNone && dy == DpadNone {
		shortest = 0
	} else if dx == DpadNone {
		shortest = db.CountSteps(DpadA, dy, depth) + db.CountSteps(dy, DpadA, depth)
	} else if dy == DpadNone {
		shortest = db.CountSteps(DpadA, dx, depth) + db.CountSteps(dx, DpadA, depth)
	} else {
		xFirst := db.CountSteps(DpadA, dx, depth) + db.CountSteps(dx, dy, depth) + db.CountSteps(dy, DpadA, depth)
		yFirst := db.CountSteps(DpadA, dy, depth) + db.CountSteps(dy, dx, depth) + db.CountSteps(dx, DpadA, depth)

		gap := DpadGap
		if depth == 1 {
			gap = NumGap
		}
		if (current.Col == gap.Col || destination.Col == gap.Col) && (current.Row == gap.Row || destination.Row == gap.Row) {
			// One of them will pass the gap
			if current.Col == gap.Col {
				shortest = xFirst
			} else {
				shortest = yFirst
			}
		} else {
			shortest = int(math.Min(float64(xFirst), float64(yFirst)))
		}
	}

	steps := shortest + current.Distance(destination)
	db.Memo[Memo{current, destination, depth - 1}] = steps
	return steps
}

var (
	Num7 = grid.Position{Row: 0, Col: 0}
	Num8 = grid.Position{Row: 0, Col: 1}
	Num9 = grid.Position{Row: 0, Col: 2}

	Num4 = grid.Position{Row: 1, Col: 0}
	Num5 = grid.Position{Row: 1, Col: 1}
	Num6 = grid.Position{Row: 1, Col: 2}

	Num1 = grid.Position{Row: 2, Col: 0}
	Num2 = grid.Position{Row: 2, Col: 1}
	Num3 = grid.Position{Row: 2, Col: 2}

	NumGap = grid.Position{Row: 3, Col: 0}
	Num0   = grid.Position{Row: 3, Col: 1}
	NumA   = grid.Position{Row: 3, Col: 2}

	NumNone = grid.Position{Row: -1, Col: -1}
)

func parseKey(key rune) grid.Position {
	switch key {
	case 'A':
		return NumA
	case '0':
		return Num0
	case '1':
		return Num1
	case '2':
		return Num2
	case '3':
		return Num3
	case '4':
		return Num4
	case '5':
		return Num5
	case '6':
		return Num6
	case '7':
		return Num7
	case '8':
		return Num8
	case '9':
		return Num9
	default:
		return NumNone
	}
}
