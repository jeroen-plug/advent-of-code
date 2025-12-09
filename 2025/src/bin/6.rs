aoc::solution!(6, "Trash Compactor");

pub fn part_1(input: &str) -> Option<u64> {
    let mut homework: Vec<Vec<&str>> = input
        .lines()
        .map(|line| line.split_whitespace().collect())
        .collect();
    let operators = homework.pop()?;

    let mut result = 0;
    for (i, op) in operators.iter().enumerate() {
        let numbers = homework
            .iter()
            .filter_map(|line| line.get(i)?.parse::<u64>().ok());
        result += match *op {
            "+" => numbers.sum::<u64>(),
            "*" => numbers.product::<u64>(),
            _ => 0,
        };
    }

    Some(result)
}

pub fn part_2(input: &str) -> Option<u64> {
    const MARGIN: usize = 2;
    let homework: Vec<&str> = input.lines().collect();

    let mut numbers: Vec<u64> = Vec::new();
    let mut result = 0;

    for col in (0..homework.first()?.len() + MARGIN).rev() {
        let mut number = 0;
        let mut op: Option<char> = None;

        for line in &homework {
            match line.as_bytes().get(col).map(|&b| b as char) {
                Some(n) if n.is_numeric() => number = 10 * number + n.to_digit(10)? as u64,
                Some(o) if o.is_ascii_punctuation() => op = Some(o),
                _ => {}
            }
        }
        numbers.push(number);

        if let Some(op) = op {
            result += match op {
                '+' => numbers.iter().sum::<u64>(),
                '*' => numbers.iter().filter(|&&n| n > 0).product::<u64>(),
                _ => 0,
            };
            numbers.clear();
        }
    }

    Some(result)
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "123 328  51 64
 45 64  387 23
  6 98  215 314
*   +   *   +
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(4277556));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(3263827));
    }
}
