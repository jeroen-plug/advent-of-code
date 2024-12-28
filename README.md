# Advent of Code Solutions
This repository contains my solutions for [Advent of Code](https://adventofcode.com/), organized by year.

| Year | Language | Stars | CI/CD Status |
|------|----------|-------|--------------|
| [2024](https://adventofcode.com/2024) | Go     | ![50 Stars](https://img.shields.io/badge/Stars-50-gold?logo=adventofcode) | [![2024](https://github.com/jeroen-plug/advent-of-code/actions/workflows/2024.yml/badge.svg)](https://github.com/jeroen-plug/advent-of-code/actions/workflows/2024.yml) |
| [2017](https://adventofcode.com/2017) | Python | ![02 Stars](https://img.shields.io/badge/Stars-02-gold?logo=adventofcode) | [![2017](https://github.com/jeroen-plug/advent-of-code/actions/workflows/2017.yml/badge.svg)](https://github.com/jeroen-plug/advent-of-code/actions/workflows/2017.yml) |

## Running the Solutions
To run the solutions, navigate to the specific year's folder and run `make`.
Other targets are available, such as `make test` to run all tests, or `make 1` to run a specific day.

Input files for each day are not included in this repository.
To automatically download the input files, you need to set the `AOC_COOKIE` variable in `tools/.env`, then run `make input`.
Alternatively, you can manually download the input files and place them in the `input` folder.
