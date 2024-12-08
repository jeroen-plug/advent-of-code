package day8

import (
	"fmt"
	"slices"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Pos [2]int

func Day8() {
	lines := input.Lines(8)

	fmt.Printf("day 8a: %d\n", day8a(lines))
	fmt.Printf("day 8b: %d\n", day8b(lines))
}

func day8a(lines []string) int {
	byFrequency := findAntennas(lines)

	var antinodes []Pos
	for _, antennas := range byFrequency {
		if len(antennas) < 1 {
			continue
		}

		for _, a := range findAntinodes(lines, antennas, true) {
			if !slices.Contains(antinodes, a) {
				antinodes = append(antinodes, a)
			}
		}
	}

	return len(antinodes)
}

func day8b(lines []string) int {
	byFrequency := findAntennas(lines)

	var antinodes []Pos
	for _, antennas := range byFrequency {
		if len(antennas) < 1 {
			continue
		}

		for _, a := range findAntinodes(lines, antennas, false) {
			if !slices.Contains(antinodes, a) {
				antinodes = append(antinodes, a)
			}
		}
		for _, a := range antennas {
			if len(a) > 1 && !slices.Contains(antinodes, a) {
				antinodes = append(antinodes, a)
			}
		}
	}

	return len(antinodes)
}

func findAntinodes(lines []string, antennas []Pos, limit bool) []Pos {
	var antinodes []Pos
	for remaining := slices.Clone(antennas); len(remaining) > 0; remaining = remaining[1:] {
		current := remaining[0]
		for _, next := range remaining[1:] {
			if limit {
				antinodes = append(antinodes, antinodesForPair(lines, current, next)...)
			} else {
				antinodes = append(antinodes, allAntinodesForPair(lines, current, next)...)
			}
		}
	}

	return antinodes
}

func antinodesForPair(lines []string, a, b Pos) []Pos {
	dl := a[0] - b[0]
	dc := a[1] - b[1]

	var antinodes []Pos
	if new := (Pos{a[0] + dl, a[1] + dc}); inBounds(lines, new) {
		antinodes = append(antinodes, new)
	}
	if new := (Pos{b[0] - dl, b[1] - dc}); inBounds(lines, new) {
		antinodes = append(antinodes, new)
	}

	return antinodes
}

func allAntinodesForPair(lines []string, a, b Pos) []Pos {
	dl := a[0] - b[0]
	dc := a[1] - b[1]

	var antinodes []Pos
	for new := (Pos{a[0] + dl, a[1] + dc}); inBounds(lines, new); new = (Pos{new[0] + dl, new[1] + dc}) {
		antinodes = append(antinodes, new)
	}
	for new := (Pos{b[0] - dl, b[1] - dc}); inBounds(lines, new); new = (Pos{new[0] - dl, new[1] - dc}) {
		antinodes = append(antinodes, new)
	}

	return antinodes
}

func inBounds(lines []string, pos Pos) bool {
	return pos[0] >= 0 && pos[0] < len(lines) && pos[1] >= 0 && pos[1] < len(lines[0])
}

func findAntennas(lines []string) [][]Pos {
	antennas := make([][]Pos, 127)

	for row, line := range lines {
		for col, c := range line {
			if c != '.' {
				antennas[c] = append(antennas[c], Pos{row, col})
			}
		}
	}

	return antennas
}
