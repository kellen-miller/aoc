package day1

import (
	"bufio"
	"fmt"
	"strconv"
)

func Part1(sc *bufio.Scanner) (string, error) {
	var (
		pos      = 50
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
		if dir != "R" {
			delta = -change
		}

		pos = (pos + delta + 100) % 100

		if pos == 0 {
			password++
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.Itoa(password), nil
}
