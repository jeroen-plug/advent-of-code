package day13

import (
	"math"
	"slices"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

func Solution() (any, any) {
	lines := input.Lines(13)
	return day13a(lines), day13b(lines)
}

func day13a(lines []string) int {
	patterns := parse(lines)

	sum := 0
	for _, p := range patterns {
		byRow, byCol := toBits(p)

		rows := findSymmetry(byRow, 0)
		cols := findSymmetry(byCol, 0)

		sum += 100*rows + cols
	}

	return sum
}

func day13b(lines []string) int {
	patterns := parse(lines)

	sum := 0
	for _, p := range patterns {
		byRow, byCol := toBits(p)
		refRows := findSymmetry(byRow, 0)
		refCols := findSymmetry(byCol, 0)
		newRows := 0
		newCols := 0

	out:
		for row, theRow := range byRow {
			for col, theCol := range byCol {
				byNewRow := slices.Clone(byRow)
				byNewCol := slices.Clone(byCol)

				byNewRow[row] = theRow ^ (1 << col)
				byNewCol[col] = theCol ^ (1 << row)

				newRows = findSymmetry(byNewRow, refRows)
				newCols = findSymmetry(byNewCol, refCols)

				if newRows+newCols > 0 {
					sum += newRows*100 + newCols
					break out
				}
			}
		}
	}

	return sum
}

func findSymmetry(pattern []uint32, ignore int) int {
	var last uint32 = math.MaxInt32
	splits := []int{0}
	for split, slice := range pattern {
		if slice == last && isSymmetric(pattern, split) && split != ignore {
			splits = append(splits, split)
		}
		last = slice
	}
	return splits[len(splits)-1]
}

func isSymmetric(pattern []uint32, split int) bool {
	size := int(math.Min(float64(split), float64(len(pattern)-split)))

	left := pattern[split-size : split]
	right := slices.Clone(pattern[split : split+size])
	slices.Reverse(right)

	return slices.Equal(left, right)
}

func toBits(pattern []string) ([]uint32, []uint32) {
	byRow := make([]uint32, len(pattern))
	byCol := make([]uint32, len(pattern[0]))

	for row, inRow := range pattern {
		for col, symbol := range inRow {
			byRow[row] <<= 1
			byCol[col] <<= 1

			if symbol == '#' {
				byRow[row]++
				byCol[col]++
			}
		}
	}

	return byRow, byCol
}

func parse(lines []string) [][]string {
	var patterns [][]string
	var pattern []string
	for _, l := range lines {
		if l == "" {
			patterns = append(patterns, pattern)
			pattern = []string{}
		} else {
			pattern = append(pattern, l)
		}
	}
	return append(patterns, pattern)
}
