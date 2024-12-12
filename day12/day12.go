package day12

import (
	"fmt"
	"math"
	"slices"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Region struct {
	Plant     byte
	Plots     []Plot
	Perimeter int
}

type Plot struct {
	Position   Position
	Perimeters int
}

func Day12() {
	lines := input.Lines(12)

	fmt.Printf("day 12a: %d\n", day12a(lines))
	fmt.Printf("day 12b: %d\n", day12b(lines))
}

func day12a(lines []string) int {
	sum := 0
	var used [][2]int
	for row, line := range lines {
		for col := range line {
			if !slices.Contains(used, [2]int{row, col}) {
				r := findRegion(lines, &used, Position{row, col})
				sum += len(r.Plots) * r.Perimeter
			}
		}
	}
	return sum
}

func day12b(lines []string) int {
	sum := 0
	var used [][2]int
	for row, line := range lines {
		for col := range line {
			if !slices.Contains(used, [2]int{row, col}) {
				r := findRegion(lines, &used, Position{row, col})
				sum += len(r.Plots) * countEdges(r)
			}
		}
	}
	return sum
}

func findRegion(lines []string, used *[][2]int, pos Position) Region {
	plant := lines[pos.Row][pos.Col]
	*used = append(*used, pos.Array())
	toCheck := []Position{pos}

	r := Region{Plant: plant}

	for len(toCheck) > 0 {
		current := toCheck[0]
		perimeters := 0
		for _, d := range AllDirections() {
			newPos := current.Move(d)
			if !newPos.InBounds(lines) || lines[current.Row][current.Col] != lines[newPos.Row][newPos.Col] {
				perimeters++
			} else if !slices.Contains(*used, newPos.Array()) {
				toCheck = append(toCheck, newPos)
				*used = append(*used, newPos.Array())
			}
		}
		r.Plots = append(r.Plots, Plot{current, perimeters})
		r.Perimeter += perimeters
		toCheck = toCheck[1:]
	}

	return r
}

func countEdges(r Region) int {
	var left, top, right, bottom int
	var plots [][2]int
	for _, p := range r.Plots {
		left = int(math.Min(float64(left), float64(p.Position.Col)))
		top = int(math.Min(float64(top), float64(p.Position.Row)))
		right = int(math.Max(float64(right), float64(p.Position.Col)))
		bottom = int(math.Max(float64(bottom), float64(p.Position.Row)))
		plots = append(plots, p.Position.Array())
	}

	edges := 0
	inside := false
	var last_candidates []float64

	// scan horizontal
	for row := top; row <= bottom; row++ {
		var candidates []float64
		for col := left; col <= right; col++ {
			in := slices.Contains(plots, [2]int{row, col})
			if inside != in {
				inside = in
				if inside {
					candidates = append(candidates, float64(col))
				} else {
					candidates = append(candidates, float64(col)-0.1)
				}
			}
		}
		if inside {
			candidates = append(candidates, float64(right+1))
		}
		for _, c := range candidates {
			if !slices.Contains(last_candidates, c) {
				edges++
			}
		}
		last_candidates = candidates
		inside = false
	}

	// scan vertical
	last_candidates = []float64{}
	for col := left; col <= right; col++ {
		var candidates []float64
		for row := top; row <= bottom; row++ {
			in := slices.Contains(plots, [2]int{row, col})
			if inside != in {
				inside = in
				if inside {
					candidates = append(candidates, float64(row))
				} else {
					candidates = append(candidates, float64(row)-0.1)
				}
			}
		}
		if inside {
			candidates = append(candidates, float64(bottom+1))
		}
		for _, c := range candidates {
			if !slices.Contains(last_candidates, c) {
				edges++
			}
		}
		last_candidates = candidates
		inside = false
	}

	return edges
}
