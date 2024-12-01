package day01

import (
	"strings"
	"testing"
)

const example = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestDay1a(t *testing.T) {
	want := 11
	res := day1a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day1a() = %d, want %d", res, want)
	}
}

func TestDay1b(t *testing.T) {
	want := 31
	res := day1b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day1b() = %d, want %d", res, want)
	}
}
