package year2024

import (
	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/internal/year2024/day1"
	"github.com/kellen-miller/aoc/go/internal/year2024/day2"
	"github.com/kellen-miller/aoc/go/internal/year2024/day3"
	"github.com/kellen-miller/aoc/go/internal/year2024/day4"
	"github.com/kellen-miller/aoc/go/internal/year2024/day5"
)

type Year struct{}

func (y *Year) Year() int {
	return 2024
}

func (y *Year) AdventDays() []internal.AdventDay {
	return []internal.AdventDay{
		new(day1.Day),
		new(day2.Day),
		new(day3.Day),
		new(day4.Day),
		new(day5.Day),
	}
}
