package day14

import (
	"strings"
	"testing"
)

const example = `O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`

func TestDay14a(t *testing.T) {
	want := 136
	res := day14a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day14a() = %d, want %d", res, want)
	}
}

func TestDay14b(t *testing.T) {
	want := 64
	res := day14b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day14b() = %d, want %d", res, want)
	}
}
