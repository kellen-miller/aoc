package day3

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/2022/go/day3/parts"
	"github.com/kellen-miller/aoc/languages/go/bridge"
)

func Part1(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day3-part1")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return strconv.Itoa(parts.SuppliesPriorityTotal(path)), nil
}

func Part2(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day3-part2")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return strconv.Itoa(parts.BadgePriorityTotal(path)), nil
}
