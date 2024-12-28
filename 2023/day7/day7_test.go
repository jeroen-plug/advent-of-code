package day7

import (
	"strings"
	"testing"
)

const example = `32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483`

func TestDay7a(t *testing.T) {
	want := 6440
	res := day7a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day7a() = %d, want %d", res, want)
	}
}

func TestDay7b(t *testing.T) {
	want := 5905
	res := day7b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day7b() = %d, want %d", res, want)
	}
}
