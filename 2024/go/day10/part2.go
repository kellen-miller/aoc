package day10

import (
	"bufio"
	"strconv"
)

func Part2(sc *bufio.Scanner) (string, error) {
	var (
		row           int
		topographyMap [][]int
	)
	for sc.Scan() {
		line := sc.Text()
		topographyRow := make([]int, len(line))
		for i, c := range line {
			ci, err := strconv.Atoi(string(c))
			if err != nil {
				return "", err
			}

			topographyRow[i] = ci
		}

		topographyMap = append(topographyMap, topographyRow)
		row++
	}

	if err := sc.Err(); err != nil {
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
