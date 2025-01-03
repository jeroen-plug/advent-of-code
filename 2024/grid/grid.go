package grid

import "math"

type Position struct {
	Row int
	Col int
}

func (p Position) Move(d Direction) Position {
	switch d {
	case Left:
		p.Col--
	case Up:
		p.Row--
	case Right:
		p.Col++
	case Down:
		p.Row++
	}
	return p
}

func (p Position) Add(offset Position) Position {
	p.Row += offset.Row
	p.Col += offset.Col
	return p
}

func (p Position) Distance(other Position) int {
	return int(math.Abs(float64(p.Col-other.Col)) + math.Abs(float64(p.Row-other.Row)))
}

func (p Position) InBounds(l []string) bool {
	return p.Row >= 0 && p.Row < len(l) && p.Col >= 0 && p.Col < len(l[0])
}

func (p Position) Array() [2]int {
	return [2]int{p.Row, p.Col}
}

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

func AllDirections() []Direction {
	return []Direction{Left, Up, Right, Down}
}

func (d Direction) Turn(offset int) Direction {
	d += Direction(offset)
	if d < 0 {
		d += 4
	}
	d %= 4
	return d
}
