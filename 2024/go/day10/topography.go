package day10

import (
	"bufio"
	"fmt"
	"strconv"
)

const (
	defaultTopographyRows = 64
	maxHeightValue        = 9
)

func readTopography(sc *bufio.Scanner) ([][]int, error) {
	topographyMap := make([][]int, 0, defaultTopographyRows)
	for sc.Scan() {
		line := sc.Text()
		row := make([]int, len(line))
		for i, c := range line {
			value, err := strconv.Atoi(string(c))
			if err != nil {
				return nil, fmt.Errorf("parse elevation at col %d: %w", i, err)
			}
			row[i] = value
		}
		topographyMap = append(topographyMap, row)
	}

	if err := sc.Err(); err != nil {
		return nil, err
	}

	return topographyMap, nil
}
