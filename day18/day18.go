package day18

import (
	"container/list"
	"fmt"
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/grid"
	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Day18() {
	lines := input.Lines(18)

	fmt.Printf("day 18a: %d\n", day18a(lines[:1024], 71))
	b := day18b(lines, 71)
	fmt.Printf("day 18b: %d,%d\n", b.Col, b.Row)
}

func day18a(lines []string, size int) int {
	walls := parse(lines)
	path, _ := bfs(walls, size)
	return len(path)
}

func day18b(lines []string, size int) grid.Position {
	walls := parse(lines)

	n := binarySearch(len(walls), func(i int) int {
		if _, ok := bfs(walls[:i], size); ok {
			return -1
		} else {
			return 1
		}
	})

	return walls[n]
}

func binarySearch(n int, compare func(i int) int) int {
	left := 0
	right := n - 1
	for left <= right {
		middle := (left + right) / 2
		cmp := compare(middle)
		if cmp < 0 {
			left = middle + 1
		} else if cmp > 0 {
			right = middle - 1
		}
	}
	return right
}

func bfs(walls []grid.Position, size int) ([]grid.Position, bool) {
	start := grid.Position{Row: 0, Col: 0}
	goal := grid.Position{Row: size - 1, Col: size - 1}

	toCheck := list.New()
	toCheck.PushBack(start)

	explored := make(map[grid.Position]grid.Position)

	for toCheck.Len() > 0 {
		e := toCheck.Front()
		current := e.Value.(grid.Position)
		toCheck.Remove(e)

		if current == goal {
			return reconstruct(explored, goal), true
		}

		for _, dir := range grid.AllDirections() {
			neighbor := current.Move(dir)
			if neighbor.Row < 0 || neighbor.Row >= size ||
				neighbor.Col < 0 || neighbor.Col >= size ||
				slices.Contains(walls, neighbor) ||
				neighbor == start {
				continue
			}

			if _, ok := explored[neighbor]; !ok {
				explored[neighbor] = current
				toCheck.PushBack(neighbor)
			}
		}
	}

	return nil, false
}

func reconstruct(explored map[grid.Position]grid.Position, current grid.Position) []grid.Position {
	var path []grid.Position
	for next, ok := explored[current]; ok; next, ok = explored[next] {
		path = append(path, next)
	}
	slices.Reverse(path)
	return path
}

func printMaze(walls []grid.Position, size int) {
	for row := range size {
		var s strings.Builder
		for col := range size {
			if slices.Contains(walls, grid.Position{Row: row, Col: col}) {
				s.WriteRune('#')
			} else {
				s.WriteRune('.')
			}
		}
		fmt.Println(s.String())
	}
}

func parse(lines []string) []grid.Position {
	var positions []grid.Position
	for _, l := range lines {
		n := strings.Split(l, ",")
		positions = append(positions, grid.Position{
			Row: input.ParseInt(n[1]),
			Col: input.ParseInt(n[0]),
		})
	}
	return positions
}
