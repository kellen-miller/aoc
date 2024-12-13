package day2

import (
	"bufio"
	"math"
	"strconv"
	"strings"
)

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
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

		var (
			isSafe    = true
			ascending bool
		)
		for i, level := range levels {
			if i == 0 {
				continue
			}

			prevLevel, err := strconv.Atoi(levels[i-1])
			if err != nil {
				return "", err
			}

			currLevel, err := strconv.Atoi(level)
			if err != nil {
				return "", err
			}

			if i == 1 {
				ascending = currLevel > prevLevel
			}

			if diff := math.Abs(float64(currLevel - prevLevel)); (diff < 1 || 3 < diff) ||
				(ascending && prevLevel > currLevel) ||
				(!ascending && currLevel > prevLevel) {
				isSafe = false
				break
			}
		}

		if isSafe {
			safe++
		}
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(safe), nil
}
