package day17

import (
	"container/heap"
	"log"
	"math"
	"slices"

	"github.com/jeroen-plug/advent-of-code/2023/grid"
	"github.com/jeroen-plug/advent-of-code/2023/input"
)

type State struct {
	Position    grid.Position
	Direction   grid.Direction
	Consecutive int
}

func Solution() (any, any) {
	lines := input.Lines(17)
	return day17a(lines), day17b(lines)
}

func day17a(lines []string) int {
	path := aStar(lines, grid.Position{Row: 0, Col: 0}, grid.Position{Row: len(lines) - 1, Col: len(lines[0]) - 1}, 0, 3)

	sum := 0
	for _, path := range path[1:] {
		sum += int(lines[path.Row][path.Col] - '0')
	}
	return sum
}

func day17b(lines []string) int {
	path := aStar(lines, grid.Position{Row: 0, Col: 0}, grid.Position{Row: len(lines) - 1, Col: len(lines[0]) - 1}, 4, 10)

	sum := 0
	for _, path := range path[1:] {
		sum += int(lines[path.Row][path.Col] - '0')
	}
	return sum
}

type GscoreDefault map[State]int

func (g GscoreDefault) Get(state State) int {
	if state.Position == (grid.Position{Row: 0, Col: 0}) {
		return 0
	}
	gScore, ok := g[state]
	if ok {
		return gScore
	} else {
		return math.MaxInt
	}
}

func (g GscoreDefault) Set(state State, score int) {
	g[state] = score
}

func aStar(lines []string, start grid.Position, goal grid.Position, minConsecutive, maxConsecutive int) []grid.Position {
	openSet := NewPriorityQueue()
	heap.Push(&openSet, &Item{state: State{Position: start, Direction: grid.Right}})

	cameFrom := make(map[State]State)
	gScore := make(GscoreDefault)

	for len(openSet) > 0 {
		item := heap.Pop(&openSet).(*Item)
		current := item.state
		if current.Position == goal && current.Consecutive >= minConsecutive {
			return reconstruct(cameFrom, current)
		}

		for _, n := range neighbors(lines, current, minConsecutive, maxConsecutive) {
			tentative_gScore := gScore.Get(current) + int(lines[n.Position.Row][n.Position.Col]-'0')

			if tentative_gScore < gScore.Get(n) {
				cameFrom[n] = current
				gScore.Set(n, tentative_gScore)
				fScore := tentative_gScore + n.Position.Distance(goal)
				i := slices.IndexFunc(openSet, func(i *Item) bool { return i.state == n })
				if i >= 0 {
					openSet.SetPriority(openSet[i], fScore)
				} else {
					heap.Push(&openSet, &Item{state: n, priority: fScore})
				}
			}
		}
	}

	log.Fatalln("Got lost, try again...")
	return nil
}

func neighbors(lines []string, current State, minConsecutive, maxConsecutive int) []State {
	var dirs []grid.Direction
	if current.Consecutive >= minConsecutive || current.Consecutive == 0 {
		dirs = append(dirs, current.Direction.Turn(-1))
		dirs = append(dirs, current.Direction.Turn(1))
	}
	if current.Consecutive < maxConsecutive {
		dirs = append(dirs, current.Direction)
	}

	var neighbors []State
	for _, dir := range dirs {
		neighbor := State{
			Position:    current.Position.Move(dir),
			Direction:   dir,
			Consecutive: current.Consecutive + 1,
		}
		if current.Direction != neighbor.Direction {
			neighbor.Consecutive = 1
		}
		if neighbor.Position.InBounds(lines) {
			neighbors = append(neighbors, neighbor)
		}
	}
	return neighbors
}

func reconstruct(cameFrom map[State]State, current State) []grid.Position {
	path := []grid.Position{current.Position}
	ok := true
	for {
		current, ok = cameFrom[current]
		if !ok {
			break
		}
		path = append(path, current.Position)
	}
	slices.Reverse(path)
	return path
}
