package day10

import (
	"strings"
	"testing"
)

const example = `..F7.
.FJ|.
SJ.L7
|F--J
LJ...`

const example2 = `FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`

func TestDay10a(t *testing.T) {
	want := 8
	res := day10a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day10a() = %d, want %d", res, want)
	}
}

func TestDay10b(t *testing.T) {
	want := 10
	res := day10b(strings.Split(example2, "\n"))

	if res != want {
		t.Fatalf("day10b() = %d, want %d", res, want)
	}
}
