package day1

import (
	"bufio"
	"fmt"
	"strconv"
)

const (
	dialPositions     = 100
	dialStartPosition = 50
	unlockedPosition  = 0
	rightRotation     = "R"
)

func Part1(sc *bufio.Scanner) (string, error) {
	var (
		pos      = dialStartPosition
		password int
	)

	for sc.Scan() {
		line := sc.Text()

		dir := line[:1]
		change, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("could not parse change %q: %w", line[1:], err)
		}

		delta := change
		if dir != rightRotation {
			delta = -change
		}

		pos = (pos + delta + dialPositions) % dialPositions

		if pos == unlockedPosition {
			password++
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.Itoa(password), nil
}
