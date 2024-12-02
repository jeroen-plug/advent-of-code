package day02

import (
	"strings"
	"testing"
)

const example = `7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`

func TestDay2a(t *testing.T) {
	want := 2
	res := day2a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day2a() = %d, want %d", res, want)
	}
}

func TestDay2b(t *testing.T) {
	want := 4
	res := day2b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day2b() = %d, want %d", res, want)
	}
}
