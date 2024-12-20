package day20

import (
	"container/list"
	"fmt"
	"maps"
	"math"
	"slices"

	"github.com/jeroen-plug/advent-of-code-2024/grid"
	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Maze map[grid.Position]Tile

type Tile struct {
	Type  Type
	Score int
}

type Type int

const (
	T_Undefined Type = iota
	T_Track
	T_Wall
)

func Day20() {
	lines := input.Lines(20)

	fmt.Printf("day 20a: %d\n", day20(lines, 2, 100))
	fmt.Printf("day 20b: %d\n", day20(lines, 20, 100))
}

func day20(lines []string, maxCheat, threshold int) int {
	maze, _, end := parse(lines)
	maze.CalculateScoresTo(end)
	cheats := maze.FindCheats(maxCheat)
	// printCheats(cheats)

	sum := 0
	for time, n := range cheats {
		if time >= threshold {
			sum += n
		}
	}
	return sum
}

func (m Maze) FindCheats(maxCheat int) map[int]int {
	cheats := make(map[int]int)
	offsets := cheatOffsets(maxCheat)

	for current, tile := range m {
		if tile.Type != T_Track {
			continue
		}
		for offset, cost := range offsets {
			next := current.Add(offset)

			if m[next].Type == T_Track {
				saved := tile.Score - m[next].Score - cost
				if saved > 0 {
					cheats[saved]++
				}
			}
		}
	}

	return cheats
}

func (m Maze) CalculateScoresTo(end grid.Position) {
	toCheck := list.New()
	toCheck.PushBack(end)
	m.SetScore(end, 0)

	for toCheck.Len() > 0 {
		v := toCheck.Front()
		current := v.Value.(grid.Position)
		toCheck.Remove(v)

		for _, dir := range grid.AllDirections() {
			next := current.Move(dir)
			if m[next].Type == T_Track && m[next].Score == math.MaxInt {
				m.SetScore(next, m[current].Score+1)
				toCheck.PushBack(next)
			}
		}
	}
}

func (m Maze) SetScore(position grid.Position, score int) {
	tile := m[position]
	tile.Score = score
	m[position] = tile
}

func printCheats(cheats map[int]int) {
	for _, time := range slices.Sorted(maps.Keys(cheats)) {
		if cheats[time] == 1 {
			fmt.Printf("- There is one cheat that save %d picoseconds\n", time)
		} else {
			fmt.Printf("- There are %d cheats that save %d picoseconds\n", cheats[time], time)
		}
	}
}

func cheatOffsets(maxCheat int) map[grid.Position]int {
	offsets := make(map[grid.Position]int)
	for d := 2; d <= maxCheat; d++ {
		for dx := -d; dx <= d; dx++ {
			dy := d - int(math.Abs(float64(dx)))
			offsets[grid.Position{Row: dy, Col: dx}] = d
			if dy != 0 {
				offsets[grid.Position{Row: -dy, Col: dx}] = d
			}
		}
	}
	return offsets
}

func parse(lines []string) (Maze, grid.Position, grid.Position) {
	maze := make(Maze)
	var start, end grid.Position

	for row, line := range lines {
		for col, char := range line {
			current := grid.Position{Row: row, Col: col}
			tile := Tile{Score: math.MaxInt}
			switch char {
			case '#':
				tile.Type = T_Wall
			case '.':
				tile.Type = T_Track
			case 'S':
				start = current
				tile.Type = T_Track
			case 'E':
				end = current
				tile.Type = T_Track
			}
			maze[current] = tile
		}
	}

	return maze, start, end
}
