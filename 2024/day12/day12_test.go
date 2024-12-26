package day12

import (
	"strings"
	"testing"
)

const example = `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`

func TestDay12a(t *testing.T) {
	want := 1930
	res := day12a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day12a() = %d, want %d", res, want)
	}
}

func TestDay12b(t *testing.T) {
	want := 1206
	res := day12b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day12b() = %d, want %d", res, want)
	}
}

func TestDay12bHoles(t *testing.T) {
	want := 368
	res := day12b(strings.Split(`AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`, "\n"))

	if res != want {
		t.Fatalf("day12b(holes) = %d, want %d", res, want)
	}
}
