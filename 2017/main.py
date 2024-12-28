#!/usr/bin/env python3

import argparse
import importlib


def main():
    parser = argparse.ArgumentParser(prog="AdventOfCode2017")
    parser.add_argument("day", help="Which day of AoC 2017 to run", type=int)
    args = parser.parse_args()

    try:
        module = importlib.import_module(f"day{args.day}.solution")

        print(f"Day {args.day}:")
        if hasattr(module, "part1") and callable(module.part1):
            print(f"  Part 1: {module.part1()}")
        else:
            print(f"  Day {args.day} does not have a callable 'part1' function.")

        if hasattr(module, "part2") and callable(module.part2):
            print(f"  Part 2: {module.part2()}")
        else:
            print(f"  Day {args.day} does not have a callable 'part2' function.")
    except ModuleNotFoundError:
        print(f"Day {args.day} not found. Make sure day{args.day}/solution.py exists.")


if __name__ == "__main__":
    main()
