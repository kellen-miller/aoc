package day8

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/languages/go/pkg/grid"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

func Part2(sc *bufio.Scanner) (string, error) {
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

func createResonantAntinodes(antennaPairs []UniqueAntennaPair, rows int, cols int) *hashset.Set[grid.Point] {
	antinodes := hashset.New[grid.Point]()
	for _, ap := range antennaPairs {
		disX, disY := ap.XYDistance()

		a1Count := addAntinodes(ap.Antenna1, antinodes, rows, cols, disX, disY)
		a2Count := addAntinodes(ap.Antenna2, antinodes, rows, cols, -disX, -disY)

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
	disX int,
	disY int,
) int {
	var (
		aX     = antenna.Position.X
		aY     = antenna.Position.Y
		aCount int
	)
	for {
		aX -= disX
		aY -= disY

		newPoint := grid.Point{X: aX, Y: aY}
		if !newPoint.IsValid(rows, cols) {
			break
		}

		antinodes.Add(newPoint)
		aCount++
	}

	return aCount
}
