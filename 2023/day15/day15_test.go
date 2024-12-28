package day15

import (
	"strings"
	"testing"
)

const example = `rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`

func TestDay15a(t *testing.T) {
	want := 1320
	res := day15a(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day15a() = %d, want %d", res, want)
	}
}

func TestDay15b(t *testing.T) {
	want := 145
	res := day15b(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("day15b() = %d, want %d", res, want)
	}
}
