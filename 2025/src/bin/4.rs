aoc::solution!(4, "Printing Department");

use std::collections::HashSet;

pub fn part_1(input: &str) -> Option<usize> {
    let grid = parse(input);
    can_reach(&grid, |(_, _, cell)| cell == '@').map(|cells| cells.len())
}

pub fn part_2(input: &str) -> Option<usize> {
    let grid = parse(input);
    let mut visited: HashSet<(usize, usize)> = HashSet::new();

    while let Some(cells) = can_reach(&grid, |(x, y, cell)| {
        cell == '@' && !visited.contains(&(x, y))
    }) {
        cells.iter().for_each(|&(x, y)| {
            visited.insert((x, y));
        });
    }

    Some(visited.len())
}

fn can_reach<P: Fn((usize, usize, char)) -> bool>(
    grid: &Grid,
    has_paper: P,
) -> Option<Vec<(usize, usize)>> {
    let result: Vec<(usize, usize)> = grid
        .iter()
        .filter(|&c| has_paper(c))
        .filter(|&(x, y, _)| {
            grid.neighbors(x, y)
                .iter()
                .filter(|&n| has_paper(*n))
                .count()
                < 4
        })
        .map(|(x, y, _)| (x, y))
        .collect();

    (!result.is_empty()).then_some(result)
}

struct Grid(Vec<Vec<char>>);

impl Grid {
    const DIRS: [(isize, isize); 8] = [
        (-1, -1),
        (-1, 0),
        (-1, 1),
        (0, -1),
        (0, 1),
        (1, -1),
        (1, 0),
        (1, 1),
    ];

    fn at(&self, x: usize, y: usize) -> Option<char> {
        self.0.get(y)?.get(x).copied()
    }

    fn neighbors(&self, x: usize, y: usize) -> Vec<(usize, usize, char)> {
        Self::DIRS
            .iter()
            .filter_map(|&(nx, ny)| Some((x.checked_add_signed(nx)?, y.checked_add_signed(ny)?)))
            .filter_map(|(nx, ny)| self.at(nx, ny).map(|n| (nx, ny, n)))
            .collect()
    }

    fn iter(&self) -> impl Iterator<Item = (usize, usize, char)> + '_ {
        self.0
            .iter()
            .enumerate()
            .flat_map(|(y, row)| row.iter().enumerate().map(move |(x, cell)| (x, y, *cell)))
    }
}

fn parse(input: &str) -> Grid {
    Grid(input.lines().map(|line| line.chars().collect()).collect())
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "..@@.@@@@.
@@@.@.@.@@
@@@@@.@.@@
@.@@@@..@.
@@.@@@@.@@
.@@@@@@@.@
.@.@.@.@@@
@.@@@.@@@@
.@@@@@@@@.
@.@.@@@.@.
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(13));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(43));
    }
}
