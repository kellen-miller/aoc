package day3

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var mulIfRegex = regexp.MustCompile(`mul\((\d+),(\d+)\)|(do\(\))|(don't\(\))`)

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
	var (
		enabled = true
		total   int
	)
	for sc.Scan() {
		matches := mulIfRegex.FindAllStringSubmatch(sc.Text(), -1)
		if len(matches) == 0 {
			return "", fmt.Errorf("no matches found")
		}

		for _, match := range matches {
			if len(match) != 5 {
				return "", fmt.Errorf("invalid match: %v", match)
			}

			if match[3] == "do()" {
				enabled = true
				continue
			} else if match[4] == "don't()" {
				enabled = false
				continue
			}

			if !enabled {
				continue
			}

			a, err := strconv.Atoi(match[1])
			if err != nil {
				return "", err
			}

			b, err := strconv.Atoi(match[2])
			if err != nil {
				return "", err
			}

			total += a * b
		}
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(total), nil
}

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
