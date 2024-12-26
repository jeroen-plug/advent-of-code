package day17

import (
	"log"
)

type Opcode int

const (
	OP_ADV Opcode = iota
	OP_BXL
	OP_BST
	OP_JNZ
	OP_BXC
	OP_OUT
	OP_BDV
	OP_CDV
)

type Register int

const (
	REG_A Register = iota
	REG_B
	REG_C
)

func computer(program []int, registers [3]int) []int {
	var out []int
	for pc := 0; pc < len(program); pc += 2 {
		switch Opcode(program[pc]) {
		case OP_ADV:
			registers[REG_A] = registers[REG_A] >> comboOperand(registers, program[pc+1])
		case OP_BXL:
			registers[REG_B] ^= program[pc+1]
		case OP_BST:
			registers[REG_B] = comboOperand(registers, program[pc+1]) & 7 // keep lowest 3 bits
		case OP_JNZ:
			if registers[REG_A] != 0 {
				littleTrampoline(&pc, program[pc+1])
			}
		case OP_BXC:
			registers[REG_B] ^= registers[REG_C]
		case OP_OUT:
			out = append(out, comboOperand(registers, program[pc+1])&7) // keep lowest 3 bits
		case OP_BDV:
			registers[REG_B] = registers[REG_A] >> comboOperand(registers, program[pc+1])
		case OP_CDV:
			registers[REG_C] = registers[REG_A] >> comboOperand(registers, program[pc+1])
		default:
			log.Fatalf("Invalid opcode: %d at %d", program[pc], pc)
		}
	}
	return out
}

func comboOperand(registers [3]int, operand int) int {
	if operand < 4 {
		return operand
	} else if operand < 7 {
		return registers[operand-4]
	} else {
		log.Fatalf("Invalid combo operand: %d", operand)
		return 0
	}
}

func littleTrampoline(pc *int, operand int) {
	// -2 to offset for the increment
	*pc = operand - 2
}
