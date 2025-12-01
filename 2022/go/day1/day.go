package day1

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/2022/go/day1/parts"
	"github.com/kellen-miller/aoc/languages/go/bridge"
)

func Part1(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day1-part1")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return strconv.Itoa(parts.MostCalories(path)), nil
}

func Part2(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day1-part2")
	if err != nil {
		return "", err
	}
	defer cleanup()

	_, sum := parts.Top3Calories(path)
	return strconv.Itoa(sum), nil
}
