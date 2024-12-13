package year2023

import (
	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/internal/year2023/day1"
	"github.com/kellen-miller/aoc/go/internal/year2023/day2"
	"github.com/kellen-miller/aoc/go/internal/year2023/day3"
)

type Year struct{}

func (y *Year) Year() int {
	return 2023
}

func (y *Year) AdventDays() []internal.AdventDay {
	return []internal.AdventDay{
		new(day1.Day),
		new(day2.Day),
		new(day3.Day),
	}
}
