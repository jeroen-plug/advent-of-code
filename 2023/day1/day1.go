package day1

import (
	"log"
	"os"
	"path"
)

func Solution() (any, any) {
	input, err := os.ReadFile(path.Join("input", "1.txt"))
	if err != nil {
		log.Fatal("Could not read input 1")
	}
	return day1a(string(input)), day1b(string(input))
}

func day1a(input string) int {
	answer := 0

	var first, last int

	for _, b := range input {
		if b == '\n' {
			answer += 10*first + last
			first = 0
			last = 0
			continue
		}

		n := int(b - '0')
		if n > 0 && n < 10 {
			if first == 0 {
				first = n
			}
			last = n
		}
	}

	// In case of a missing trailing newline
	answer += 10*first + last
	return answer
}

func day1b(input string) int {
	answer := 0

	var first, last int

	onDigit := func(n int) {
		if first == 0 {
			first = n
		}
		last = n
	}

	getNext := func(i int, n int) string {
		if i+n >= len(input) {
			return input[i:i]
		}
		return input[i : i+n]
	}

	for i, b := range input {
		if b == '\n' {
			answer += 10*first + last
			first = 0
			last = 0
			continue
		}

		n := int(b - '0')
		if n > 0 && n < 10 {
			onDigit(n)
			continue
		}

		switch getNext(i, 3) {
		case "one":
			onDigit(1)
		case "two":
			onDigit(2)
		case "thr":
			if getNext(i, len("three")) == "three" {
				onDigit(3)
			}
		case "fou":
			if getNext(i, len("four")) == "four" {
				onDigit(4)
			}
		case "fiv":
			if getNext(i, len("five")) == "five" {
				onDigit(5)
			}
		case "six":
			onDigit(6)
		case "sev":
			if getNext(i, len("seven")) == "seven" {
				onDigit(7)
			}
		case "eig":
			if getNext(i, len("eight")) == "eight" {
				onDigit(8)
			}
		case "nin":
			if getNext(i, len("nine")) == "nine" {
				onDigit(9)
			}
		}
	}

	// In case of a missing trailing newline
	answer += 10*first + last
	return answer
}
