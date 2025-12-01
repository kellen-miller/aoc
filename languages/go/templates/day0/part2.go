package day0

import (
	"bufio"
	"fmt"
)

func Part2(sc *bufio.Scanner) (string, error) {
	for sc.Scan() {
		// Do something with sc.Text()
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return "", nil
}
