# Advent of Code (aoc)

## Description

This repository contains my solutions to the Advent of Code challenges.

## How To Run
Solutions are organized as `/<year>/<language>/...` so that each language can keep year-specific code together. Shared helpers for a language live under `languages/<language>`.

Available flags when running via Task:

- `--year` - Year of the challenge, default is `all`
- `--day` - Day of the challenge, default is `all`
- `--part` - Part of the challenge, default is `all`
- `--lang` - Language identifier, default is `all`

```bash
# Run all challenges for every language/year combination
task run

# Run a specific problem for Go
task run -- --year 2024 --day 1 --part 1 --lang go

# Run all days/parts for Go in 2023
task run -- --year 2023 --lang go
```

## Languages

- [go](languages/go/)

> ℹ️ Adding a new Go day? Run `go generate ./languages/go/cmd/advent` (or `task run` which does it for you) so the runner picks it up automatically.
