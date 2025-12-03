package day4

import (
	"bufio"
	"strconv"
)

func Part2(sc *bufio.Scanner) (string, error) {
	var wordSearch [][]rune //nolint:prealloc // We don't know the size of the word search.
	for sc.Scan() {
		wordSearch = append(wordSearch, []rune(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	upper := []direction{
		{row: -1, col: -1},
		{row: -1, col: 1},
	}
	lower := []direction{
		{row: 1, col: 1},
		{row: 1, col: -1},
	}

	var xmas int
	for i := range wordSearch {
		for j := range wordSearch[i] {
			if wordSearch[i][j] == 'A' && checkIsXMAS(wordSearch, i, j, upper, lower) {
				xmas++
			}
		}
	}

	return strconv.Itoa(xmas), nil
}

func checkIsXMAS(wordSearch [][]rune, i int, j int, upper, lower []direction) bool {
	upperDiagonalChars, ok := checkUpperDiagonalsForMS(wordSearch, i, j, upper)
	if !ok {
		return false
	}

	return checkDiagonalsNotEqual(wordSearch, i, j, upperDiagonalChars, lower)
}

func checkUpperDiagonalsForMS(wordSearch [][]rune, i int, j int, upper []direction) ([]rune, bool) {
	chars := make([]rune, 0, len(upper))
	for _, dir := range upper {
		if i+dir.row < 0 || j+dir.col < 0 || j+dir.col >= len(wordSearch[i]) {
			continue
		}

		char := wordSearch[i+dir.row][j+dir.col]
		if char != 'M' && char != 'S' {
			continue
		}

		chars = append(chars, char)
	}

	return chars, len(chars) == len(upper)
}

func checkDiagonalsNotEqual(wordSearch [][]rune, i int, j int, upperDiagonalChars []rune, lower []direction) bool {
	for k, dir := range lower {
		if i+dir.row >= len(wordSearch) || j+dir.col >= len(wordSearch[i]) || j+dir.col < 0 {
			return false
		}

		char := wordSearch[i+dir.row][j+dir.col]
		if char == upperDiagonalChars[k] {
			return false
		}

		if char != 'M' && char != 'S' {
			return false
		}
	}

	return true
}
