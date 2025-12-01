//nolint:forbidigo // this package is outdated and not worth updating
package main

import (
	"fmt"

	"github.com/kellen-miller/aoc/2022/go/day4/parts"
)

func main() {
	input := "2022/go/day4/input.txt"

	fmt.Println("--- Day 4: Camp Cleanup ---")

	fmt.Println("Part 1:", "Redundant pairs =>", parts.RedundantCleanup(input))

	fmt.Println("Part 2:", "Overlapping pairs =>", parts.OverlappingSections(input))
}
