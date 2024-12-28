package day3

import (
	"strings"
	"testing"
)

const example3 = `467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`

func TestDay3a(t *testing.T) {
	want := 4361
	res := day3a(strings.Split(example3, "\n"))

	if res != want {
		t.Fatalf("day3a() = %d, want %d", res, want)
	}
}

func TestDay3b(t *testing.T) {
	want := 467835
	res := day3b(strings.Split(example3, "\n"))

	if res != want {
		t.Fatalf("day3b() = %d, want %d", res, want)
	}
}
