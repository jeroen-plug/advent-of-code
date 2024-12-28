package day11

import (
	"strings"
	"testing"
)

const example = `...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`

func TestDay11a(t *testing.T) {
	want := 374
	res := day11(strings.Split(example, "\n"), 2)

	if res != want {
		t.Fatalf("day11(2) = %d, want %d", res, want)
	}
}

func TestDay11b10(t *testing.T) {
	want := 1030
	res := day11(strings.Split(example, "\n"), 10)

	if res != want {
		t.Fatalf("day11(10) = %d, want %d", res, want)
	}
}

func TestDay11b100(t *testing.T) {
	want := 8410
	res := day11(strings.Split(example, "\n"), 100)

	if res != want {
		t.Fatalf("day11(100) = %d, want %d", res, want)
	}
}
