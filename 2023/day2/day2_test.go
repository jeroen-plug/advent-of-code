package day2

import (
	"reflect"
	"strings"
	"testing"
)

const example2 = `Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`

func TestDay2parse(t *testing.T) {
	lines := strings.Split(example2, "\n")

	want := day2game{1, []day2set{
		{4, 0, 3}, {1, 2, 6}, {0, 2, 0},
	}}
	res := day2parse(lines[0])

	if !reflect.DeepEqual(res, want) {
		t.Fatalf("day2parse() = %d, want %d", res, want)
	}
}

func TestDay2possible(t *testing.T) {
	lines := strings.Split(example2, "\n")

	if !day2possible(day2parse(lines[0]), day2set{12, 13, 14}) {
		t.Fatalf("day2possible(1) = false, want true")
	}
	if !day2possible(day2parse(lines[1]), day2set{12, 13, 14}) {
		t.Fatalf("day2possible(2) = false, want true")
	}
	if !day2possible(day2parse(lines[4]), day2set{12, 13, 14}) {
		t.Fatalf("day2possible(5) = false, want true")
	}

	if day2possible(day2parse(lines[2]), day2set{12, 13, 14}) {
		t.Fatalf("day2possible(3) = true, want false")
	}
	if day2possible(day2parse(lines[3]), day2set{12, 13, 14}) {
		t.Fatalf("day2possible(4) = true, want false")
	}
}

func TestDay2power(t *testing.T) {
	lines := strings.Split(example2, "\n")
	want := []int{48, 12, 1560, 630, 36}

	for i, line := range lines {
		power := day2power(day2parse(line))
		if power != want[i] {
			t.Fatalf("day2power(%d) = %d, want %d", i+1, power, want[i])
		}
	}
}
