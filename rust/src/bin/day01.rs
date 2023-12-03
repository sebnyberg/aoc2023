use std::fs::read_to_string;

pub fn part1(fname: &'static str) -> i32 {
    let mut res: i32 = 0;
    for line in read_to_string(fname).unwrap().lines() {
        let mut first: Option<i32> = None;
        let mut last: Option<i32> = None;
        for ch in line.chars() {
            if !ch.is_numeric() {
                continue;
            }
            if first.is_none() {
                first = Some(ch.to_digit(10).unwrap() as i32);
            }
            last = Some(ch.to_digit(10).unwrap() as i32);
        }
        match (first, last) {
            (Some(first), Some(last)) => res += first * 10 + last,
            _ => (),
        }
    }
    res
}

pub fn part2(fname: &'static str) -> i32 {
    let mut res: i32 = 0;
    fn update(first: &mut Option<i32>, last: &mut Option<i32>, x: i32) {
        if first.is_none() {
            *first = Some(x);
        }
        *last = Some(x);
    }

    for line in read_to_string(fname).unwrap().lines() {
        let mut first: Option<i32> = None;
        let mut last: Option<i32> = None;
        for (i, ch) in line.chars().enumerate() {
            if ch.is_numeric() {
                update(&mut first, &mut last, ch.to_digit(10).unwrap() as i32);
                continue;
            }
            for (idx, s) in [
                "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
            ]
            .iter()
            .enumerate()
            {
                if line[i..].starts_with(s) {
                    update(&mut first, &mut last, (idx + 1) as i32);
                }
            }
        }
        match (first, last) {
            (Some(first), Some(last)) => res += first * 10 + last,
            _ => (),
        }
    }
    res
}

fn main() {
    // println!("{}", part1("src/bin/day01input"));
    println!("{}", part2("src/bin/day01input"));
}
