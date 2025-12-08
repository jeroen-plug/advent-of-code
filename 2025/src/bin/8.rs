aoc::solution!(8, "");

use std::collections::HashMap;

#[derive(Debug, Clone, Copy, Default, Eq, Hash, PartialEq)]
struct Point {
    x: u64,
    y: u64,
    z: u64,
}

pub fn part_1(input: &str) -> Option<u64> {
    part_1_inner(input, 1000)
}

fn part_1_inner(input: &str, max_iterations: usize) -> Option<u64> {
    let (mut sizes, _, _) = solve(input, Some(max_iterations))?;
    sizes.sort();
    Some(sizes.pop()? * sizes.pop()? * sizes.pop()?)
}

pub fn part_2(input: &str) -> Option<u64> {
    let (_, a, b) = solve(input, None)?;
    Some(a.x * b.x)
}

fn solve(input: &str, max_iterations: Option<usize>) -> Option<(Vec<u64>, Point, Point)> {
    let junctions = parse(input);
    let max_size = junctions.len() as u64;
    let mut distances = measure_distances(&junctions);

    let mut circuits: HashMap<usize, usize> = HashMap::new();
    let mut next_circuit = 1usize;
    let mut sizes: Vec<u64> = Vec::new();

    let mut a = 0usize;
    let mut b = 0usize;
    for _ in 0..max_iterations.unwrap_or(distances.len()) {
        (a, b, _) = distances.pop()?;

        match (circuits.get(&a).copied(), circuits.get(&b).copied()) {
            (None, None) => {
                circuits.insert(a, next_circuit);
                circuits.insert(b, next_circuit);

                if sizes.len() <= next_circuit {
                    sizes.resize(next_circuit + 1, 0);
                }
                sizes[next_circuit] = 2;

                next_circuit += 1;
            }
            (Some(c), None) | (None, Some(c)) => {
                circuits.insert(a, c);
                circuits.insert(b, c);

                sizes[c] += 1;
                if sizes[c] >= max_size {
                    break;
                }
            }
            (Some(ca), Some(cb)) if ca != cb => {
                let to_move: Vec<usize> = circuits
                    .iter()
                    .filter_map(|(j, c)| if *c == cb { Some(*j) } else { None })
                    .collect();
                for j in to_move {
                    circuits.insert(j, ca);
                }

                sizes[ca] += sizes[cb];
                sizes[cb] = 0;
                if sizes[ca] >= max_size {
                    break;
                }
            }
            _ => {}
        }
    }

    Some((sizes, junctions[a], junctions[b]))
}

fn measure_distances(junctions: &Vec<Point>) -> Vec<(usize, usize, u64)> {
    let mut distances = Vec::new();
    for (i, a) in junctions.iter().enumerate() {
        for (j, b) in junctions[(i + 1)..junctions.len()].iter().enumerate() {
            distances.push((i, i + 1 + j, distance(a, b)));
        }
    }
    distances.sort_by(|a, b| b.2.partial_cmp(&a.2).unwrap());
    distances
}

fn distance(a: &Point, b: &Point) -> u64 {
    let dx = a.x.abs_diff(b.x);
    let dy = a.y.abs_diff(b.y);
    let dz = a.z.abs_diff(b.z);
    dx * dx + dy * dy + dz * dz
}

fn parse(input: &str) -> Vec<Point> {
    input
        .lines()
        .filter_map(|line| {
            let mut numbers = line.splitn(3, ',').filter_map(|d| d.parse::<u64>().ok());
            Some(Point {
                x: numbers.next()?,
                y: numbers.next()?,
                z: numbers.next()?,
            })
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "162,817,812
57,618,57
906,360,560
592,479,940
352,342,300
466,668,158
542,29,236
431,825,988
739,650,466
52,470,668
216,146,977
819,987,18
117,168,530
805,96,715
346,949,466
970,615,88
941,993,340
862,61,35
984,92,344
425,690,689
";

    #[test]
    fn test_part_1() {
        let result = part_1_inner(INPUT, 10);
        assert_eq!(result, Some(40));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(25272));
    }
}
