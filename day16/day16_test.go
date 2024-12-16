package day16

import (
	"strings"
	"testing"
)

const exampleSmall = `###############
#.......#....E#
#.#.###.#.###.#
#.....#.#...#.#
#.###.#####.#.#
#.#.#.......#.#
#.#.#####.###.#
#...........#.#
###.#.#####.#.#
#...#.....#.#.#
#.#.#.###.#.#.#
#.....#...#.#.#
#.###.#.#.#.#.#
#S..#.....#...#
###############`

const exampleLarge = `#################
#...#...#...#..E#
#.#.#.#.#.#.#.#.#
#.#.#.#...#...#.#
#.#.#.#.###.#.#.#
#...#.#.#.....#.#
#.#.#.#.#.#####.#
#.#...#.#.#.....#
#.#.#####.#.###.#
#.#.#.......#...#
#.#.###.#####.###
#.#.#...#.....#.#
#.#.#.#####.###.#
#.#.#.........#.#
#.#.#.#########.#
#S#.............#
#################`

var tests = []struct {
	name  string
	input []string
	wantA int
	wantB int
}{
	{"small", strings.Split(exampleSmall, "\n"), 7036, 45},
	{"large", strings.Split(exampleLarge, "\n"), 11048, 64},
}

func TestDay16a(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := day16a(tt.input)
			if res != tt.wantA {
				t.Fatalf("day16a() = %d, want %d", res, tt.wantA)
			}
		})
	}
}

func TestDay16b(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := day16b(tt.input)
			if res != tt.wantB {
				t.Fatalf("day16b() = %d, want %d", res, tt.wantB)
			}
		})
	}
}
