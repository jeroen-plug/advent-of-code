aoc::solution!(10, "Factory");

use good_lp::{Expression, Solution, SolverModel, constraint, default_solver, variables};
use itertools::Itertools;

#[derive(Debug, Default)]
struct Machine {
    indicator: u16,
    buttons: Vec<Vec<usize>>,
    joltage: Vec<u32>,
}

pub fn part_1(input: &str) -> Option<usize> {
    let machines = parse(input);
    let mut sum = 0;

    for machine in &machines {
        'search: for k in 1..=machine.buttons.len() {
            for combo in machine.buttons.iter().combinations(k) {
                let indicator = combo
                    .into_iter()
                    .flatten()
                    .fold(0u16, |acc, &i| acc ^ (1 << i));

                if indicator == machine.indicator {
                    sum += k;
                    break 'search;
                }
            }
        }
    }

    Some(sum)
}

pub fn part_2(input: &str) -> Option<usize> {
    let machines = parse(input);
    let mut sum = 0;

    for machine in &machines {
        variables! {vars: 0 <= x[machine.buttons.len()] (integer); }
        let objective: Expression = x.iter().sum();

        let constraints = machine.joltage.iter().enumerate().map(|(j, &target)| {
            let presses: Expression = machine
                .buttons
                .iter()
                .enumerate()
                .filter_map(|(b, button)| {
                    if button.contains(&j) {
                        Some(x[b])
                    } else {
                        None
                    }
                })
                .sum();
            constraint!(presses == target)
        });

        sum += vars
            .minimise(&objective)
            .using(default_solver)
            .with_all(constraints)
            .solve()
            .ok()?
            .eval(objective)
            .round() as usize;
    }

    Some(sum)
}

fn parse(input: &str) -> Vec<Machine> {
    input
        .lines()
        .map(|line| {
            let mut machine = Machine::default();
            line.split_whitespace()
                .for_each(|token| parse_token(&mut machine, token));
            machine
        })
        .collect()
}

fn parse_token(machine: &mut Machine, token: &str) {
    match token.chars().next() {
        Some('[') => {
            machine.indicator = token
                .trim_matches(['[', ']'])
                .chars()
                .enumerate()
                .fold(0, |acc, (i, c)| acc | (((c == '#') as u16) << i));
        }
        Some('(') => machine.buttons.push(
            token
                .trim_matches(['(', ')'])
                .split(',')
                .map(|b| b.parse().unwrap())
                .collect(),
        ),
        Some('{') => {
            machine.joltage = token
                .trim_matches(['{', '}'])
                .split(',')
                .map(|j| j.parse().unwrap())
                .collect();
        }
        _ => {}
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}
[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
[.###.#] (0,1,2,3,4) (0,3,4) (0,1,2,4,5) (1,2) {10,11,11,5,10,5}
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(7));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(33));
    }
}
