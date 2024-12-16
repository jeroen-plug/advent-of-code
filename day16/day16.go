package day16

import (
	"container/heap"
	"fmt"
	"log"
	"math"
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/grid"
	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type PosDir struct {
	Row int
	Col int
	Dir grid.Direction
}

func ToPosDir(pos grid.Position, dir grid.Direction) PosDir {
	return PosDir{
		Row: pos.Row,
		Col: pos.Col,
		Dir: dir,
	}
}

func Day16() {
	lines := input.Lines(16)

	fmt.Printf("day 16a: %d\n", day16a(lines))
	fmt.Printf("day 16b: %d\n", day16b(lines))
}

func day16a(lines []string) int {
	maze, start, end := separateStartEnd(lines)
	cost, _ := aStar(maze, ToPosDir(start, grid.Right), end, distance)
	return cost
}

func day16b(lines []string) int {
	maze, start, end := separateStartEnd(lines)
	_, tiles := aStar(maze, ToPosDir(start, grid.Right), end, distance)
	return tiles
}

func aStar(maze []string, start PosDir, end grid.Position, heuristic func(current PosDir, end grid.Position) int) (int, int) {
	openSet := make(PriorityQueue, 0)
	heap.Init(&openSet)
	heap.Push(&openSet, &Item{posDir: start})

	cameFrom := make(map[PosDir][]PosDir)

	gScore := make(map[PosDir]int)
	gScore[start] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*Item)
		if current.posDir.Row == end.Row && current.posDir.Col == end.Col {
			return gScore[current.posDir], countTiles(cameFrom, current.posDir)
		}

		for _, neighbor := range neighbors(maze, current.posDir) {
			tentativeGScore := gScore[current.posDir] + neighbor.cost
			if gs, ok := gScore[neighbor.posDir]; !ok || tentativeGScore < gs {
				cameFrom[neighbor.posDir] = []PosDir{current.posDir}
				gScore[neighbor.posDir] = tentativeGScore
				fScore := tentativeGScore + heuristic(neighbor.posDir, end)

				i := slices.IndexFunc(openSet, func(i *Item) bool { return i.posDir == neighbor.posDir })
				if i >= 0 {
					openSet.SetPriority(openSet[i], fScore)
				} else {
					heap.Push(&openSet, &Item{
						posDir:   neighbor.posDir,
						priority: fScore,
					})
				}
			} else if tentativeGScore == gScore[neighbor.posDir] {
				// Keep track of alternative routes
				cameFrom[neighbor.posDir] = append(cameFrom[neighbor.posDir], current.posDir)
			}
		}
	}

	log.Fatalln("Got lost, try again")
	return 0, 0
}

func countTiles(cameFrom map[PosDir][]PosDir, current PosDir) int {
	toCheck := []PosDir{current}
	var tiles []grid.Position

	for len(toCheck) > 0 {
		current = toCheck[0]
		position := grid.Position{Row: current.Row, Col: current.Col}
		if !slices.Contains(tiles, position) {
			tiles = append(tiles, position)
		}
		for _, next := range cameFrom[current] {
			if !slices.Contains(toCheck, next) {
				toCheck = append(toCheck, next)
			}
		}
		toCheck = toCheck[1:]
	}

	return len(tiles)
}

type Neighbor struct {
	posDir PosDir
	cost   int
}

func neighbors(maze []string, current PosDir) []Neighbor {
	left, right := current, current
	move := grid.Position{Row: current.Row, Col: current.Col}

	left.Dir = current.Dir.Turn(-1)
	right.Dir = current.Dir.Turn(1)
	move = move.Move(current.Dir)

	neighbors := []Neighbor{
		{left, 1000},
		{right, 1000},
	}

	if maze[move.Row][move.Col] == '.' {
		neighbors = append(neighbors, Neighbor{ToPosDir(move, current.Dir), 1})
	}

	return neighbors
}

func distance(a PosDir, b grid.Position) int {
	return int(math.Abs(float64(a.Col-b.Col)) + math.Abs(float64(a.Row-b.Row)))
}

func separateStartEnd(lines []string) ([]string, grid.Position, grid.Position) {
	var maze []string
	var start, end grid.Position

	for row, l := range lines {
		if col := strings.IndexRune(l, 'S'); col > 0 {
			start = grid.Position{Row: row, Col: col}
			l = strings.Replace(l, "S", ".", 1)
		}
		if col := strings.IndexRune(l, 'E'); col > 0 {
			end = grid.Position{Row: row, Col: col}
			l = strings.Replace(l, "E", ".", 1)
		}
		maze = append(maze, l)
	}

	return maze, start, end
}
