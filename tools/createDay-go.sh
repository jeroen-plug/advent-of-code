#!/bin/bash

if [ $# -ne 1 ]; then
    echo "Usage: $0 <day>"
    exit 1
fi

DAY=$1
YEAR=$(basename "$(pwd)")
if ! [[ "$YEAR" =~ ^2[0-9]{3}$ ]]; then
	echo "Run this script from the root of a year directory."
	exit 1
fi

TOOLS=$(dirname "$(readlink -f "$0")")
"$TOOLS/downloadInput.sh" $DAY

# Go specific

if [ ! -f day$DAY/day$DAY.go ]; then
    mkdir -p day$DAY

    cat > day$DAY/day$DAY.go <<EOF
package day$DAY

import (
	"github.com/jeroen-plug/advent-of-code/$YEAR/input"
)

func Solution() (any, any) {
	lines := input.Lines($DAY)
	return Part1(lines), Part2(lines)
}

func Part1(lines []string) int {
    // TODO: Implement Part1
    return 0
}

func Part2(lines []string) int {
    // TODO: Implement Part2
    return 0
}
EOF
fi

if [ ! -f day$DAY/day${DAY}_test.go ]; then
    mkdir -p day$DAY

    cat > day$DAY/day${DAY}_test.go <<EOF
package day$DAY

import (
	"strings"
	"testing"
)

const example = \`\`

func TestPart1(t *testing.T) {
	want := 0
	res := Part1(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("Part1() = %d, want %d", res, want)
	}
}

func TestPart2(t *testing.T) {
	want := 0
	res := Part2(strings.Split(example, "\n"))

	if res != want {
		t.Fatalf("Part2() = %d, want %d", res, want)
	}
}
EOF
fi

dayEntry="$DAY: {\"\", day$DAY.Solution},"
if ! grep -q "$dayEntry" main.go; then
    lastLine=$(grep -E '^[[:space:]]*[0-9]+:.*Solution\},[[:space:]]*$' main.go | tail -n 1)
    sed -i "s|$lastLine|$lastLine\n\t\t$dayEntry|" main.go
fi

importLine="\"github.com/jeroen-plug/advent-of-code/$YEAR/day$DAY\""
if ! grep -q "$importLine" main.go; then
    lastLine=$(grep -E '^[[:space:]]*\"github.com/jeroen-plug/advent-of-code/[0-9]+/day[0-9]+\"$' main.go | tail -n 1)
    sed -i "s|$lastLine|$lastLine\n\t$importLine|" main.go
fi

go fmt main.go
