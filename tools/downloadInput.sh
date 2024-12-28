#!/bin/bash

TOOLS=$(dirname "$(readlink -f "$0")")

if [ ! -f "$TOOLS/.env" ]; then
  echo "AOC_COOKIE=" > .env
fi

source "$TOOLS/.env"

download() {
    YEAR=$1
    DAY=$2
    if [ -f "input/$DAY.txt" ]; then
        return
    fi
    mkdir -p input
    curl -s --cookie "session=$AOC_COOKIE" "https://adventofcode.com/$YEAR/day/$DAY/input" --output "input/$DAY.txt"
    echo "Downloaded input for day $DAY of $YEAR"
}

YEAR=$(basename "$(pwd)")
if ! [[ "$YEAR" =~ ^2[0-9]{3}$ ]]; then
	echo "Run this script from the root of a year directory."
	exit 1
fi

if [ $# -eq 0 ]; then
    for ((i=1; i<=25; i++)); do
        download $YEAR $i
    done
elif [ $# -eq 1 ]; then
    download $YEAR $1
else
    echo "Usage: $0 <day> Download input for a specific day"
    echo "       $0       Download input for all days"
    exit 1
fi
