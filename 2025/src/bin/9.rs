aoc::solution!(9, "Movie Theater");

#[derive(Debug, Clone, Copy)]
struct Point {
    x: u64,
    y: u64,
}

pub fn part_1(input: &str) -> Option<u64> {
    let boundary = parse(input);
    solve(&boundary, |_| true)
}

pub fn part_2(input: &str) -> Option<u64> {
    let boundary = parse(input);
    solve(&boundary, |rect| !rect_blocked(&boundary, rect))
}

fn solve(boundary: &[Point], filter: impl Fn((Point, Point)) -> bool) -> Option<u64> {
    boundary
        .iter()
        .enumerate()
        .flat_map(|(i, a)| {
            boundary[(i + 1)..]
                .iter()
                .filter(|&b| filter((*a, *b)))
                .map(|b| area(*a, *b))
        })
        .max()
}

fn rect_blocked(boundary: &[Point], rect: (Point, Point)) -> bool {
    boundary.iter().enumerate().any(|(i, a)| {
        let b = boundary[(i + 1) % boundary.len()];
        intersects((rect.0.x, rect.1.x), (a.x, b.x)) && intersects((rect.0.y, rect.1.y), (a.y, b.y))
    })
}

fn intersects(a: (u64, u64), b: (u64, u64)) -> bool {
    let a_start = a.0.min(a.1);
    let a_end = a.0.max(a.1);
    let b_start = b.0.min(b.1);
    let b_end = b.0.max(b.1);

    a_start < b_end && b_start < a_end
}

fn area(a: Point, b: Point) -> u64 {
    (a.x.abs_diff(b.x) + 1) * (a.y.abs_diff(b.y) + 1)
}

fn parse(input: &str) -> Vec<Point> {
    input
        .lines()
        .filter_map(|line| {
            line.split_once(',').and_then(|(x, y)| {
                Some(Point {
                    x: x.parse().ok()?,
                    y: y.parse().ok()?,
                })
            })
        })
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "7,1
11,1
11,7
9,7
9,5
2,5
2,3
7,3
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(50));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(24));
    }
}
