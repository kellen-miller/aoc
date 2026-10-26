package day5

import (
	"bufio"
	"fmt"
	"strings"
)

type pageManual struct {
	rules   map[string][]string
	updates [][]string
}

func readPageManual(sc *bufio.Scanner) (pageManual, error) {
	manual := pageManual{rules: make(map[string][]string)}
	parseUpdate := false

	for sc.Scan() {
		line := sc.Text()
		if line == "" {
			parseUpdate = true
			continue
		}

		if parseUpdate {
			manual.updates = append(manual.updates, strings.Split(line, ","))
			continue
		}

		x, y, ok := strings.Cut(line, "|")
		if !ok {
			return pageManual{}, fmt.Errorf("could not parse page order rule: %s", line)
		}

		manual.rules[x] = append(manual.rules[x], y)
	}

	if err := sc.Err(); err != nil {
		return pageManual{}, fmt.Errorf("scan page updates: %w", err)
	}

	return manual, nil
}
