package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

const evenSplitDivisor = 2

func Part1(sc *bufio.Scanner) (string, error) {
	var invalidTotal int

	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if line == "" {
			continue
		}

		lineTotal, err := sumInvalidFromLine(line)
		if err != nil {
			return "", err
		}
		invalidTotal += lineTotal
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.Itoa(invalidTotal), nil
}

func part1ValidateID(id string) bool {
	if id == "" { // empty
		return false
	}

	if id[0] == '0' { // leading zero
		return false
	}

	if len(id)%evenSplitDivisor != 0 { // odd length
		return true
	}

	mid := len(id) / evenSplitDivisor
	return id[:mid] != id[mid:]
}

func sumInvalidFromLine(line string) (int, error) {
	idRanges := strings.Split(line, ",")
	var total int
	for _, idRange := range idRanges {
		idRange = strings.TrimSpace(idRange)
		if idRange == "" {
			continue
		}

		firstID, lastID, err := parseRangeBounds(idRange)
		if err != nil {
			return 0, err
		}

		for id := firstID; id <= lastID; id++ {
			if !part1ValidateID(strconv.Itoa(id)) {
				total += id
			}
		}
	}

	return total, nil
}

func parseRangeBounds(idRange string) (int, int, error) {
	firstIDStr, lastIDStr, ok := strings.Cut(idRange, "-")
	if !ok {
		return 0, 0, fmt.Errorf("could not parse id range: %s", idRange)
	}

	firstID, err := strconv.Atoi(strings.TrimSpace(firstIDStr))
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse id %s: %w", firstIDStr, err)
	}

	lastID, err := strconv.Atoi(strings.TrimSpace(lastIDStr))
	if err != nil {
		return 0, 0, fmt.Errorf("could not parse id %s: %w", lastIDStr, err)
	}

	return firstID, lastID, nil
}
