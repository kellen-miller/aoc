package day5

import (
	"bufio"

	"github.com/kellen-miller/aoc/2022/go/day5/parts"
	"github.com/kellen-miller/aoc/languages/go/bridge"
)

func Part1(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day5-part1")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return parts.RearrangeCrates(path), nil
}

func Part2(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day5-part2")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return parts.RearrangeCratesMulti(path), nil
}
