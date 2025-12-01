//nolint:errcheck // package is outdated and not maintained
package parts

import (
	"strconv"
	"strings"

	"github.com/kellen-miller/aoc/2022/go/shared"
	"github.com/kellen-miller/aoc/languages/go/pkg/io"
)

func RedundantCleanup(input string) int {
	if input == "" {
		input = shared.Input
	}

	sc, closeFn := io.GetScanner(input)
	defer closeFn()

	var total int
	for sc.Scan() {
		total += isAssignmentRedundant(sc.Text())
	}

	return total
}

func isAssignmentRedundant(pairs string) int {
	var (
		elves = strings.Split(pairs, ",")

		elf1Min, elf1Max = getElfSections(elves[0])
		elf2Min, elf2Max = getElfSections(elves[1])
	)

	if (elf1Min <= elf2Min && elf1Max >= elf2Max) ||
		(elf2Min <= elf1Min && elf2Max >= elf1Max) {
		return 1
	}

	return 0
}

func getElfSections(elf string) (int, int) {
	var (
		sections = strings.Split(elf, "-")

		section1, _ = strconv.Atoi(sections[0])
		section2, _ = strconv.Atoi(sections[1])
	)

	return section1, section2
}
