package day3

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

var mulRegex = regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
	var total int
	for sc.Scan() {
		t, err := findTotal(sc.Text())
		if err != nil {
			return "", err
		}

		total += t
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(total), nil
}

func findTotal(s string) (int, error) {
	matches := mulRegex.FindAllStringSubmatch(s, -1)
	if len(matches) == 0 {
		return 0, nil
	}

	var total int
	for _, match := range matches {
		if len(match) != 3 {
			return 0, fmt.Errorf("invalid match: %v", match)
		}

		a, err := strconv.Atoi(match[1])
		if err != nil {
			return 0, err
		}

		b, err := strconv.Atoi(match[2])
		if err != nil {
			return 0, err
		}

		total += a * b
	}

	return total, nil
}
