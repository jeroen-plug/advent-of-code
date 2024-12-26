package day17

import (
	"slices"
	"strconv"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2024/input"
)

func Solution() (any, any) {
	lines := input.Lines(17)
	return day17a(lines), day17b(lines)
}

func day17a(lines []string) string {
	program, registers := parse(lines)
	return asString(computer(program, registers))
}

func day17b(lines []string) int {
	program, registers := parse(lines)
	input := make([]int, len(program))
	input[len(input)-1] = 1

	for i := len(program) - 1; i >= 0; {
		for ; input[i] <= 7; input[i]++ {
			registers[REG_A] = asRegister(input)
			if computer(program, registers)[i] == program[i] {
				break
			}
		}
		if input[i] > 7 {
			input[i] = 0
			i++
			input[i]++
		} else {
			i--
		}
	}

	return asRegister(input)
}

func parse(lines []string) ([]int, [3]int) {
	var program []int
	var registers [3]int

	registers[REG_A] = input.ParseInt(strings.TrimPrefix(lines[0], "Register A: "))
	registers[REG_B] = input.ParseInt(strings.TrimPrefix(lines[1], "Register B: "))
	registers[REG_C] = input.ParseInt(strings.TrimPrefix(lines[2], "Register C: "))

	for _, n := range strings.Split(strings.TrimPrefix(lines[4], "Program: "), ",") {
		program = append(program, input.ParseInt(n))
	}

	return program, registers
}

func asString(numbers []int) string {
	var s []string
	for _, n := range numbers {
		s = append(s, strconv.Itoa(n))
	}
	return strings.Join(s, ",")
}

func asRegister(numbers []int) int {
	reg := 0
	for _, n := range slices.Backward(numbers) {
		reg <<= 3
		reg += n
	}
	return reg
}
