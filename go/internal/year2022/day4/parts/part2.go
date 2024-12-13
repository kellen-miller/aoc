package parts

import (
	"strings"

	"github.com/kellen-miller/advent-of-code/go/internal"
	"github.com/kellen-miller/advent-of-code/go/pkg/io"
)

func OverlappingSections(input string) int {
	if input == "" {
		input = internal.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var total int
	for sc.Scan() {
		total += isOverlap(sc.Text())
	}

	return total
}

func isOverlap(pairs string) int {
	var (
		elves = strings.Split(pairs, ",")

		elf1Min, elf1Max = getElfSections(elves[0])
		elf2Min, elf2Max = getElfSections(elves[1])
	)

	if (elf1Min <= elf2Min && elf2Min <= elf1Max) ||
		(elf2Min <= elf1Min && elf1Min <= elf2Max) {
		return 1
	}

	return 0
}
