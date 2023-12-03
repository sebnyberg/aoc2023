use std::fs::read_to_string;

struct GameRow {
    idx: usize,
    color_counts: Vec<(u32, u32, u32)>,
}

fn parse(row: &str) -> GameRow {
    let parts: Vec<&str> = row.split(" ").collect();
    let mut res = GameRow {
        idx: parts.get(1).unwrap()[..parts[1].len() - 1]
            .parse::<usize>()
            .unwrap(),
        color_counts: Vec::new(),
    };
    let mut count: (u32, u32, u32) = (0, 0, 0);
    for i in (2..parts.len()).step_by(2) {
        let x = parts[i].parse::<u32>().unwrap();
        let color = parts[i + 1];
        if color.starts_with("red") {
            count.0 += x;
        } else if color.starts_with("green") {
            count.1 += x;
        } else {
            count.2 += x;
        }
        if color.ends_with(";") || i + 2 >= parts.len() {
            res.color_counts.push(count);
            count = (0, 0, 0);
        }
    }
    return res;
}

fn part1(fname: &'static str) -> i32 {
    // The grammar is as follows
    // ROW := GAME ":" SET [";" SET]...
    // GAME := "Game" INT
    // SET := BALLCOUNT ["," BALLCOUNT]...
    // BALLCOUNT := INT COLOR
    // INT := [1-9][0-9]+
    // COLOR := ["blue"|"red"|"green"]
    //
    let mut res: i32 = 0;

    'outer: for line in read_to_string(fname).unwrap().lines() {
        if line == "" {
            continue;
        }
        let game = parse(line);
        for set in &game.color_counts {
            if set.0 > 12 || set.1 > 13 || set.2 > 14 {
                continue 'outer;
            }
        }
        res += game.idx as i32;
    }
    res
}

#[allow(dead_code)]
#[allow(unused_variables)]
fn part2(fname: &'static str) -> i32 {
    return 0;
}

fn main() {
    println!("{}", part1("src/bin/day02input"));
    println!("{}", part2("src/bin/day02input"));
}
