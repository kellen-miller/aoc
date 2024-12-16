package year2024

import (
	"github.com/kellen-miller/aoc/go/internal/advent"
	"github.com/kellen-miller/aoc/go/internal/year2024/day1"
	"github.com/kellen-miller/aoc/go/internal/year2024/day2"
	"github.com/kellen-miller/aoc/go/internal/year2024/day3"
	"github.com/kellen-miller/aoc/go/internal/year2024/day4"
	"github.com/kellen-miller/aoc/go/internal/year2024/day5"
	"github.com/kellen-miller/aoc/go/internal/year2024/day6"
	"github.com/kellen-miller/aoc/go/internal/year2024/day7"
)

type Year struct{}

func (y *Year) AdventYear() int {
	return 2024
}

func (y *Year) AdventDays() []advent.Day {
	return []advent.Day{
		new(day1.Day),
		new(day2.Day),
		new(day3.Day),
		new(day4.Day),
		new(day5.Day),
		new(day6.Day),
		new(day7.Day),
	}
}
