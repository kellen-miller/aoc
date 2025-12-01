package day4

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/2022/go/day4/parts"
	"github.com/kellen-miller/aoc/languages/go/bridge"
)

func Part1(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day4-part1")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return strconv.Itoa(parts.RedundantCleanup(path)), nil
}

func Part2(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day4-part2")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return strconv.Itoa(parts.OverlappingSections(path)), nil
}
