package day9

import (
	"strings"
	"testing"
)

const example = `0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`

func TestDay9a(t *testing.T) {
	want := 114
	res := day9a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day9a() = %d, want %d", res, want)
	}
}

func TestDay9b(t *testing.T) {
	want := 2
	res := day9b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day9b() = %d, want %d", res, want)
	}
}
