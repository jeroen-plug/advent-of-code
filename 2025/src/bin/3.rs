aoc::solution!(3, "Lobby");

pub fn part_1(input: &str) -> Option<u64> {
    Some(input.lines().map(|b| largest_joltage(b, 2)).sum())
}

pub fn part_2(input: &str) -> Option<u64> {
    Some(input.lines().map(|b| largest_joltage(b, 12)).sum())
}

fn largest_joltage(bank: &str, digits: usize) -> u64 {
    let mut joltage = vec!['0'; digits];

    for (i, battery) in bank.chars().enumerate() {
        let better = (0..digits).find(|&d| battery > joltage[d] && i + digits - d <= bank.len());
        if let Some(d) = better {
            joltage[d] = battery;
            joltage[d + 1..].fill('0');
        }
    }

    joltage
        .iter()
        .fold(0, |acc, &b| acc * 10 + b.to_digit(10).unwrap() as u64)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "987654321111111
811111111111119
234234234234278
818181911112111
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(357));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(3121910778619));
    }
}
