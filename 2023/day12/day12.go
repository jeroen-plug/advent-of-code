package day12

import (
	"fmt"
	"log"
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code/2023/input"
)

type Record struct {
	Row    string
	Groups []int
}

var (
	Operational rune = '.'
	Damaged          = '#'
	Unknown          = '?'
)

func Solution() (any, any) {
	lines := input.Lines(12)
	return day12a(lines), day12b(lines)
}

func day12a(lines []string) int {
	records := parse(lines)

	sum := 0
	for _, r := range records {
		sum += countPermutations(r)

		if countPermutations(r) != countPermutations2(r) {
			fmt.Println(r)
		}
	}

	return sum
}

func day12b(lines []string) int {
	records := parse(lines)

	sum := 0
	for _, r := range records {
		sum += countPermutations(unfold(r))
	}

	return sum
}

func unfold(r Record) Record {
	return Record{
		fmt.Sprintf("%s?%s?%s?%s?%s", r.Row, r.Row, r.Row, r.Row, r.Row),
		slices.Concat(r.Groups, r.Groups, r.Groups, r.Groups, r.Groups),
	}
}

type State struct {
	Group        int
	Amount       int
	Permutations int
}

func countPermutations(record Record) int {
	states := []State{{0, 0, 1}}
	for _, condition := range record.Row + string(Operational) {
		// fmt.Printf("\nCondition: %s\n", string(condition))
		var newStates []State
		for _, state := range states {
			// fmt.Printf("  In: %d\n", state)
			switch condition {
			case Operational:
				if s, valid := onOperational(record, state); valid {
					newStates = addOrAppend(newStates, s)
				}
			case Damaged:
				if s, valid := onDamaged(record, state); valid {
					newStates = addOrAppend(newStates, s)
				}
			case Unknown:
				if s, valid := onOperational(record, state); valid {
					newStates = addOrAppend(newStates, s)
				}
				if s, valid := onDamaged(record, state); valid {
					newStates = addOrAppend(newStates, s)
				}
			}
		}
		states = newStates
		// fmt.Println(states)
	}

	result := states[slices.IndexFunc(states, func(s State) bool { return s.Amount == 0 && s.Group == len(record.Groups) })]
	return result.Permutations
}

func addOrAppend(states []State, state State) []State {
	if i := slices.IndexFunc(states, func(s State) bool { return s.Amount == state.Amount && s.Group == state.Group }); i >= 0 {
		states[i].Permutations += state.Permutations
		// fmt.Printf("    Increment: %d\n", states[i])
	} else {
		states = append(states, state)
		// fmt.Printf("    Out: %d\n", state)
	}
	return states
}

func onOperational(r Record, s State) (State, bool) {
	if s.Amount != 0 && s.Amount != r.Groups[s.Group] {
		return State{}, false
	}
	if s.Amount != 0 {
		s.Group++
		s.Amount = 0
	}
	return s, true
}

func onDamaged(r Record, s State) (State, bool) {
	if s.Amount++; s.Group >= len(r.Groups) || s.Amount > r.Groups[s.Group] {
		return State{}, false
	} else {
		return s, true
	}
}

func parse(lines []string) []Record {
	var records []Record
	for _, l := range lines {
		fields := strings.Fields(l)

		var groups []int
		for _, g := range strings.Split(fields[1], ",") {
			groups = append(groups, input.ParseInt(g))
		}

		records = append(records, Record{fields[0], groups})
	}
	return records
}

// old slow version

func countPermutations2(r Record) int {
	totalDamaged := 0
	for _, i := range r.Groups {
		totalDamaged += i
	}

	missingDamaged := totalDamaged - strings.Count(r.Row, string(Damaged))
	missingOperational := strings.Count(r.Row, string(Unknown)) - missingDamaged
	permutations := permute(missingOperational, missingDamaged, "")

	sum := 0
	for _, p := range permutations {
		if isValid(applyPermutation(r.Row, p), r.Groups) {
			sum++
		}
	}

	return sum
}

func isValid(row string, groups []int) bool {
	currentGroup := 0
	currentSize := 0
	for _, condition := range row {
		switch condition {
		case Operational:
			if currentSize != 0 {
				currentGroup++
				currentSize = 0
			}
		case Damaged:
			if currentSize++; currentGroup >= len(groups) || currentSize > groups[currentGroup] {
				return false
			}
		default:
			log.Fatal("Invalid row:", row)
		}
	}
	return true
}

func applyPermutation(row string, permutation string) string {
	p := strings.NewReader(permutation)
	var newRow strings.Builder
	for _, condition := range row {
		if condition == Unknown {
			condition, _, _ = p.ReadRune()
		}
		newRow.WriteRune(condition)
	}
	return newRow.String()
}

func permute(operational, damaged int, s string) []string {
	if operational == 0 {
		return []string{s + strings.Repeat(string(Damaged), damaged)}
	} else if damaged == 0 {
		return []string{s + strings.Repeat(string(Operational), operational)}
	} else {
		return slices.Concat(
			permute(operational-1, damaged, s+string(Operational)),
			permute(operational, damaged-1, s+string(Damaged)),
		)
	}
}
