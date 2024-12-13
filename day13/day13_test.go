package day13

import (
	"strings"
	"testing"
)

const example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279`

func TestDay13a(t *testing.T) {
	want := 480
	res := day13a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day13a() = %d, want %d", res, want)
	}
}

func TestDay13b(t *testing.T) {
	want := 875318608908
	res := day13b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day13b() = %d, want %d", res, want)
	}
}
