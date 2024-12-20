package day20

import (
	"fmt"
	"strings"
	"testing"
)

const example = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestDay20(t *testing.T) {
	lines := strings.Split(example, "\n")
	var tests = []struct {
		maxCheat  int
		threshold int
		want      int
	}{
		{2, 0, 44},
		{2, 10, 10},
		{2, 50, 1},
		{20, 50, 285},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("max-%d/threshold-%d)", tt.maxCheat, tt.threshold), func(t *testing.T) {
			res := day20(lines, tt.maxCheat, tt.threshold)
			if res != tt.want {
				t.Fatalf("got %d, want %d", res, tt.want)
			}
		})
	}
}

func TestCalculateScores(t *testing.T) {
	want := 84

	maze, start, end := parse(strings.Split(example, "\n"))
	maze.CalculateScoresTo(end)
	res := maze[start].Score

	if res != want {
		t.Fatalf("maze.CalculateScoresTo() = %d, want %d", res, want)
	}
}
