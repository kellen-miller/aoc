package day2

import (
	"bufio"
	"strconv"
	"strings"
)

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
	var input [][]int //nolint:prealloc // We don't know the size of the input
	for sc.Scan() {
		row := strings.Fields(sc.Text())

		var levels []int
		for _, l := range row {
			level, err := strconv.Atoi(l)
			if err != nil {
				return "", err
			}

			levels = append(levels, level)
		}

		input = append(input, levels)
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	var safeLevels int
	for _, level := range input {
		if isLevelSafe(level, true, true) ||
			isLevelSafe(level, true, false) ||
			isLevelSafe(level, false, true) ||
			isLevelSafe(level, false, false) {
			safeLevels++
		}
	}

	return strconv.Itoa(safeLevels), nil
}

// validateLevel tries to validate the sequence in one direction, either forward or backward.
func isLevelSafe(level []int, allowRemovals bool, isForward bool) bool {
	// We'll simulate iteration in reverse by adjusting indices
	// If isForward: i starts from 0 to len(level)-2
	// If !isForward: i starts from len(level)-1 down to 1
	indices := make([]int, len(level))
	if levelLen := len(level); isForward {
		for i := range levelLen {
			indices[i] = i
		}
	} else {
		for i := range levelLen {
			indices[i] = levelLen - 1 - i
		}
	}

	for i := range len(indices) - 1 {
		var (
			currentIndex = indices[i]
			nextIndex    = indices[i+1]
		)

		if diff := level[nextIndex] - level[currentIndex]; diff >= 1 && diff <= 3 {
			continue
		}

		if !allowRemovals {
			return false
		}

		return isLevelSafe(removeIndex(level, currentIndex), false, isForward) ||
			isLevelSafe(removeIndex(level, nextIndex), false, isForward)
	}

	return true
}

func removeIndex(level []int, idx int) []int {
	newLevel := make([]int, 0, len(level)-1)
	newLevel = append(newLevel, level[:idx]...)
	newLevel = append(newLevel, level[idx+1:]...)
	return newLevel
}
