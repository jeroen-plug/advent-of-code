package day3

import (
	"bufio"
	"log"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

func Solution() (any, any) {
	f, err := os.Open(path.Join("input", "3.txt"))
	if err != nil {
		log.Fatal("Could not read input 3")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var grid []string

	for scanner.Scan() {
		grid = append(grid, scanner.Text())
	}

	return day3a(grid), day3b(grid)
}

func day3a(grid []string) int {
	var sum int
	var number strings.Builder
	isPart := false

	for ln, line := range grid {
		for col, char := range line {
			if isSymbol(byte(char)) || isSymbolUpDown(grid, ln, col) {
				isPart = true
			}

			if char >= '0' && char <= '9' {
				number.WriteRune(char)
				if col < len(line)-1 {
					continue
				}
			}

			if isPart && number.Len() > 0 {
				n, err := strconv.Atoi(number.String())
				if err != nil {
					log.Printf("Could not parse part number %s", number.String())
				}
				sum += n
			}

			isPart = isSymbol(byte(char)) || isSymbolUpDown(grid, ln, col)
			number.Reset()
		}
	}

	return sum
}

type day3gear struct {
	ln    int
	col   int
	parts []int
}

func createOrGetIndex(gears []day3gear, ln int, col int) ([]day3gear, int) {
	i := slices.IndexFunc(gears, func(g day3gear) bool {
		return g.ln == ln && g.col == col
	})

	if i < 0 {
		gears = append(gears, day3gear{ln, col, []int{}})
		i = len(gears) - 1
	}

	return gears, i
}

func day3b(grid []string) int {
	var gears []day3gear
	var number strings.Builder
	var currentGears []int
	var i int

	for ln, line := range grid {
		for col, char := range line {
			if isGearAt(grid, ln-1, col) {
				gears, i = createOrGetIndex(gears, ln-1, col)
				currentGears = append(currentGears, i)
			}
			if isGearAt(grid, ln, col) {
				gears, i = createOrGetIndex(gears, ln, col)
				currentGears = append(currentGears, i)
			}
			if isGearAt(grid, ln+1, col) {
				gears, i = createOrGetIndex(gears, ln+1, col)
				currentGears = append(currentGears, i)
			}

			if char >= '0' && char <= '9' {
				number.WriteRune(char)
				if col < len(line)-1 {
					continue
				}
			}

			if len(currentGears) > 0 && number.Len() > 0 {
				n, err := strconv.Atoi(number.String())
				if err != nil {
					log.Printf("Could not parse part number %s", number.String())
				}
				gears[currentGears[0]].parts = append(gears[currentGears[0]].parts, n)
			}

			number.Reset()
			currentGears = nil

			if isGearAt(grid, ln-1, col) {
				gears, i = createOrGetIndex(gears, ln-1, col)
				currentGears = append(currentGears, i)
			}
			if isGearAt(grid, ln, col) {
				gears, i = createOrGetIndex(gears, ln, col)
				currentGears = append(currentGears, i)
			}
			if isGearAt(grid, ln+1, col) {
				gears, i = createOrGetIndex(gears, ln+1, col)
				currentGears = append(currentGears, i)
			}
		}
	}

	var sum int
	for _, g := range gears {
		if len(g.parts) == 2 {
			sum += g.parts[0] * g.parts[1]
		}
	}

	return sum
}

func isSymbol(char byte) bool {
	return !(char >= '0' && char <= '9' || char == '.')
}

func isSymbolUpDown(grid []string, ln int, col int) bool {
	return ln > 0 && isSymbol(grid[ln-1][col]) || ln < len(grid)-1 && isSymbol(grid[ln+1][col])
}

func isGearAt(grid []string, ln int, col int) bool {
	return ln >= 0 && ln < len(grid) && col >= 0 && col < len(grid[0]) && grid[ln][col] == '*'
}
