aoc::solution!(5, "Cafeteria");

pub fn part_1(input: &str) -> Option<usize> {
    let (ranges, ingredients) = parse(input);

    Some(
        ingredients
            .iter()
            .filter(|&i| ranges.iter().any(|(start, end)| i >= start && i <= end))
            .count(),
    )
}

pub fn part_2(input: &str) -> Option<u64> {
    let (ranges, _) = parse(input);

    let mut merged = Vec::new();
    for (start, end) in ranges {
        if let Some((_, last_end)) = merged.last_mut() {
            if start <= *last_end + 1 {
                *last_end = end.max(*last_end);
                continue;
            }
        }
        merged.push((start, end));
    }

    Some(merged.iter().map(|&(start, end)| end - start + 1).sum())
}

fn parse(input: &str) -> (Vec<(u64, u64)>, Vec<u64>) {
    let (ranges, ingredients) = input.split_once("\n\n").unwrap();

    let mut ranges: Vec<(u64, u64)> = ranges
        .lines()
        .filter_map(|line| {
            let (a, b) = line.split_once('-')?;
            Some((a.parse().ok()?, b.parse().ok()?))
        })
        .collect();
    ranges.sort();

    let ingredients = ingredients
        .lines()
        .filter_map(|line| line.parse().ok())
        .collect();

    (ranges, ingredients)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "3-5
10-14
16-20
12-18
13-14

1
5
8
11
17
32
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(3));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(14));
    }
}
