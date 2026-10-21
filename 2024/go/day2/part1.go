package day2

import (
	"bufio"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func Part1(sc *bufio.Scanner) (string, error) {
	var safe int
	for sc.Scan() {
		line := sc.Text()

		levels := strings.Split(line, " ")

		if len(levels) == 0 {
			continue
		}

		if len(levels) == 1 {
			safe++
			continue
		}

		isSafe, err := checkLevels(true, levels)
		if err != nil {
			return "", fmt.Errorf("check report %q: %w", line, err)
		}

		if isSafe {
			safe++
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("scan reports: %w", err)
	}

	return strconv.Itoa(safe), nil
}

func checkLevels(isSafe bool, levels []string) (bool, error) {
	var ascending bool
	for i, level := range levels {
		if i == 0 {
			continue
		}

		prevLevel, err := strconv.Atoi(levels[i-1])
		if err != nil {
			return false, fmt.Errorf("parse previous level %q: %w", levels[i-1], err)
		}

		currLevel, err := strconv.Atoi(level)
		if err != nil {
			return false, fmt.Errorf("parse current level %q: %w", level, err)
		}

		if i == 1 {
			ascending = currLevel > prevLevel
		}

		if diff := math.Abs(float64(currLevel - prevLevel)); (diff < 1 || 3 < diff) ||
			(ascending && prevLevel > currLevel) ||
			(!ascending && currLevel > prevLevel) {
			return false, nil
		}
	}

	return isSafe, nil
}
