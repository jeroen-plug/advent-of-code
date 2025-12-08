aoc::solution!(8, "");

#[derive(Debug, Clone, Copy, Eq, Hash, PartialEq)]
struct Point {
    x: u64,
    y: u64,
    z: u64,
}

struct DisjointSet {
    parent: Vec<usize>,
    size: Vec<usize>,
}

impl DisjointSet {
    fn new(n: usize) -> Self {
        DisjointSet {
            parent: (0..n).collect(),
            size: vec![1; n],
        }
    }

    fn find(&mut self, x: usize) -> usize {
        if self.parent[x] != x {
            self.parent[x] = self.find(self.parent[x]);
        }
        self.parent[x]
    }

    fn union(&mut self, a: usize, b: usize) {
        let mut a = self.find(a);
        let mut b = self.find(b);

        if a == b {
            return;
        }

        if self.size[a] < self.size[b] {
            std::mem::swap(&mut a, &mut b);
        }

        self.parent[b] = a;
        self.size[a] += self.size[b];
    }

    fn find_size(&mut self, x: usize) -> usize {
        let x = self.find(x);
        self.size[x]
    }
}

pub fn part_1(input: &str) -> Option<usize> {
    part_1_inner(input, 1000)
}

fn part_1_inner(input: &str, max_iterations: usize) -> Option<usize> {
    let (circuits, _, _) = solve(input, Some(max_iterations))?;
    let mut sizes: Vec<usize> = circuits
        .parent
        .iter()
        .enumerate()
        .filter(|(i, p)| i == *p)
        .map(|(i, _)| circuits.size[i])
        .collect();
    sizes.sort_unstable();
    Some(sizes.pop()? * sizes.pop()? * sizes.pop()?)
}

pub fn part_2(input: &str) -> Option<u64> {
    let (_, a, b) = solve(input, None)?;
    Some(a.x * b.x)
}

fn solve(input: &str, max_iterations: Option<usize>) -> Option<(DisjointSet, Point, Point)> {
    let junctions = parse(input);
    let mut distances = measure_distances(&junctions);
    let mut circuits = DisjointSet::new(junctions.len());

    let mut a = 0usize;
    let mut b = 0usize;
    for _ in 0..max_iterations.unwrap_or(distances.len()) {
        (a, b, _) = distances.pop()?;
        circuits.union(a, b);
        if circuits.find_size(a) >= junctions.len() {
            break;
        }
    }

    Some((circuits, junctions[a], junctions[b]))
}

fn measure_distances(junctions: &[Point]) -> Vec<(usize, usize, u64)> {
    let mut distances = Vec::new();
    for (i, a) in junctions.iter().enumerate() {
        for (j, b) in junctions[(i + 1)..junctions.len()].iter().enumerate() {
            distances.push((i, i + 1 + j, distance(a, b)));
        }
    }
    distances.sort_by_key(|d| std::cmp::Reverse(d.2));
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
