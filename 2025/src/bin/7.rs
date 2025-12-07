aoc::solution!(7, "Laboratories");

use std::collections::HashMap;

pub fn part_1(input: &str) -> Option<u64> {
    solve(input).map(|(splits, _)| splits)
}

pub fn part_2(input: &str) -> Option<u64> {
    solve(input).map(|(_, timelines)| timelines)
}

fn solve(input: &str) -> Option<(u64, u64)> {
    let mut beams = HashMap::new();
    let mut splits = 0;

    for line in input.lines().step_by(2) {
        for (i, c) in line.chars().enumerate() {
            match c {
                'S' => {
                    beams.insert(i, 1);
                }
                '^' if beams.contains_key(&i) => {
                    splits += 1;
                    let n = beams.remove(&i).unwrap();
                    *beams.entry(i + 1).or_default() += n;
                    *beams.entry(i - 1).or_default() += n;
                }
                _ => {}
            }
        }
    }

    Some((splits, beams.values().sum()))
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = ".......S.......
...............
.......^.......
...............
......^.^......
...............
.....^.^.^.....
...............
....^.^...^....
...............
...^.^...^.^...
...............
..^...^.....^..
...............
.^.^.^.^.^...^.
...............
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(21));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(40));
    }
}
