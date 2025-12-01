package day0

import (
	"bufio"
	"fmt"
)

func Part1(sc *bufio.Scanner) (string, error) {
	for sc.Scan() {
		// Do something with sc.Text()
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return "", nil
}
