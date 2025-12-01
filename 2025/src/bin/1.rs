aoc::solution!(1, "Secret Entrance");

const DIAL_SIZE: i32 = 100;
const DIAL_START: i32 = 50;

pub fn part_1(input: &str) -> Option<u32> {
    let mut dial = DIAL_START;
    let mut password = 0;

    for turn in parse(input) {
        dial += turn;
        dial %= DIAL_SIZE;
        if dial == 0 {
            password += 1;
        }
    }

    Some(password)
}

pub fn part_2(input: &str) -> Option<u32> {
    let mut dial = DIAL_START;
    let mut password = 0;

    for turn in parse(input) {
        // The dial does not pass '0' if it starts there
        if dial == 0 && turn < 0 {
            password -= 1;
        }

        dial += turn;
        password += dial.div_euclid(DIAL_SIZE).unsigned_abs();
        dial = dial.rem_euclid(DIAL_SIZE);

        if dial == 0 && turn < 0 {
            password += 1;
        }
    }

    Some(password)
}

fn parse(input: &str) -> impl Iterator<Item = i32> {
    input.lines().filter_map(|line| {
        if line.is_empty() {
            return None;
        }
        let (dir, rest) = line.split_at(1);
        let n: i32 = rest.parse().ok()?;
        Some(if dir == "L" { -n } else { n })
    })
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT: &str = "L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT);
        assert_eq!(result, Some(3));
    }

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT);
        assert_eq!(result, Some(6));
    }

    #[test]
    fn test_part_2_extra() {
        let cases = [("L50,L110", 2), ("L50,R10,L10", 2), ("L49,L101", 2)];

        for (input, expected) in cases {
            let result = part_2(&input.replace(',', "\n"));
            assert_eq!(result, Some(expected), "Failed on input: {input:?}");
        }
    }
}
