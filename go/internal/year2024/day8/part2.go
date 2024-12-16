package day8

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/go/pkg/grid"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
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

	return strconv.Itoa(createResonantAntinodes(antennaPairs, rows, cols).Size()), nil
}

func createResonantAntinodes(antennaPairs []UniqueAntennaPair, rows, cols int) *hashset.Set[grid.Point] {
	var antinodes = hashset.New[grid.Point]()
	for _, ap := range antennaPairs {
		xDis, yDis := ap.Antenna2.Position.XYDistance(ap.Antenna1.Position)

		a1Count := addAntinodes(ap.Antenna1, antinodes, rows, cols, xDis, yDis)
		a2Count := addAntinodes(ap.Antenna2, antinodes, rows, cols, -xDis, -yDis)

		if a1Count == 1 {
			antinodes.Add(ap.Antenna1.Position)
		}

		if a2Count == 1 {
			antinodes.Add(ap.Antenna2.Position)
		}

		if a1Count >= 2 || a2Count >= 2 {
			antinodes.Add(ap.Antenna1.Position)
			antinodes.Add(ap.Antenna2.Position)
		}
	}

	return antinodes
}

func addAntinodes(
	antenna Antenna,
	antinodes *hashset.Set[grid.Point],
	rows int,
	cols int,
	xDis int,
	yDis int,
) int {
	var (
		aX     = antenna.Position.X
		aY     = antenna.Position.Y
		aCount int
	)
	for {
		aX -= xDis
		aY -= yDis

		newPoint := grid.Point{X: aX, Y: aY}
		if !newPoint.IsValid(rows, cols) {
			break
		}

		antinodes.Add(newPoint)
		aCount++
	}

	return aCount
}
