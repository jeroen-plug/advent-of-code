package day13

import (
	"strings"
	"testing"
)

const example = `#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`

func TestDay13a(t *testing.T) {
	want := 405
	res := day13a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day13a() = %d, want %d", res, want)
	}
}

func TestDay13b(t *testing.T) {
	want := 400
	res := day13b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day13b() = %d, want %d", res, want)
	}
}
