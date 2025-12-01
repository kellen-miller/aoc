package day6

import (
	"bufio"
	"fmt"

	"github.com/kellen-miller/aoc/2022/go/day6/parts"
	"github.com/kellen-miller/aoc/languages/go/bridge"
)

func Part1(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day6-part1")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return fmt.Sprint(parts.StartOfPacket(path)), nil
}

func Part2(sc *bufio.Scanner) (string, error) {
	path, cleanup, err := bridge.MaterializeScanner(sc, "2022-day6-part2")
	if err != nil {
		return "", err
	}
	defer cleanup()

	return fmt.Sprint(parts.StartOfMessage(path)), nil
}
