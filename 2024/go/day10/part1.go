package day10

import (
	"bufio"
	"strconv"

	"github.com/kellen-miller/aoc/languages/go/pkg/grid"
	"github.com/ugurcsen/gods-generic/sets/hashset"
)

func Part1(sc *bufio.Scanner) (string, error) {
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
				score += dfs(i, j, topographyMap, -1, hashset.New[grid.Point]())
			}
		}
	}

	return strconv.Itoa(score), nil
}

func dfs(i int, j int, topographyMap [][]int, prevHeight int, ninesSeen *hashset.Set[grid.Point]) int {
	if i < 0 || i >= len(topographyMap) || j < 0 || j >= len(topographyMap[0]) {
		return 0
	}

	height := topographyMap[i][j]
	if height-prevHeight != 1 {
		return 0
	}

	if height == 9 {
		if ninesSeen == nil {
			return 1
		}

		p := grid.Point{X: i, Y: j}
		if !ninesSeen.Contains(p) {
			ninesSeen.Add(p)
			return 1
		}
	}

	return dfs(i-1, j, topographyMap, height, ninesSeen) +
		dfs(i+1, j, topographyMap, height, ninesSeen) +
		dfs(i, j-1, topographyMap, height, ninesSeen) +
		dfs(i, j+1, topographyMap, height, ninesSeen)
}
