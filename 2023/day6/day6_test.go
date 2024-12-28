package day6

import (
	"strings"
	"testing"
)

const example = `Time:      7  15   30
Distance:  9  40  200`

func TestDay6a(t *testing.T) {
	want := 288
	res := day6a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day6a() = %d, want %d", res, want)
	}
}

func TestDay6b(t *testing.T) {
	want := 71503
	res := day6b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day6b() = %d, want %d", res, want)
	}
}
