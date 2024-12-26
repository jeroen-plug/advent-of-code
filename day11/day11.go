package day11

import (
	"math"
	"math/big"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Solution() (any, any) {
	data := strings.TrimSuffix(input.String(11), "\n")
	return day11(data, 25), day11(data, 75)
	// fmt.Printf("day 11 extra (100000): %d\n", day11("0 1 10 99 999", 100000))
}

func day11(data string, blinks int) *big.Int {
	stones := make(map[int]*big.Int) // how many of each number
	for _, s := range strings.Fields(data) {
		i := input.ParseInt(s)
		stones[i] = Add(stones[i], big.NewInt(1))
	}

	for range blinks {
		newStones := make(map[int]*big.Int)
		for s, n := range stones {
			if n.Cmp(big.NewInt(0)) == 0 {
				continue
			} else if s == 0 {
				newStones[1] = Add(newStones[1], n)
			} else if digits := int(math.Log10(float64(s)) + 1); digits%2 == 0 {
				split := int(math.Pow10(digits / 2))
				iLeft := s / split
				iRight := s % split
				newStones[iLeft] = Add(newStones[iLeft], n)
				newStones[iRight] = Add(newStones[iRight], n)
			} else {
				i := s * 2024
				newStones[i] = Add(newStones[i], n)
			}
		}
		stones = newStones
	}

	sum := big.NewInt(0)
	for _, n := range stones {
		sum.Add(sum, n)
	}

	return sum
}

func Add(a, b *big.Int) *big.Int {
	if a == nil {
		a = big.NewInt(0)
	}
	return a.Add(a, b)
}
