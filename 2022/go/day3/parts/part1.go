package parts

import (
	"github.com/kellen-miller/aoc/2022/go/shared"
	"github.com/kellen-miller/aoc/languages/go/pkg/io"
)

func SuppliesPriorityTotal(input string) int {
	if input == "" {
		input = shared.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var total int
	for sc.Scan() {
		total += findRucksackPriority(sc.Text())
	}
	return total
}

const (
	// Subtracting 'A' from all chars, the max difference will be 'z' (122) - 'A' (65) = 57
	// Adding 1 to the difference to account for the 0 index.
	maxLetterDifference = 'z' - 'A' + 1
)

func letterScore(letter rune) int {
	// 'a' - 'a' = 0 + 1 = 1
	if letter > 'Z' {
		return int(letter-'a') + 1
	}

	// 'A' - 'A' = 0 + 26 + 1 = 27
	return int(letter-'A') + 26 + 1
}

func findRucksackPriority(line string) int {
	var (
		compartment1Counts = make([]int, maxLetterDifference)
		compartment2Start  = len(line) / 2
	)

	for i, char := range line {
		if i < compartment2Start {
			compartment1Counts[char-'A']++
		} else if compartment1Counts[char-'A'] > 0 {
			return letterScore(char)
		}
	}

	return 0
}
