package day19

import (
	"strings"
	"testing"
)

const example = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestDay19a(t *testing.T) {
	want := 6
	res := day19a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day19a() = %d, want %d", res, want)
	}
}

func TestDay19b(t *testing.T) {
	want := 16
	res := day19b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day19b() = %d, want %d", res, want)
	}
}
