package day4

import (
	"bufio"
	"strconv"
)

type direction struct {
	row int
	col int
}

func Part1(sc *bufio.Scanner) (string, error) {
	var wordSearch [][]rune //nolint:prealloc // Preallocating is not possible here
	for sc.Scan() {
		wordSearch = append(wordSearch, []rune(sc.Text()))
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(findCountOfWord(wordSearch, "XMAS")), nil
}

func findCountOfWord(wordSearch [][]rune, word string) int {
	directions := []direction{
		{row: 0, col: 1},   // right
		{row: 0, col: -1},  // left
		{row: 1, col: 0},   // down
		{row: -1, col: 0},  // up
		{row: 1, col: 1},   // down right
		{row: 1, col: -1},  // down left
		{row: -1, col: 1},  // up right
		{row: -1, col: -1}, // up left
	}

	var count int
	for i := range wordSearch {
		for j := range wordSearch[i] {
			if rune(word[0]) != wordSearch[i][j] {
				continue
			}

			for _, dir := range directions {
				count += dfsFindCountOfWord(wordSearch, i, j, dir, word)
			}
		}
	}

	return count
}

func dfsFindCountOfWord(wordSearch [][]rune, i int, j int, dir direction, word string) int {
	if i < 0 || j < 0 || i >= len(wordSearch) || j >= len(wordSearch[i]) {
		return 0
	}

	if rune(word[0]) != wordSearch[i][j] {
		return 0
	}

	if len(word) == 1 {
		return 1
	}

	return dfsFindCountOfWord(wordSearch, i+dir.row, j+dir.col, dir, word[1:])
}
