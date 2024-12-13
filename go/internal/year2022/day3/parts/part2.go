package parts

import (
	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/pkg/io"
)

const (
	groupSize = 3
)

func BadgePriorityTotal(input string) int {
	if input == "" {
		input = internal.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var total int
	for {
		groups := make([]string, groupSize)
		for i := 0; i < groupSize; i++ {
			if !sc.Scan() {
				return total
			}

			groups[i] = sc.Text()
		}

		total += findGroupBadgePriority(groups)
	}
}

func findGroupBadgePriority(groups []string) int {
	countGroupsBadgeFound := make([]int, maxLetterDifference)

	for _, group := range groups {
		badgesFoundInGroup := make([]bool, maxLetterDifference)

		for _, badge := range group {
			badgeVal := badge - 'A'

			if !badgesFoundInGroup[badgeVal] {
				badgesFoundInGroup[badgeVal] = true
				countGroupsBadgeFound[badgeVal]++
			}

			if countGroupsBadgeFound[badgeVal] == len(groups) {
				return letterScore(badge)
			}
		}
	}

	return 0
}
