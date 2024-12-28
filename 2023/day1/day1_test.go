package day1

import (
	"testing"
)

func TestDay1a(t *testing.T) {
	want := 142
	res := day1a(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`)

	if res != want {
		t.Fatalf("day1a() = %d, want %d", res, want)
	}
}

func TestDay1b(t *testing.T) {
	want := 281
	res := day1b(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`)

	if res != want {
		t.Fatalf("day1b() = %d, want %d", res, want)
	}
}
