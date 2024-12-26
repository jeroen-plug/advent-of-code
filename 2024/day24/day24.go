package day24

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2024/input"
)

type GateType int

const (
	GateAnd GateType = iota
	GateOr
	GateXor
)

type Gate struct {
	Input  [2]string
	Output string
	Type   GateType
}

func Solution() (any, any) {
	lines := input.Lines(24)
	return day24a(lines), fmt.Sprintf("%s; Full adder diagram and input for swaps are in ./out24", day24b(lines))
}

func day24a(lines []string) int {
	state, gates := parse(lines)

	unusedGates := slices.Clone(gates)
	active := true
	for active {
		active = false
		var nextGates []Gate
		for _, g := range unusedGates {
			a, aOk := state[g.Input[0]]
			b, bOk := state[g.Input[1]]
			if aOk && bOk {
				active = true
				state[g.Output] = applyGate(a, b, g.Type)
			} else {
				nextGates = append(nextGates, g)
			}
		}
		unusedGates = nextGates
	}

	res := 0
	for connection, value := range state {
		if connection[0] == 'z' {
			offset := input.ParseInt(connection[1:])
			res |= value << offset
		}
	}

	return res
}

func day24b(lines []string) string {
	os.Mkdir("out24", os.ModeDir)

	initial, gates := parse(lines)
	swaps := getSwaps()
	for _, swap := range swaps {
		i0 := slices.IndexFunc(gates, func(g Gate) bool { return g.Output == swap[0] })
		i1 := slices.IndexFunc(gates, func(g Gate) bool { return g.Output == swap[1] })
		gates[i0].Output = swap[1]
		gates[i1].Output = swap[0]
	}

	f, err := os.Create(path.Join("out24", "full_adder.puml"))
	if err != nil {
		log.Fatal("Could not write out24")
	}
	defer f.Close()
	f.WriteString(toPlantUml(initial, gates))

	result := slices.Concat(swaps...)
	slices.Sort(result)
	return strings.Join(result, ",")
}

func applyGate(a, b int, gateType GateType) int {
	result := 0
	switch gateType {
	case GateAnd:
		result = (a & b)
	case GateOr:
		result = (a | b)
	case GateXor:
		result = (a ^ b)
	default:
		log.Fatalf("Unknown gate %d", gateType)
	}
	return result
}

func parse(lines []string) (map[string]int, []Gate) {
	initial := make(map[string]int)
	var gates []Gate

	beforeBreak := true
	for _, l := range lines {
		if l == "" {
			beforeBreak = false
		} else if beforeBreak {
			split := strings.Split(l, ":")
			initial[split[0]] = input.ParseInt(strings.TrimSpace(split[1]))
		} else {
			fields := strings.Fields(l)
			gates = append(gates, Gate{
				Input:  [2]string{fields[0], fields[2]},
				Output: fields[4],
				Type:   toGateType(fields[1]),
			})
		}
	}
	return initial, gates
}

func toGateType(s string) GateType {
	if s == "AND" {
		return GateAnd
	} else if s == "OR" {
		return GateOr
	} else {
		return GateXor
	}
}

func toPlantUml(initial map[string]int, gates []Gate) string {
	var nodes, connecions []string
	nodes = append(nodes, "interface input")

	for connection, value := range initial {
		connecions = append(connecions, fmt.Sprintf("input --> %s : %d", connection, value))
	}
	for _, gate := range gates {
		element := ""
		switch gate.Type {
		case GateAnd:
			element = "action"
		case GateOr:
			element = "process"
		case GateXor:
			element = "hexagon"
		}
		nodes = append(nodes, fmt.Sprintf("%s %s", element, gate.Output))
		connecions = append(connecions, fmt.Sprintf("%s --> %s", gate.Input[0], gate.Output))
		connecions = append(connecions, fmt.Sprintf("%s --> %s", gate.Input[1], gate.Output))
	}

	slices.Sort(connecions)
	return fmt.Sprintf("@startuml\n%s\n\n%s\n@enduml\n", strings.Join(nodes, "\n"), strings.Join(connecions, "\n"))
}

func getSwaps() [][]string {
	file := path.Join("out24", "swaps.csv")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		f, err := os.Create(file)
		if err != nil {
			log.Fatal("Could not create swaps.csv")
		}
		f.Close()
	}

	f, err := os.Open(file)
	if err != nil {
		log.Fatal("Could not read swaps.csv")
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var swaps [][]string

	for scanner.Scan() {
		swaps = append(swaps, strings.Split(scanner.Text(), ","))
	}

	return swaps
}
