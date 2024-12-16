package internal

import (
	"github.com/kellen-miller/aoc/go/internal/advent"
	"github.com/kellen-miller/aoc/go/internal/year2023"
	"github.com/kellen-miller/aoc/go/internal/year2024"
)

func Years() []advent.Year {
	return []advent.Year{
		new(year2023.Year),
		new(year2024.Year),
	}
}
