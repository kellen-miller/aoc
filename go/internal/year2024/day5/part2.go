package day5

import (
	"bufio"
)

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
	for sc.Scan() {
		// Do something with sc.Text()
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return "", nil
}
