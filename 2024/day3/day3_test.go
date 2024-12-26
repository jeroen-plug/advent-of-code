package day3

import (
	"testing"
)

const example_a = `xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))`
const example_b = `xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))`

func TestDay3a(t *testing.T) {
	want := 161
	res := day3a(example_a)

	if res != want {
		t.Fatalf("day3a() = %d, want %d", res, want)
	}
}

func TestDay3b(t *testing.T) {
	want := 48
	res := day3b(example_b)

	if res != want {
		t.Fatalf("day3b() = %d, want %d", res, want)
	}
}
