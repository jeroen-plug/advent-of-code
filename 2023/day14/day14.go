package day14

import (
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

type Direction int

const (
	Left Direction = iota
	Up
	Right
	Down
)

func Solution() (any, any) {
	lines := input.Lines(14)
	return day14a(lines), day14b(lines)
}

func day14a(lines []string) int {
	platform := slices.Clone(lines)
	return weight(slide(platform, Up))
}

func day14b(lines []string) int {
	p := slices.Clone(lines)

	cycles := []string{strings.Join(p, "-")}
	found := -1
	for {
		p = cycle(p)
		oneLine := strings.Join(p, "-")
		if found = slices.Index(cycles, oneLine); found >= 0 {
			break
		}
		cycles = append(cycles, oneLine)
	}

	offset := found
	size := len(cycles) - offset

	goal := 1000000000
	index := ((goal - offset) % size) + offset

	return weight(strings.Split(cycles[index], "-"))
}

func cycle(p []string) []string {
	p = slide(p, Up)
	p = slide(p, Left)
	p = slide(p, Down)
	p = slide(p, Right)
	return p
}

func slide(p []string, dir Direction) []string {
	if dir == Down {
		slices.Reverse(p)
	}
	for row := range p {
		if dir == Right {
			p[row] = reverse(p[row])
		}
		for col := range p[row] {
			if p[row][col] != '.' {
				continue
			}

			if dir == Up || dir == Down {
				for r, nextRow := range p[row+1:] {
					nextSpot := nextRow[col]
					if nextSpot == 'O' {
						p[row+1+r] = nextRow[:col] + "." + nextRow[col+1:]
						p[row] = p[row][:col] + "O" + p[row][col+1:]
						break
					} else if nextSpot == '#' {
						break
					}
				}
			} else {
				for c, nextSpot := range p[row][col+1:] {
					if nextSpot == 'O' {
						p[row] = p[row][:col] + "O" + p[row][col+1:col+1+c] + "." + p[row][col+2+c:]
						break
					} else if nextSpot == '#' {
						break
					}
				}
			}
		}
		if dir == Right {
			p[row] = reverse(p[row])
		}
	}
	if dir == Down {
		slices.Reverse(p)
	}
	return p
}

func weight(p []string) int {
	weight := 0
	for row := range p {
		for col := range p[row] {
			if p[row][col] == 'O' {
				weight += len(p) - row
			}
		}
	}
	return weight
}

func reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
