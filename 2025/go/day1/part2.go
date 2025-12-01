package day1

import (
	"bufio"
	"fmt"
	"strconv"
)

func Part2(sc *bufio.Scanner) (string, error) {
	var (
		pos      = 50
		prevPos  = 50
		password int
	)

	for sc.Scan() {
		line := sc.Text()

		dir := line[:1]
		change, err := strconv.Atoi(line[1:])
		if err != nil {
			return "", fmt.Errorf("could not parse change %q: %w", line[1:], err)
		}

		password += change / 100
		change %= 100
		prevPos = pos

		var rollover bool
		if dir == "R" {
			pos += change
			if pos > 99 {
				pos -= 100
				rollover = prevPos != 0
			}
		} else {
			pos -= change
			if pos < 0 {
				pos += 100
				rollover = prevPos != 0
			}
		}

		if pos == 0 || rollover {
			password++
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.Itoa(password), nil
}
