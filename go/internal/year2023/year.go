package year2023

import (
	"github.com/kellen-miller/advent-of-code/go/internal"
	"github.com/kellen-miller/advent-of-code/go/internal/year2023/day1"
)

type Year struct{}

func (y *Year) Year() int {
	return 2023
}

func (y *Year) AdventDays() []internal.AdventDay {
	return []internal.AdventDay{
		new(day1.Day),
	}
}
