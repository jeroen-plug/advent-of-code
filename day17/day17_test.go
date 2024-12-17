package day17

import (
	"strings"
	"testing"
)

const exampleA = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

// ADV 1 => A >>= 1
// OUT A
// JNZ

const exampleB = `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

// ADV 3 => A >>= 3
// OUT A
// JNZ

func TestDay17a(t *testing.T) {
	want := "4,6,3,5,6,3,5,2,1,0"
	res := day17a(strings.Split(exampleA, "\n"))

	if res != want {
		t.Fatalf("day17a() = %s, want %s", res, want)
	}
}

func TestDay17b(t *testing.T) {
	want := 117440
	res := day17b(strings.Split(exampleB, "\n"))

	if res != want {
		t.Fatalf("day17b() = %d, want %d", res, want)
	}
}
