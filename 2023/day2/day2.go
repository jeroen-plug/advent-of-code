package day2

import (
	"bufio"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

type day2set struct {
	red   int
	green int
	blue  int
}

type day2game struct {
	id      int
	subsets []day2set
}

func Solution() (any, any) {
	f, err := os.Open(path.Join("input", "2.txt"))
	if err != nil {
		log.Fatal("Could not read input 2")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var resA, resB int

	for scanner.Scan() {
		g := day2parse(scanner.Text())
		if day2possible(g, day2set{12, 13, 14}) {
			resA += g.id
		}
		resB += day2power(g)
	}

	return resA, resB
}

func day2max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func day2power(game day2game) int {
	fewest := day2set{}
	for _, s := range game.subsets {
		fewest.red = day2max(fewest.red, s.red)
		fewest.green = day2max(fewest.green, s.green)
		fewest.blue = day2max(fewest.blue, s.blue)
	}
	return fewest.red * fewest.green * fewest.blue
}

func day2possible(game day2game, set day2set) bool {
	for _, s := range game.subsets {
		if s.red > set.red || s.green > set.green || s.blue > set.blue {
			return false
		}
	}
	return true
}

func day2parse(line string) day2game {
	game := strings.SplitN(line, ":", 2)
	ids, found := strings.CutPrefix(game[0], "Game ")
	if !found {
		log.Fatalf("Could not find game id in %s", game[0])
	}
	id, err := strconv.Atoi(ids)
	if err != nil {
		log.Fatalf("Could not parse game id %s", ids)
	}

	var subsets []day2set
	for _, s := range strings.Split(game[1], ";") {
		subset := day2set{}
		for _, color := range strings.Split(s, ",") {
			value := strings.Split(strings.Trim(color, " "), " ")
			switch value[1] {
			case "red":
				subset.red, err = strconv.Atoi(value[0])
			case "green":
				subset.green, err = strconv.Atoi(value[0])
			case "blue":
				subset.blue, err = strconv.Atoi(value[0])
			}
			if err != nil {
				log.Fatalf("Could not parse game subset %s", s)
			}
		}
		subsets = append(subsets, subset)
	}

	return day2game{id, subsets}
}
