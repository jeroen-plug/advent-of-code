package day16

import (
	"strings"
	"testing"
)

const example = `.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`

func TestDay16a(t *testing.T) {
	want := 46
	res := day16a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day16a() = %d, want %d", res, want)
	}
}

func TestDay16b(t *testing.T) {
	want := 51
	res := day16b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day16b() = %d, want %d", res, want)
	}
}
