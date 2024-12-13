package day4

import (
	"bufio"
	"strconv"
)

var upperDiagonals = [][]int{
	{-1, -1}, // upper left
	{-1, 1},  // upper right
}

var lowerDiagonals = [][]int{
	{1, 1},  // lower right
	{1, -1}, // lower left
}

func (d *Day) Part2(sc *bufio.Scanner) (string, error) {
	var wordSearch [][]rune //nolint:prealloc // We don't know the size of the word search.
	for sc.Scan() {
		wordSearch = append(wordSearch, []rune(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	var xmas int
	for i := range wordSearch {
		for j := range wordSearch[i] {
			if wordSearch[i][j] == 'A' && checkIsXMAS(wordSearch, i, j) {
				xmas++
			}
		}
	}

	return strconv.Itoa(xmas), nil
}

func checkIsXMAS(wordSearch [][]rune, i int, j int) bool {
	upperDiagonalChars, ok := checkUpperDiagonalsForMS(wordSearch, i, j)
	if !ok {
		return false
	}

	return checkDiagonalsNotEqual(wordSearch, i, j, upperDiagonalChars)
}

func checkUpperDiagonalsForMS(wordSearch [][]rune, i int, j int) ([]rune, bool) {
	chars := make([]rune, 0, len(upperDiagonals))
	for _, dir := range upperDiagonals {
		if i+dir[0] < 0 || j+dir[1] < 0 || j+dir[1] >= len(wordSearch[i]) {
			continue
		}

		char := wordSearch[i+dir[0]][j+dir[1]]
		if char != 'M' && char != 'S' {
			continue
		}

		chars = append(chars, char)
	}

	return chars, len(chars) == 2 //nolint:mnd // Two upper diagonals set if both are 'M' or 'S'.
}

func checkDiagonalsNotEqual(wordSearch [][]rune, i int, j int, upperDiagonalChars []rune) bool {
	for k, dir := range lowerDiagonals {
		if i+dir[0] >= len(wordSearch) || j+dir[1] >= len(wordSearch[i]) || j+dir[1] < 0 {
			return false
		}

		char := wordSearch[i+dir[0]][j+dir[1]]
		if char == upperDiagonalChars[k] {
			return false
		}

		if char != 'M' && char != 'S' {
			return false
		}
	}

	return true
}
