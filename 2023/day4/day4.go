package day4

import (
	"bufio"
	"log"
	"math"
	"os"
	"path"
	"slices"
	"strconv"
	"strings"
)

type day4card struct {
	card int
	win  []int
	have []int
	mult int
}

func Solution() (any, any) {
	f, err := os.Open(path.Join("input", "4.txt"))
	if err != nil {
		log.Fatal("Could not read input 4")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return day4a(lines), day4b(lines)
}

func day4a(lines []string) int {
	cards := day4parse(lines)
	var points int

	for _, c := range cards {
		var wins int

		for _, have := range c.have {
			if slices.Contains(c.win, have) {
				wins++
			}
		}

		// int(2^-1) == int(0.5) == 0
		points += int(math.Pow(2, float64(wins-1)))
	}

	return points
}

func day4b(lines []string) int {
	cards := day4parse(lines)
	total := len(cards)

	for i, c := range cards {
		var wins int

		for _, have := range c.have {
			if slices.Contains(c.win, have) {
				wins++
				if i+wins < len(cards) {
					cards[i+wins].mult += c.mult
					total += c.mult
				}
			}
		}
	}

	return total
}

func day4parse(lines []string) []day4card {
	var cards []day4card

	for _, l := range lines {
		s := strings.SplitN(l, ":", 2)
		numbers := strings.Trim(s[1], " ")
		c, _ := strings.CutPrefix(s[0], "Card")
		card, err := strconv.Atoi(strings.Trim(c, " "))
		if err != nil {
			log.Printf("Can't parse number '%s'", strings.Trim(s[0], " "))
		}

		s = strings.SplitN(numbers, "|", 2)
		var (
			win  []int
			have []int
		)

		for _, w := range strings.Split(s[0], " ") {
			if len(strings.Trim(w, " ")) == 0 {
				continue
			}
			n, err := strconv.Atoi(w)
			if err != nil {
				log.Printf("Can't parse number '%s'", w)
			}
			win = append(win, n)
		}

		for _, h := range strings.Split(s[1], " ") {
			if len(strings.Trim(h, " ")) == 0 {
				continue
			}
			n, err := strconv.Atoi(h)
			if err != nil {
				log.Printf("Can't parse number '%s'", h)
			}
			have = append(have, n)
		}

		cards = append(cards, day4card{card, win, have, 1})
	}

	return cards
}
