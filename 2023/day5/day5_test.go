package day5

import (
	"strings"
	"testing"
)

const example = `seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

func TestDay5a(t *testing.T) {
	want := 35
	res := day5a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day5a() = %d, want %d", res, want)
	}
}

func TestDay5b(t *testing.T) {
	want := 46
	res := day5b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day5b() = %d, want %d", res, want)
	}
}
