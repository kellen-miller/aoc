package day2

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func Part2(sc *bufio.Scanner) (string, error) {
	var invalidTotal int

	for sc.Scan() {
		idRanges := strings.Split(sc.Text(), ",")
		if len(idRanges) == 0 {
			break
		}

		for _, idRange := range idRanges {
			firstIDStr, lastIDStr, ok := strings.Cut(idRange, "-")
			if !ok {
				return "", fmt.Errorf("could not parse id range: %s", idRange)
			}

			firstID, err := strconv.Atoi(firstIDStr)
			if err != nil {
				return "", fmt.Errorf("could not parse id %d: %w", firstID, err)
			}

			lastID, err := strconv.Atoi(lastIDStr)
			if err != nil {
				return "", fmt.Errorf("could not parse id %d: %w", lastID, err)
			}

			for i := firstID; i <= lastID; i++ {
				if !part2ValidateID(strconv.Itoa(i)) {
					invalidTotal += i
				}
			}
		}
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.Itoa(invalidTotal), nil
}

func part2ValidateID(id string) bool {
	if len(id) == 0 { // empty
		return false
	}

	if id[0] == '0' { // leading zero
		return false
	}

	return false
}
