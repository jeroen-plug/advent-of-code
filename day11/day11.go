package day11

import (
	"fmt"
	"math"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Day11() {
	data := strings.TrimSuffix(input.String(11), "\n")

	fmt.Printf("day 11a: %d\n", day11(data, 25))
	fmt.Printf("day 11b: %d\n", day11(data, 75))
}

func day11(data string, blinks int) int {
	stones := make(map[int]int) // how many of each number
	for _, s := range strings.Fields(data) {
		stones[input.ParseInt(s)]++
	}

	for range blinks {
		newStones := make(map[int]int)
		for s, n := range stones {
			if n == 0 {
				continue
			} else if s == 0 {
				newStones[1] += n
			} else if digits := int(math.Log10(float64(s)) + 1); digits%2 == 0 {
				split := int(math.Pow10(digits / 2))
				newStones[s/split] += n
				newStones[s%split] += n
			} else {
				newStones[s*2024] += n
			}
		}
		stones = newStones
	}

	sum := 0
	for _, n := range stones {
		sum += n
	}

	return sum
}
