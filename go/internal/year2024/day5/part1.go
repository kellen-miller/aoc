package day5

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func (d *Day) Part1(sc *bufio.Scanner) (string, error) {
	var (
		parseUpdate    bool
		pageUpdates    [][]string
		pageOrderRules = make(map[string][]string)
	)
	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			parseUpdate = true
			continue
		}

		if parseUpdate {
			pageUpdates = append(pageUpdates, strings.Split(line, ","))
			continue
		}

		x, y, ok := strings.Cut(line, "|")
		if !ok {
			return "", fmt.Errorf("could not parse page order rule: %s", line)
		}

		pageOrderRules[x] = append(pageOrderRules[x], y)
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	var middlePageSum int
	for _, pageUpdate := range pageUpdates {
		if ordered := isOrdered(pageUpdate, pageOrderRules); !ordered {
			continue
		}

		middleIdx := len(pageUpdate) / 2
		middlePage, err := strconv.Atoi(pageUpdate[middleIdx])
		if err != nil {
			return "", err
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
