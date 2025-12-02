aoc::solution!(2, "Gift Shop");

pub fn part_1(input: &str) -> Option<u64> {
    solve(input, |id| {
        if id.len() % 2 != 0 {
            return true;
        }
        check_id(id, vec![2])
    })
}

pub fn part_2(input: &str) -> Option<u64> {
    solve(input, |id| check_id(id, factors(id.len())))
}

fn solve(input: &str, check: impl Fn(&str) -> bool) -> Option<u64> {
    Some(
        parse(input)
            .into_iter()
            .flat_map(|(start, end)| start..=end)
            .filter(|id| !check(&id.to_string()))
            .sum(),
    )
}

fn check_id(id: &str, factors: Vec<usize>) -> bool {
    for f in factors {
        let size = id.len() / f;
        let base = &id[0..size];

        if (1..f).all(|i| &id[i * size..(i + 1) * size] == base) {
            return false;
        }
    }

    true
}

fn factors(x: usize) -> Vec<usize> {
    let mut factors = Vec::new();

    for i in 2..=(x / 2) {
        if x.is_multiple_of(i) {
            factors.push(i);
        }
    }
    if x > 1 {
        factors.push(x);
    }

    factors
}

fn parse(input: &str) -> Vec<(u64, u64)> {
    let first_line = input.lines().next().unwrap();
    first_line
        .split(',')
        .filter_map(|r| r.split_once('-'))
        .map(|(a, b)| (a.parse::<u64>().unwrap(), b.parse::<u64>().unwrap()))
        .collect()
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(1227775554));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(4174379265));
    }
}
