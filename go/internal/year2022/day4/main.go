package main

import (
	"fmt"

	"github.com/kellen-miller/advent-of-code/go/internal/year2022/day4/parts"
)

func main() {
	input := "internal/2022/day4/input.txt"

	fmt.Println("--- Day 4: Camp Cleanup ---")

	fmt.Println("Part 1:", "Redundant pairs =>", parts.RedundantCleanup(input))

	fmt.Println("Part 2:", "Overlapping pairs =>", parts.OverlappingSections(input))
}
