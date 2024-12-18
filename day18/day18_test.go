package day18

import (
	"strings"
	"testing"

	"github.com/jeroen-plug/advent-of-code-2024/grid"
)

const example = `5,4
4,2
4,5
3,0
2,1
6,3
2,4
1,5
0,6
3,3
2,6
5,1
1,2
5,5
2,5
6,5
1,4
0,4
6,4
1,1
6,1
1,0
0,5
1,6
2,0`

func TestDay18a(t *testing.T) {
	want := 22
	res := day18a(strings.Split(example, "\n")[:12], 7)

	if res != want {
		t.Fatalf("day18a() = %d, want %d", res, want)
	}
}

func TestDay18b(t *testing.T) {
	want := grid.Position{Row: 1, Col: 6}
	res := day18b(strings.Split(example, "\n"), 7)

	if res != want {
		t.Fatalf("day18b() = %d, want %d", res, want)
	}
}
