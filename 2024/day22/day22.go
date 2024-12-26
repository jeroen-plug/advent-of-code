package day22

import (
	"maps"
	"slices"
	"sync"

	"github.com/jeroen-plug/advent-of-code/2024/input"
)

func Solution() (any, any) {
	lines := input.Lines(22)
	return day22a(lines), day22b(lines, 32)
}

func day22a(lines []string) int {
	sum := 0
	var secret MonkeySecret
	for _, seed := range lines {
		secret.SetSeed(input.ParseInt(seed))
		for range 2000 {
			secret.Next()
		}
		sum += int(secret)
	}
	return sum
}

func day22b(lines []string, parallel int) int {
	results := make(chan map[Pattern]int)
	var wg sync.WaitGroup

	for chunk := range slices.Chunk(lines, len(lines)/parallel) {
		wg.Add(1)
		go func(chunk []string) {
			defer wg.Done()
			results <- findPatternProfits(chunk)
		}(chunk)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	profits := make(map[Pattern]int)
	for result := range results {
		for pattern, profit := range result {
			profits[pattern] += profit
		}
	}

	bestProfit := 0
	for profit := range maps.Values(profits) {
		if profit > bestProfit {
			bestProfit = profit
		}
	}

	return bestProfit
}

func findPatternProfits(seeds []string) map[Pattern]int {
	patterns := make(map[Pattern]int)
	for _, seed := range seeds {
		buyer := newBuyer(input.ParseInt(seed))
		var seen []Pattern
		for i := 4; i < len(buyer.Deltas); i++ {
			current := Pattern(buyer.Deltas[i-4 : i])
			if slices.Contains(seen, current) {
				continue
			}
			patterns[current] += buyer.Prices[i]
			seen = append(seen, current)
		}
	}
	return patterns
}

type MonkeySecret int

func (s *MonkeySecret) SetSeed(seed int) {
	*s = MonkeySecret(seed)
}

func (s *MonkeySecret) Next() {
	const bm = 16777216 - 1
	(*s) ^= ((*s) << 6) & bm  // * 64
	(*s) ^= ((*s) >> 5) & bm  // / 32
	(*s) ^= ((*s) << 11) & bm // * 2048
}

func (s MonkeySecret) CurrentPrice() int {
	return int(s) % 10
}

type Pattern [4]int

type Buyer struct {
	Prices []int
	Deltas []int
}

func newBuyer(seed int) Buyer {
	var secret MonkeySecret
	secret.SetSeed(seed)

	buyer := Buyer{Prices: []int{secret.CurrentPrice()}}
	for i := range 2000 {
		secret.Next()
		currentPrice := secret.CurrentPrice()
		buyer.Prices = append(buyer.Prices, currentPrice)
		buyer.Deltas = append(buyer.Deltas, buyer.Prices[i+1]-buyer.Prices[i])
	}
	return buyer
}
