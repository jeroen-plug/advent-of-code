package day8

import (
	"strings"
	"testing"
)

const exampleA = `LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`

const exampleB = `LR

AAA = (AAB, XXX)
AAB = (XXX, AAZ)
AAZ = (AAB, XXX)
BBA = (BBB, XXX)
BBB = (BBC, BBC)
BBC = (BBZ, BBZ)
BBZ = (BBB, BBB)
XXX = (XXX, XXX)`

func TestDay8a(t *testing.T) {
	want := 6
	res := day8a(strings.Split(exampleA, "\n"))

	if res != want {
		t.Fatalf("day8a() = %d, want %d", res, want)
	}
}

func TestDay8b(t *testing.T) {
	want := 6
	res := day8b(strings.Split(exampleB, "\n"))

	if res != want {
		t.Fatalf("day8b() = %d, want %d", res, want)
	}
}

func TestFindLcm(t *testing.T) {
	want := 252
	res := findLcm([]int{2, 7, 3, 9, 4})

	if res != want {
		t.Fatalf("findLcm([2, 7, 3, 9, 4]) = %d, want %d", res, want)
	}
}
