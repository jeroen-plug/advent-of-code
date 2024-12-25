package day25

import (
	"strings"
	"testing"
)

const example = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`

func TestDay25a(t *testing.T) {
	want := 3
	res := day25a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day25a() = %d, want %d", res, want)
	}
}
