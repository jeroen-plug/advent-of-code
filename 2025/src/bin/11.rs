aoc::solution!(11, "Reactor");

use std::collections::HashMap;

pub fn part_1(input: &str) -> Option<u64> {
    let tree = parse(input);
    let mut memo = HashMap::new();
    Some(dfs(&tree, "you", true, true, &mut memo))
}

pub fn part_2(input: &str) -> Option<u64> {
    let tree = parse(input);
    let mut memo = HashMap::new();
    Some(dfs(&tree, "svr", false, false, &mut memo))
}

fn dfs<'a>(
    tree: &HashMap<&str, Vec<&'a str>>,
    device: &'a str,
    dac: bool,
    fft: bool,
    memo: &mut HashMap<(&'a str, bool, bool), u64>,
) -> u64 {
    if let Some(&cached) = memo.get(&(device, dac, fft)) {
        return cached;
    }

    let (dac, fft) = match device {
        "dac" => (true, fft),
        "fft" => (dac, true),
        _ => (dac, fft),
    };

    let result = match (dac && fft && device == "out", tree.get(device)) {
        (true, _) => 1,
        (_, Some(d)) => d.iter().map(|out| dfs(tree, out, dac, fft, memo)).sum(),
        _ => 0,
    };

    memo.insert((device, dac, fft), result);
    result
}

fn parse(input: &str) -> HashMap<&str, Vec<&str>> {
    let mut tree = HashMap::new();
    for (device, outputs) in input.lines().filter_map(|line| line.split_once(':')) {
        tree.insert(device, outputs.split_whitespace().collect());
    }
    tree
}

#[cfg(test)]
mod tests {
    use super::*;

    const INPUT1: &str = "aaa: you hhh
you: bbb ccc
bbb: ddd eee
ccc: ddd eee fff
ddd: ggg
eee: out
fff: out
ggg: out
hhh: ccc fff iii
iii: out
";

    #[test]
    fn test_part_1() {
        let result = part_1(INPUT1);
        assert_eq!(result, Some(5));
    }

    const INPUT2: &str = "svr: aaa bbb
aaa: fft
fft: ccc
bbb: tty
tty: ccc
ccc: ddd eee
ddd: hub
hub: fff
eee: dac
dac: fff
fff: ggg hhh
ggg: out
hhh: out
";

    #[test]
    fn test_part_2() {
        let result = part_2(INPUT2);
        assert_eq!(result, Some(2));
    }
}
