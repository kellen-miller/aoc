package day2

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/2022/go/day2/parts"
	"github.com/kellen-miller/aoc/languages/go/bridge"
)

func Part1(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day2-part1")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return strconv.Itoa(parts.TotalScore(path)), nil
}

func Part2(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day2-part2")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return strconv.Itoa(parts.SetRoundResult(path)), nil
}
