package day23

import (
	"slices"
	"strings"

	"github.com/jeroen-plug/advent-of-code-2024/input"
)

type Network map[string][]string

func Solution() (any, any) {
	lines := input.Lines(23)
	return day23a(lines), day23b(lines)
}

func day23a(lines []string) int {
	network := make(Network)
	sum := 0
	for _, line := range lines {
		split := strings.Split(line, "-")
		left, right := split[0], split[1]

		for _, other := range network[left] {
			if slices.Contains(network[right], other) && (left[0] == 't' || right[0] == 't' || other[0] == 't') {
				sum++
			}
		}

		network[left] = append(network[left], right)
		network[right] = append(network[right], left)
	}
	return sum
}

func day23b(lines []string) string {
	network := parse(lines)

	var largest []string
	for _, clique := range bronKerbosch(network) {
		if len(clique) > len(largest) {
			largest = clique
		}
	}
	slices.Sort(largest)
	return strings.Join(largest, ",")
}

func bronKerbosch(network Network) [][]string {
	var cliques [][]string

	var bk func(R, P, X Set)
	bk = func(R, P, X Set) {
		if len(P) == 0 && len(X) == 0 {
			var clique []string
			for computer := range R {
				clique = append(clique, computer)
			}
			cliques = append(cliques, clique)
		}

		for v := range P {
			N := ToSet(network[v])
			bk(R.Union(Set{v: {}}), P.Intersection(N), X.Intersection(N))
			delete(P, v)
			X[v] = struct{}{}
		}
	}

	P := make(Set)
	for computer := range network {
		P[computer] = struct{}{}
	}
	bk(make(Set), P, make(Set))

	return cliques
}

func parse(lines []string) Network {
	network := make(Network)
	for _, line := range lines {
		split := strings.Split(line, "-")
		left, right := split[0], split[1]
		network[left] = append(network[left], right)
		network[right] = append(network[right], left)
	}
	return network
}
