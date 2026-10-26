package day5

import (
	"bufio"
	"fmt"
	"strconv"
)

func Part1(sc *bufio.Scanner) (string, error) {
	manual, err := readPageManual(sc)
	if err != nil {
		return "", err
	}

	var middlePageSum int
	for _, pageUpdate := range manual.updates {
		if ordered := isOrdered(pageUpdate, manual.rules); !ordered {
			continue
		}

		middleIdx := len(pageUpdate) / 2 //nolint:mnd // Middle index is calculated based on the length of the slice.
		middlePage, err := strconv.Atoi(pageUpdate[middleIdx])
		if err != nil {
			return "", fmt.Errorf("parse middle page %q: %w", pageUpdate[middleIdx], err)
		}

		middlePageSum += middlePage
	}

	return strconv.Itoa(middlePageSum), nil
}

func isOrdered(pageUpdate []string, pageOrderRules map[string][]string) bool {
	indexMap := make(map[string]int)
	for i, page := range pageUpdate {
		indexMap[page] = i
	}

	for x, ys := range pageOrderRules {
		idxX, ok := indexMap[x]
		if !ok {
			continue // x is not in the page update
		}

		for _, y := range ys {
			if idxY, ok := indexMap[y]; ok && idxX > idxY {
				return false
			}
		}
	}

	return true
}
