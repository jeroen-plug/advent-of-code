package day4

import (
	"strings"
	"testing"
)

const example = `MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`

func TestDay4a(t *testing.T) {
	want := 18
	res := day4a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day4a() = %d, want %d", res, want)
	}
}

func TestDay4b(t *testing.T) {
	want := 9
	res := day4b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day4b() = %d, want %d", res, want)
	}
}
