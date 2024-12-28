package day15

import (
	"log"
	"slices"
	"strings"
	"unicode"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

type Box []Lens

type Lens struct {
	Label       string
	FocalLength int
}

func Solution() (any, any) {
	lines := input.Lines(15)
	return day15a(lines), day15b(lines)
}

func day15a(lines []string) int {
	sum := 0
	for _, step := range strings.Split(lines[0], ",") {
		sum += hash(step)
	}
	return sum
}

func day15b(lines []string) int {
	var hashmap [256]Box

	for _, step := range strings.Split(lines[0], ",") {
		operation, lens := parseStep(step)
		box := hash(lens.Label)

		switch operation {
		case '=':
			hashmap[box] = addLens(hashmap[box], lens)
		case '-':
			hashmap[box] = removeLens(hashmap[box], lens)
		default:
			log.Fatalf("Invalid operation: '%s'", string(operation))
		}
	}

	focusingPower := 0
	for box, lenses := range hashmap {
		for slot, lens := range lenses {
			focusingPower += (box + 1) * (slot + 1) * lens.FocalLength
		}
	}

	return focusingPower
}

func addLens(box Box, lens Lens) Box {
	if existing := slices.IndexFunc(box, func(l Lens) bool { return l.Label == lens.Label }); existing >= 0 {
		box[existing] = lens
	} else {
		box = append(box, lens)
	}
	return box
}

func removeLens(box Box, lens Lens) Box {
	return slices.DeleteFunc(box, func(l Lens) bool { return l.Label == lens.Label })
}

func parseStep(step string) (rune, Lens) {
	var label strings.Builder
	var operation rune
	var focalLength int

	for _, r := range step {
		if unicode.IsLetter(r) {
			label.WriteRune(r)
		} else if unicode.IsNumber(r) {
			focalLength = input.ParseInt(string(r))
		} else {
			operation = r
		}
	}

	return operation, Lens{label.String(), focalLength}
}

func hash(s string) int {
	h := 0
	for _, c := range s {
		h = ((h + int(c)) * 17) % 256
	}
	return h
}
