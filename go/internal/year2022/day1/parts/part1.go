package parts

import (
	"strconv"

	"github.com/kellen-miller/aoc/go/internal/year2022"
	"github.com/kellen-miller/aoc/go/pkg/io"
)

func MostCalories(input string) int {
	if input == "" {
		input = year2022.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		maxCals     int
		currentCals int
	)

	for sc.Scan() {
		cals, err := strconv.Atoi(sc.Text())
		currentCals += cals

		if err != nil { // new line
			if currentCals > maxCals {
				maxCals = currentCals
			}
			currentCals = 0
		}
	}

	if currentCals > maxCals {
		maxCals = currentCals
	}

	return maxCals
}
