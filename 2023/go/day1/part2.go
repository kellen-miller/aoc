package day1

import (
	"bufio"
	"fmt"
	"strconv"
	"unicode"
)

func Part2(sc *bufio.Scanner) (string, error) {
	lenLookup, wordLookup := spelledNumberLookups()
	var sum int
	for sc.Scan() {
		line := sc.Text()

		left := findLeftChar(line, lenLookup, wordLookup)
		right := findRightChar(line, lenLookup, wordLookup)

		val, err := strconv.Atoi(string(left) + string(right))
		if err != nil {
			return "", fmt.Errorf("parse calibration value %q: %w", line, err)
		}

		sum += val
	}

	if err := sc.Err(); err != nil {
		return "", err
	}

	return strconv.Itoa(sum), nil
}

func findLeftChar(line string, lenLookup map[int][]string, wordLookup map[string]uint8) uint8 {
	for i := range len(line) {
		if unicode.IsDigit(rune(line[i])) {
			return line[i]
		}

		for k, v := range lenLookup {
			if len(line)-i < k {
				continue
			}

			for _, word := range v {
				if line[i:i+k] == word {
					return wordLookup[word]
				}
			}
		}
	}

	return 0
}

func findRightChar(line string, lenLookup map[int][]string, wordLookup map[string]uint8) uint8 {
	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(rune(line[i])) {
			return line[i]
		}

		for k, v := range lenLookup {
			if i < k {
				continue
			}

			for _, word := range v {
				if line[i-k+1:i+1] == word {
					return wordLookup[word]
				}
			}
		}
	}

	return 0
}

func spelledNumberLookups() (map[int][]string, map[string]uint8) {
	lenLookup := map[int][]string{
		3: {"one", "two", "six"},
		4: {"zero", "four", "five", "nine"},
		5: {"three", "seven", "eight"},
	}

	wordLookup := map[string]uint8{
		"zero":  '0',
		"one":   '1',
		"two":   '2',
		"three": '3',
		"four":  '4',
		"five":  '5',
		"six":   '6',
		"seven": '7',
		"eight": '8',
		"nine":  '9',
	}

	return lenLookup, wordLookup
}
