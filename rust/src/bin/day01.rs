use std::fs;
use std::fs::File;
use std::io::{self, BufRead};

fn main() {
    // List cwd
    let cwd = std::env::current_dir().unwrap();
    println!("cwd: {:?}", cwd);

    // Open the file
    let file = File::open("src/bin/day01input").expect("Could not open file");

    // let reader = io::BufReader::new(file);

    println!("Hello, world1");
}
