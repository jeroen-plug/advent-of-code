package day15

import (
	"strings"
	"testing"
)

const largeExample = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^`

const smallExample = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v>>v<<`

const wideExample = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

func TestDay15a(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  int
	}{
		{"small", strings.Split(smallExample, "\n"), 2028},
		{"large", strings.Split(largeExample, "\n"), 10092},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := day15a(tt.input)
			if res != tt.want {
				t.Fatalf("day15a() = %d, want %d", res, tt.want)
			}
		})
	}
}

func TestDay15b(t *testing.T) {
	var tests = []struct {
		name  string
		input []string
		want  int
	}{
		{"wide", strings.Split(wideExample, "\n"), 618},
		{"large", strings.Split(largeExample, "\n"), 9021},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := day15b(tt.input)
			if res != tt.want {
				t.Fatalf("day15b() = %d, want %d", res, tt.want)
			}
		})
	}
}
