package day1

import (
	"fmt"
	"log"
	"math"
	"slices"
	"strconv"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

func Day1() {
	lines := input.Lines(1)

	fmt.Printf("day 1a: %d\n", day1a(lines))
	fmt.Printf("day 1b: %d\n", day1b(lines))
}

func day1a(lines []string) int {
	list1, list2 := parse(lines)

	var distance int
	for i := 0; i < len(list1); i++ {
		distance += int(math.Abs(float64(list1[i] - list2[i])))
	}

	return distance
}

func day1b(lines []string) int {
	list1, list2 := parse(lines)

	var (
		similarity int
		occurances int
		i          int
	)
	for _, n := range list1 {
		if i == 0 || list2[i-1] != n {
			occurances = 0
		}

		for ; i < len(list2) && list2[i] <= n; i++ {
			if list2[i] == n {
				occurances++
			}
		}

		similarity += n * occurances
	}

	return similarity
}

func parse(lines []string) ([]int, []int) {
	var (
		list1 []int
		list2 []int
	)

	for _, l := range lines {
		numbers := strings.Fields(l)

		n1, err := strconv.Atoi(numbers[0])
		if err != nil {
			log.Fatalf("Could not parse number %s", numbers[0])
		}
		n2, err := strconv.Atoi(numbers[1])
		if err != nil {
			log.Fatalf("Could not parse number %s", numbers[0])
		}

		list1 = append(list1, n1)
		list2 = append(list2, n2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	return list1, list2
}
