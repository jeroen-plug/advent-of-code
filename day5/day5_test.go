package day5

import (
	"strings"
	"testing"
)

const example = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestDay5a(t *testing.T) {
	want := 143
	res := day5a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day5a() = %d, want %d", res, want)
	}
}

func TestDay5b(t *testing.T) {
	want := 123
	res := day5b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day5b() = %d, want %d", res, want)
	}
}
