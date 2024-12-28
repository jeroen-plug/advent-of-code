package day17

import (
	"strings"
	"testing"
)

const example = `2413432311323
3215453535623
3255245654254
3446585845452
4546657867536
1438598798454
4457876987766
3637877979653
4654967986887
4564679986453
1224686865563
2546548887735
4322674655533`

func TestDay17a(t *testing.T) {
	want := 102
	res := day17a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day17a() = %d, want %d", res, want)
	}
}

func TestDay17b(t *testing.T) {
	want := 94
	res := day17b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day17b() = %d, want %d", res, want)
	}
}
