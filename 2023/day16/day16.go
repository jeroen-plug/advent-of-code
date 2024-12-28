package day16

import (
	"log"
	"slices"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

type Position [2]int
type PosDir [3]int

const (
	Left int = iota
	Up
	Right
	Down
)

func Solution() (any, any) {
	lines := input.Lines(16)
	return day16a(lines), day16b(lines)
}

func day16a(lines []string) int {
	return calculateEnergized(lines, PosDir{0, 0, Right})
}

func day16b(lines []string) int {
	max := make(chan int)

	go func() {
		m := 0
		for l := range lines {
			if new := calculateEnergized(lines, PosDir{l, 0, Right}); new > m {
				m = new
			}
		}
		max <- m
	}()
	go func() {
		m := 0
		for l := range lines {
			if new := calculateEnergized(lines, PosDir{l, len(lines[0]) - 1, Left}); new > m {
				m = new
			}
		}
		max <- m
	}()
	go func() {
		m := 0
		for c := range lines[0] {
			if new := calculateEnergized(lines, PosDir{0, c, Down}); new > m {
				m = new
			}
		}
		max <- m
	}()
	go func() {
		m := 0
		for c := range lines[0] {
			if new := calculateEnergized(lines, PosDir{len(lines) - 1, c, Up}); new > m {
				m = new
			}
		}
		max <- m
	}()

	m1, m2, m3, m4 := <-max, <-max, <-max, <-max
	return slices.Max([]int{m1, m2, m3, m4})
}

func calculateEnergized(lines []string, initial PosDir) int {
	var (
		stack     []PosDir
		visited   []PosDir
		energized []Position
	)

	stack = append(stack, initial)
	visited = append(visited, initial)
	energized = append(energized, Position{initial[0], initial[1]})

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		var next []PosDir

		switch lines[current[0]][current[1]] {
		case '.':
			next = append(next, move(current, current[2]))

		case '/':
			switch current[2] {
			case Left:
				next = append(next, move(current, Down))
			case Up:
				next = append(next, move(current, Right))
			case Right:
				next = append(next, move(current, Up))
			case Down:
				next = append(next, move(current, Left))
			default:
				log.Fatalln("Unknown dir: ", current[2])
			}
		case '\\':
			switch current[2] {
			case Left:
				next = append(next, move(current, Up))
			case Up:
				next = append(next, move(current, Left))
			case Right:
				next = append(next, move(current, Down))
			case Down:
				next = append(next, move(current, Right))
			default:
				log.Fatalln("Unknown dir: ", current[2])
			}

		case '|':
			if current[2]%2 == 0 {
				next = append(next, move(current, Up))
				next = append(next, move(current, Down))
			} else {
				next = append(next, move(current, current[2]))
			}
		case '-':
			if current[2]%2 == 0 {
				next = append(next, move(current, current[2]))
			} else {
				next = append(next, move(current, Left))
				next = append(next, move(current, Right))
			}

		default:
			log.Fatalln("Unknown object at: ", current)
		}

		for _, n := range next {
			if inBounds(lines, n) && !slices.Contains(visited, n) {
				stack = append(stack, n)
				visited = append(visited, n)

				if !slices.Contains(energized, Position{n[0], n[1]}) {
					energized = append(energized, Position{n[0], n[1]})
				}
			}
		}
	}

	return len(energized)
}

func move(pos PosDir, dir int) PosDir {
	newPos := PosDir{pos[0], pos[1], dir}
	switch dir {
	case Left:
		newPos[1]--
	case Up:
		newPos[0]--
	case Right:
		newPos[1]++
	case Down:
		newPos[0]++
	default:
		log.Fatalln("Unknown dir: ", dir)
	}
	return newPos
}

func inBounds(lines []string, pos PosDir) bool {
	return pos[0] >= 0 && pos[0] < len(lines) && pos[1] >= 0 && pos[1] < len(lines[0])
}
