package day1

import (
	"bufio"
	"errors"
	"fmt"
	"math"
	"regexp"
	"slices"
	"strconv"
)

var locationIDsInputRegex = regexp.MustCompile(`^(\d+)\s+(\d+)$`)

func Part1(sc *bufio.Scanner) (string, error) {
	var (
		leftLocationsIDs  []int
		rightLocationsIDs []int
	)
	for sc.Scan() {
		line := sc.Text()

		matches := locationIDsInputRegex.FindStringSubmatch(line)
		if matches == nil || len(matches) != 3 {
			return "", errors.New("invalid input: expected two integers separated by whitespace")
		}

		leftLocationID, err := strconv.Atoi(matches[1])
		if err != nil {
			return "", fmt.Errorf("parse left location ID %q: %w", matches[1], err)
		}

		rightLocationID, err := strconv.Atoi(matches[2])
		if err != nil {
			return "", fmt.Errorf("parse right location ID %q: %w", matches[2], err)
		}

		leftLocationsIDs = append(leftLocationsIDs, leftLocationID)
		rightLocationsIDs = append(rightLocationsIDs, rightLocationID)
	}

	slices.Sort(leftLocationsIDs)
	slices.Sort(rightLocationsIDs)

	if len(leftLocationsIDs) != len(rightLocationsIDs) {
		return "", errors.New("invalid input: left and right location IDs must have the same length")
	}

	var distance float64
	for i := range leftLocationsIDs {
		distance += math.Abs(float64(leftLocationsIDs[i] - rightLocationsIDs[i]))
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("scan location IDs: %w", err)
	}

	return strconv.FormatFloat(distance, 'f', -1, 64), nil
}
