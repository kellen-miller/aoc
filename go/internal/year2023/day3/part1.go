package day3

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/kellen-miller/advent-of-code/go/pkg/io"
)

func FindValidParts(input string) int {
	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		sum                    int
		potentialPartLocations [][]int
		lastLine               string
		lineIdx                int
	)

	for sc.Scan() {
		line := sc.Text()
		potentialPartLocations = append(potentialPartLocations, make([]int, len(line)))

		for i := 0; i < len(line); i++ {
			if line[i] == '.' {
				continue
			}

			r := rune(line[i])
			if unicode.IsDigit(r) {
				part, digits, isPartNumber := extractPotentialPart(line, lastLine, i)
				if isPartNumber {
					sum += part
				} else {
					recordPotentialPartLocation(potentialPartLocations[lineIdx], part, digits, i)
				}

				i += digits - 1
			} else if lineIdx > 0 && isSymbol(r) {
				sum += checkLastLineForAdjacentPart(potentialPartLocations[lineIdx-1], i)
			}
		}

		lastLine = line
		lineIdx++
	}

	return sum
}

func extractPotentialPart(line string, lastLine string, startIndex int) (int, int, bool) {
	var (
		nb           strings.Builder
		isPartNumber bool
	)

	for i := startIndex; i < len(line) && unicode.IsDigit(rune(line[i])); i++ {
		nb.WriteByte(line[i])
		if !isPartNumber {
			isPartNumber = checkForAdjacentSymbol(line, lastLine, i)
		}
	}

	number, err := strconv.Atoi(nb.String())
	if err != nil {
		panic("invalid part number")
	}

	return number, nb.Len(), isPartNumber
}

func recordPotentialPartLocation(locations []int, number int, digits int, startIndex int) {
	for i := startIndex; i < len(locations); i++ {
		locations[i] = number

		if timesRecorded := i - startIndex; timesRecorded == digits-1 {
			return
		}
	}
}

func checkForAdjacentSymbol(currLine string, lastLine string, i int) bool {
	if i > 0 && isSymbol(rune(currLine[i-1])) {
		return true
	}

	if i < len(currLine)-1 && isSymbol(rune(currLine[i+1])) {
		return true
	}

	if lastLine != "" {
		if i > 0 && isSymbol(rune(lastLine[i-1])) {
			return true
		}

		if i < len(lastLine)-1 && isSymbol(rune(lastLine[i+1])) {
			return true
		}

		if isSymbol(rune(lastLine[i])) {
			return true
		}
	}

	return false
}

func checkLastLineForAdjacentPart(lastLine []int, i int) int {
	var sum int

	if lastLine[i] > 0 {
		sum += lastLine[i]
		zeroSamePart(lastLine, i) // zero out the rest of locations recorded for the part so we don't double count
		return sum
	}

	if i > 0 && lastLine[i-1] > 0 {
		sum += lastLine[i-1]
	}

	if i < len(lastLine)-1 && lastLine[i+1] > 0 {
		sum += lastLine[i+1]
		zeroSamePart(lastLine, i+1) // zero out the rest of locations recorded for the part so we don't double count
	}

	return sum
}

func isSymbol(c rune) bool {
	return !unicode.IsDigit(c) && c != '.'
}

func zeroSamePart(line []int, col int) {
	number := line[col]

	for i := col; i < len(line) && line[i] == number; i++ {
		line[i] = 0
	}
}
