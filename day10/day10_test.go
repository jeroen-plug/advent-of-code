package day10

import (
	"strings"
	"testing"
)

const example = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestDay10a(t *testing.T) {
	want := 36
	res := day10a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day10a() = %d, want %d", res, want)
	}
}

func TestDay10b(t *testing.T) {
	want := 81
	res := day10b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day10b() = %d, want %d", res, want)
	}
}
