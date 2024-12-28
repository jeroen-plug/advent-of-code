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

# Python specific

if [ ! -f day$DAY/solution.py ]; then
    mkdir -p day$DAY

    cat > day$DAY/solution.py <<EOF
import inputs


def part1():
    data = inputs.string($DAY)
    # TODO: Implement part1


def part2():
    data = inputs.string($DAY)
    # TODO: Implement part2
EOF
fi

if [ ! -f day$DAY/test_solution.py ]; then
    mkdir -p day$DAY

    cat > day$DAY/test_solution.py <<EOF
import pytest
from unittest.mock import patch
from . import solution

example = """"""


class TestDay$DAY:
    def test_part1(self, data, expect):
        with patch("inputs.lines", return_value=example.split("\n")):
            assert solution.part1() == expect

    def test_part2(self, data, expect):
        with patch("inputs.lines", return_value=example.split("\n")):
            assert solution.part2() == expect
EOF
fi

touch day$DAY/__init__.py
