#!/bin/bash

TOOLS=$(dirname "$(readlink -f "$0")")

if [ ! -f "$TOOLS/.env" ]; then
  echo "AOC_COOKIE=" > .env
  echo "AOC_EMAIL=" >> .env
fi

source "$TOOLS/.env"

download() {
    YEAR=$1
    DAY=$2
    if [ -f "input/$DAY.txt" ]; then
        return
    fi
    mkdir -p input
    URL="$(git remote get-url origin | sed -E 's|^git@github.com:(.+)\.git$|https://github.com/\1|')"
    EMAIL="${AOC_EMAIL:-$(git config user.email)}"
    curl -s --cookie "session=$AOC_COOKIE" --user-agent "downloadInput.sh (+$URL; mailto:$EMAIL)" "https://adventofcode.com/$YEAR/day/$DAY/input" --output "input/$DAY.txt"
    echo "Downloaded input for day $DAY of $YEAR"
}

YEAR=$(basename "$(pwd)")
if ! [[ "$YEAR" =~ ^2[0-9]{3}$ ]]; then
	echo "Run this script from the root of a year directory."
	exit 1
fi

if [ $# -eq 0 ]; then
    days=25
    [ "$YEAR" -ge 2025 ] && days=12
    for ((i=1; i<=$days; i++)); do
        download $YEAR $i
    done
elif [ $# -eq 1 ]; then
    download $YEAR $1
else
    echo "Usage: $0 <day> Download input for a specific day"
    echo "       $0       Download input for all days"
    exit 1
fi
