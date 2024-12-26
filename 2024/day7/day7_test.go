package day7

import (
	"strings"
	"testing"
)

const example = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestDay7a(t *testing.T) {
	want := 3749
	res := day7a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day7a() = %d, want %d", res, want)
	}
}

func TestDay7b(t *testing.T) {
	want := 11387
	res := day7b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day7b() = %d, want %d", res, want)
	}
}
