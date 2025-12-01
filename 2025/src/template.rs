use std::fmt::Display;
use std::{env, fs};

#[must_use]
pub fn read_input(day: u8) -> String {
    let cwd = env::current_dir().unwrap();
    let filepath = cwd.join("input").join(format!("{day}.txt"));
    let f = fs::read_to_string(filepath);
    f.expect("could not open input file")
}

pub fn print_result<T: Display>(result: Option<T>, part: u8) {
    match result {
        Some(result) => println!("    Part {part}: {result}"),
        None => println!("    Part {part}: X"),
    }
}

#[macro_export]
macro_rules! solution {
    ($day:expr, $title:expr) => {
        $crate::solution!(@impl $day, $title, [part_1, 1] [part_2, 2]);
    };
    ($day:expr, $title:expr, 1) => {
        $crate::solution!(@impl $day, $title, [part_1, 1]);
    };
    ($day:expr, $title:expr, 2) => {
        $crate::solution!(@impl $day, $title, [part_2, 2]);
    };

    (@impl $day:expr, $title:expr, $( [$func:expr, $part:expr] )*) => {
        fn main() {
            let input = $crate::template::read_input($day);
            println!("\x1b[1;33m2025 Day {}: \x1b[22m{}\x1b[0m", $day, $title);
            $( $crate::template::print_result($func(&input), $part); )*
        }
    };
}
