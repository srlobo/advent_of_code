use std::{env, i32};
use std::fs;

fn main() {
    let args: Vec<String> = env::args().collect();
    let f_name = &args[1];

    let mut total = 0;
    for line in fs::read_to_string(f_name).unwrap().lines() {
        let mut first_digit = '\0';
        let mut last_digit = '\0';
        for c in line.chars(){
            if c.is_numeric() && first_digit == '\0' {
                first_digit = c;
            }
            if c.is_numeric() {
                last_digit = c;
            }
        }
        dbg!(first_digit);
        dbg!(last_digit);
        let concat = format!("{}{}",first_digit, last_digit);
        total += concat.parse::<i32>().unwrap();

    }
    dbg!(total);
}
