# Advent of Code (aoc)

## Description

This repository contains my solutions to the Advent of Code challenges.

## How To Run
Solutions are organized as `/<year>/<language>/...` so that each language can keep year-specific code together. Shared helpers for a language live under `languages/<language>`.

Available flags when running via Task (pass as `KEY=VALUE`, for example `task run YEAR=2024 DAY=1`):

- `--year` - Year of the challenge, default is `all`
- `--day` - Day of the challenge, default is `all`
- `--part` - Part of the challenge, default is `all`
- `--lang` - Language identifier, default is `all` (set with `LANGUAGE=...` to avoid clashing with system `LANG`)

```bash
# Run all challenges for every language/year combination
task run

# Run a specific problem for Go
task run YEAR=2024 DAY=1 PART=1 LANGUAGE=go

# Run all days/parts for Go in 2023
task run YEAR=2023 LANGUAGE=go
```

## Languages

- [go](languages/go/)

### Scaffolding Days

Language templates live under `templates/<language>` and are rendered with Go's `text/template`. Use the helper task to hydrate them:

```bash
# Create 2024/day12 in Go (LANGUAGE defaults to go)
task new-day YEAR=2024 DAY=12 LANGUAGE=go
```

When scaffolding Go days the task automatically formats the generated code and refreshes the Go runner (`go generate ./languages/go/cmd/advent`), so new days are runnable immediately.
