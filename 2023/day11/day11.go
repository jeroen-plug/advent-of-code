package day11

import (
	"math"
	"slices"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

func Solution() (any, any) {
	lines := input.Lines(11)
	return day11(lines, 2), day11(lines, 1e6)
}

func day11(lines []string, expansion int) int {
	galaxies := findGalaxies(lines)
	colLookup, rowLookup := generateEmptyLookup(lines)

	sum := 0
	for i, g1 := range galaxies {
		for _, g2 := range galaxies[i:] {
			sum += distance(g1, g2) + (expansion-1)*expand(colLookup, rowLookup, g1, g2)
		}
	}

	return sum
}

func distance(a, b [2]int) int {
	return int(math.Abs(float64(a[0]-b[0])) + math.Abs(float64(a[1]-b[1])))
}

func expand(colLookup []int, rowLookup []int, a, b [2]int) int {
	return distance([2]int{colLookup[a[0]], rowLookup[a[1]]}, [2]int{colLookup[b[0]], rowLookup[b[1]]})
}

func findGalaxies(lines []string) [][2]int {
	var galaxies [][2]int
	for r, row := range lines {
		for c, col := range row {
			if col == '#' {
				galaxies = append(galaxies, [2]int{c, r})
			}
		}
	}
	return galaxies
}

func generateEmptyLookup(lines []string) ([]int, []int) {
	emptyCols, emptyRows := findEmpty(lines)
	var colLookup []int
	var rowLookup []int

	i := 0
	for c := range lines[0] {
		if slices.Contains(emptyCols, c) {
			i++
		}
		colLookup = append(colLookup, i)
	}
	i = 0
	for r := range lines {
		if slices.Contains(emptyRows, r) {
			i++
		}
		rowLookup = append(rowLookup, i)
	}

	return colLookup, rowLookup
}

func findEmpty(lines []string) ([]int, []int) {
	var emptyCols []int
	var emptyRows []int

	for r, row := range lines {
		rowEmpty := true
		for c, col := range row {
			if r == 0 {
				// init emptyCols
				emptyCols = append(emptyCols, c)
			}
			if col != '.' {
				rowEmpty = false
				emptyCols = slices.DeleteFunc(emptyCols, func(n int) bool { return n == c })
			}
		}
		if rowEmpty {
			emptyRows = append(emptyRows, r)
		}
	}

	return emptyCols, emptyRows
}
