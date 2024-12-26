package day6

import (
	"strings"
	"testing"
)

const example = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestDay6a(t *testing.T) {
	want := 41
	res := day6a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day6a() = %d, want %d", res, want)
	}
}

func TestDay6b(t *testing.T) {
	want := 6
	res := day6b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day6b() = %d, want %d", res, want)
	}
}
