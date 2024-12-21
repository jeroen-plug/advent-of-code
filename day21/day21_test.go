package day21

import (
	"strings"
	"testing"
)

const example = `029A
980A
179A
456A
379A`

func TestDay21a(t *testing.T) {
	want := 126384
	res := day21a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day21a() = %d, want %d", res, want)
	}
}

func TestDay21b(t *testing.T) {
	want := 154115708116294
	res := day21b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day21b() = %d, want %d", res, want)
	}
}
