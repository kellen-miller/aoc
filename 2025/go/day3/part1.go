package day3

import (
	"bufio"
	"fmt"
	"strconv"
)

const minBankLength = 2

func Part1(sc *bufio.Scanner) (string, error) {
	var totalJoltage int64
	for sc.Scan() {
		totalJoltage += bestBankJoltage(sc.Text())
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.FormatInt(totalJoltage, 10), nil
}

func bestBankJoltage(bank string) int64 {
	if len(bank) < minBankLength {
		return 0
	}

	var (
		lastIndex        = len(bank) - 1
		bestLeadingDigit = -1
		bestPairSeen     int64
	)

	for i, battery := range bank {
		batteryJoltage := int(battery - '0')

		if bestLeadingDigit >= 0 {
			potentialPair := int64(bestLeadingDigit*10 + batteryJoltage)
			if potentialPair > bestPairSeen {
				bestPairSeen = potentialPair
			}
		}

		if i == lastIndex {
			break
		}

		if batteryJoltage >= bestLeadingDigit {
			bestLeadingDigit = batteryJoltage
		}
	}

	return bestPairSeen
}
