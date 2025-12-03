package day3

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/ugurcsen/gods-generic/trees/binaryheap"
)

const bankOutputDigits = 12

func Part2(sc *bufio.Scanner) (string, error) {
	var totalJoltage int
	for sc.Scan() {
		bank := sc.Text()

		bankOutput := make([]int, len(bank))
		minHeap := binaryheap.NewWith[battery](compareBatteries)
		for i, bat := range bank {
			joltage := int(bat - '0')

			bankOutput[i] = joltage
			minHeap.Push(battery{
				joltage:   joltage,
				bankIndex: i,
			})
		}

		for range len(bank) - bankOutputDigits {
			minLeftBat, ok := minHeap.Pop()
			if !ok {
				return "", errors.New("something went wrong with heap, ran out of batteries to remove")
			}

			bankOutput[minLeftBat.bankIndex] = 0
		}

		var bestBankOutputSb strings.Builder
		for _, bat := range bankOutput {
			if bat != 0 {
				bestBankOutputSb.WriteString(strconv.Itoa(bat))
			}
		}

		bestBankOutput, err := strconv.Atoi(bestBankOutputSb.String())
		if err != nil {
			return "", fmt.Errorf("could not convert best bank output %s to int: %w", bestBankOutputSb.String(), err)
		}

		totalJoltage += bestBankOutput
	}

	if err := sc.Err(); err != nil {
		return "", fmt.Errorf("read input: %w", err)
	}

	return strconv.Itoa(totalJoltage), nil
}

type battery struct {
	joltage   int
	bankIndex int
}

func compareBatteries(a, b battery) int {
	if a.joltage > b.joltage {
		return 1
	}

	if a.joltage < b.joltage {
		return -1
	}

	if a.bankIndex > b.bankIndex {
		return 1
	}

	if b.bankIndex < b.bankIndex {
		return -1
	}

	return 0
}
