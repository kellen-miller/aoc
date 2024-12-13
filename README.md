# Advent of Code (aoc)

## Description

This repository contains my solutions to the Advent of Code challenges.

## How To Run
Available flags:

- `--year` - Year of the challenge, default is `all`
- `--day` - Day of the challenge, default is `all`
- `--part` - Part of the challenge, default is `all`

```bash
# Run all challenges
task run

# Run specific problem
task run -- --year 2024 --day 1 --part 1

# Run all problems for a specific language
task run:go

# Run specific problem for a specific language
task run:go -- --year 2024 --day 1 --part 1
```

## Languages

- [go](go/)