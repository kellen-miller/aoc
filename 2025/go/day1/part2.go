package day1

import (
	"bufio"
	"fmt"
	"strconv"
)

func Part2(sc *bufio.Scanner) (string, error) {
	var (
		pos      = dialStartPosition
		prevPos  int
		password int
	)

	for sc.Scan() {
		line := sc.Text()

		dir := line[:1]
		change, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("could not parse change %q: %w", line[1:], err)
		}

		password += change / dialPositions
		change %= dialPositions
		prevPos = pos

		var rollover bool
		if dir == rightRotation {
			pos += change
			if pos >= dialPositions {
				pos -= dialPositions
				rollover = prevPos != unlockedPosition
			}
		} else {
			pos -= change
			if pos < 0 {
				pos += dialPositions
				rollover = prevPos != unlockedPosition
			}
		}

		if pos == unlockedPosition || rollover {
			password++
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.Itoa(password), nil
}
