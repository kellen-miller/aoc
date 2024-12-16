package day8

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/go/pkg/grid"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
	var (
		antennas     = make(map[rune][]Antenna)
		antennaPairs []UniqueAntennaPair
		rows         int
		cols         int
	)
	for sc.Scan() {
		antennaPairs = append(antennaPairs, parseLine(sc.Text(), rows, antennas, antennaPairs)...)
		rows++
		if cols == 0 {
			cols = len(sc.Text())
		}
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(createAntinodes(antennaPairs, rows, cols).Size()), nil
}

func createAntinodes(antennaPairs []UniqueAntennaPair, rows, cols int) *hashset.Set[grid.Point] {
	antinodes := hashset.New[grid.Point]()
	for _, ap := range antennaPairs {
		xDis, yDis := ap.Antenna2.Position.XYDistance(ap.Antenna1.Position)

		antinode1 := grid.Point{
			X: ap.Antenna1.Position.X - xDis,
			Y: ap.Antenna1.Position.Y - yDis,
		}

		if antinode1.X >= 0 && antinode1.Y >= 0 && antinode1.X < cols && antinode1.Y < rows {
			antinodes.Add(antinode1)
		}

		antinode2 := grid.Point{
			X: ap.Antenna2.Position.X + xDis,
			Y: ap.Antenna2.Position.Y + yDis,
		}

		if antinode2.X >= 0 && antinode2.Y >= 0 && antinode2.X < cols && antinode2.Y < rows {
			antinodes.Add(antinode2)
		}
	}

	return antinodes
}
