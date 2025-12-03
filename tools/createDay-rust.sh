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

# Rust specific

if [ ! -f src/bin/$DAY.rs ]; then
    mkdir -p src/bin

    cat > src/bin/$DAY.rs <<EOF
aoc::solution!($DAY, "");

pub fn part_1(input: &str) -> Option<u64> {
    None
}

pub fn part_2(input: &str) -> Option<u64> {
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, None);
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, None);
    }
}
EOF
fi
