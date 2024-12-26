package day8

import (
	"strings"
	"testing"
)

const example = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestDay8a(t *testing.T) {
	want := 14
	res := day8a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day8a() = %d, want %d", res, want)
	}
}

func TestDay8b(t *testing.T) {
	want := 34
	res := day8b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day8b() = %d, want %d", res, want)
	}
}
