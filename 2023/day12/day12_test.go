package day12

import (
	"strings"
	"testing"
)

const example = `???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`

func TestDay12a(t *testing.T) {
	want := 21
	res := day12a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day12a() = %d, want %d", res, want)
	}
}

func TestDay12b(t *testing.T) {
	want := 525152
	res := day12b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day12b() = %d, want %d", res, want)
	}
}
