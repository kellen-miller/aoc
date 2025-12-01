package day4

import (
	"bufio"
	"strconv"
)

var dirs = [][]int{
	{0, 1},   // right
	{0, -1},  // left
	{1, 0},   // down
	{-1, 0},  // up
	{1, 1},   // down right
	{1, -1},  // down left
	{-1, 1},  // up right
	{-1, -1}, // up left
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
	var count int
	for i := range wordSearch {
		for j := range wordSearch[i] {
			if rune(word[0]) != wordSearch[i][j] {
				continue
			}

			for _, dir := range dirs {
				count += dfsFindCountOfWord(wordSearch, i, j, dir, word)
			}
		}
	}

	return count
}

func dfsFindCountOfWord(wordSearch [][]rune, i int, j int, dir []int, word string) int {
	if i < 0 || j < 0 || i >= len(wordSearch) || j >= len(wordSearch[i]) {
		return 0
	}

	if rune(word[0]) != wordSearch[i][j] {
		return 0
	}

	if len(word) == 1 {
		return 1
	}

	return dfsFindCountOfWord(wordSearch, i+dir[0], j+dir[1], dir, word[1:])
}
