package day3

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mulIfRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)

func Part2(sc *bufio.Scanner) (string, error) {
	var (
		enabled = true
		total   int
	)
	for sc.Scan() {
		matches := mulIfRegex.FindAllStringSubmatch(sc.Text(), -1)
		if len(matches) == 0 {
			return "", errors.New("no matches found")
		}

		for _, match := range matches {
			var (
				mul int
				err error
			)
			mul, enabled, err = checkMatch(match, enabled)
			if err != nil {
				return "", err
			}

			total += mul
		}
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(total), nil
}

func checkMatch(match []string, enabled bool) (int, bool, error) {
	if len(match) != 5 { //nolint:mnd // 5 is the expected length
		return 0, enabled, fmt.Errorf("invalid match: %v", match)
	}

	if match[3] == "do()" {
		return 0, true, nil
	} else if match[4] == "don't()" {
		return 0, false, nil
	}

	if !enabled {
		return 0, enabled, nil
	}

	a, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, enabled, err
	}

	b, err := strconv.Atoi(match[2])
	if err != nil {
		return 0, enabled, err
	}

	return a * b, enabled, nil
}

//nolint:unparam // The function signature must match the one in the internal package.
func part2SplitCutRegex(sc *bufio.Scanner) (string, error) {
	var program strings.Builder
	for sc.Scan() {
		program.WriteString(sc.Text())
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	var (
		doStarts = strings.Split(program.String(), "do()")
		dos      strings.Builder
	)
	for _, doStart := range doStarts {
		do, _, _ := strings.Cut(doStart, "don't()")
		dos.WriteString(do)
	}

	total, err := findTotal(dos.String())
	if err != nil {
		return "", err
	}

	return strconv.Itoa(total), nil
}
