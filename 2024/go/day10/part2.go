package day10

import (
	"bufio"
	"strconv"
)

func Part2(sc *bufio.Scanner) (string, error) {
	topographyMap, err := readTopography(sc)
	if err != nil {
		return "", err
	}

	var score int
	for i := range topographyMap {
		for j := range topographyMap[i] {
			if topographyMap[i][j] == 0 {
				score += dfs(i, j, topographyMap, -1, nil)
			}
		}
	}

	return strconv.Itoa(score), nil
}
