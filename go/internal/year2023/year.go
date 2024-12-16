package year2023

import (
	"github.com/kellen-miller/aoc/go/internal/advent"
	"github.com/kellen-miller/aoc/go/internal/year2023/day1"
	"github.com/kellen-miller/aoc/go/internal/year2023/day2"
	"github.com/kellen-miller/aoc/go/internal/year2023/day3"
)

type Year struct{}

func (y *Year) AdventYear() int {
	return 2023
}

func (y *Year) AdventDays() []advent.Day {
	return []advent.Day{
		new(day1.Day),
		new(day2.Day),
		new(day3.Day),
	}
}
