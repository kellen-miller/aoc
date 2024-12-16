package day8

import "github.com/kellen-miller/aoc/go/pkg/grid"

func parseLine(
	line string,
	row int,
	antennas map[rune][]Antenna,
	antennaPairs []UniqueAntennaPair,
) []UniqueAntennaPair {
	var newAntennaPairs []UniqueAntennaPair
	for i, c := range line {
		if c == '.' {
			continue
		}

		antenna := Antenna{
			Position: grid.Point{X: i, Y: row},
			Value:    c,
		}

		antennas[c] = append(antennas[c], antenna)

		avs, ok := antennas[c]
		if !ok {
			panic("antenna slice not found for c: " + string(c))
		}

		for _, a := range avs {
			if a.Equals(&antenna) {
				continue
			}

			newAp := UniqueAntennaPair{
				Antenna1: a,
				Antenna2: antenna,
			}

			var found bool
			for _, ap := range antennaPairs {
				if ap.Equals(&newAp) {
					found = true
					break
				}
			}

			if !found {
				newAntennaPairs = append(newAntennaPairs, newAp)
			}
		}
	}

	return newAntennaPairs
}
