package day3

import (
	"strings"
	"unicode"

	"github.com/jeroen-plug/advent-of-code/2024/input"
)

type Instruction struct {
	name string
	args []int
}

func Solution() (any, any) {
	data := input.String(3)
	return day3a(data), day3b(data)
}

func day3a(data string) int {
	var result int

	for _, i := range parse(data) {
		if strings.HasSuffix(i.name, "mul") && len(i.args) == 2 {
			result += i.args[0] * i.args[1]
		}
	}

	return result
}

func day3b(data string) int {
	var result int
	enabled := true

	for _, i := range parse(data) {
		if strings.HasSuffix(i.name, "mul") && len(i.args) == 2 && enabled {
			result += i.args[0] * i.args[1]
		} else if strings.HasSuffix(i.name, "do") && len(i.args) == 0 {
			enabled = true
		} else if strings.HasSuffix(i.name, "don't") && len(i.args) == 0 {
			enabled = false
		}
	}

	return result
}

type ParseState int

const (
	Searching ParseState = iota
	ParseFunction
	ParseArg
)

func parse(data string) []Instruction {
	var (
		state  ParseState
		buffer strings.Builder

		function string
		args     []int

		instructions []Instruction
	)

	for _, c := range data {
		switch state {
		case Searching:
			if unicode.IsLetter(c) || c == '\'' {
				buffer.WriteRune(c)
				state = ParseFunction
			}

		case ParseFunction:
			if unicode.IsLetter(c) || c == '\'' {
				buffer.WriteRune(c)
			} else if c == '(' {
				function = buffer.String()
				buffer.Reset()
				state = ParseArg
			} else {
				buffer.Reset()
				state = Searching
			}

		case ParseArg:
			if unicode.IsDigit(c) {
				buffer.WriteRune(c)
			} else if c == ',' || c == ')' {
				if buffer.Len() > 0 {
					args = append(args, input.ParseInt(buffer.String()))
					buffer.Reset()
				}
				if c == ')' {
					instructions = append(instructions, Instruction{function, args})
					args = []int{}
					state = Searching
				}
			} else {
				buffer.Reset()
				args = []int{}
				state = Searching
				if unicode.IsLetter(c) {
					buffer.WriteRune(c)
					state = ParseFunction
				}
			}
		}
	}

	return instructions
}
