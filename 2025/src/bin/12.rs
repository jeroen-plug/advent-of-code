aoc::solution!(12, "Christmas Tree Farm");

const SHAPES: usize = 6;
const THRESHOLD: usize = 85;

pub fn part_1(input: &str) -> Option<usize> {
    let mut chunks = input.split("\n\n");

    let shapes: Vec<_> = chunks
        .by_ref()
        .take(SHAPES)
        .map(|shape| shape.chars().filter(|&c| c == '#').count())
        .collect();

    Some(
        chunks
            .next()?
            .lines()
            .filter_map(|region| {
                let (size, quantities) = region.split_once(':')?;
                let (w, h) = size.split_once('x')?;
                let area = w.parse::<usize>().ok()? * h.parse::<usize>().ok()?;
                let required = quantities
                    .split_whitespace()
                    .zip(&shapes)
                    .try_fold(0, |acc, (n, shape)| {
                        Some(acc + (shape * n.parse::<usize>().ok()?))
                    })?;
                (required * 100 <= THRESHOLD * area).then_some(())
            })
            .count(),
    )
}

pub fn part_2(_: &str) -> Option<usize> {
    None
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "0:
###
##.
##.

1:
###
##.
.##

2:
.##
###
##.

3:
##.
###
##.

4:
###
#..
###

5:
###
.#.
###

4x4: 0 0 0 0 2 0
12x5: 1 0 1 0 2 2
12x5: 1 0 1 0 3 2
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(2));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, None);
    }
}
