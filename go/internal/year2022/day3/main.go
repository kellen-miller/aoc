//nolint:forbidigo // this package is outdated and not worth updating
package main

import (
	"fmt"

	"github.com/kellen-miller/aoc/go/internal/year2022/day3/parts"
)

func main() {
	input := "internal/2022/day3/input.txt"

	fmt.Println("--- Day 3: Rucksack Reorganization ---")

	fmt.Println("Part 1:", "Total supplies priority =>", parts.SuppliesPriorityTotal(input))

	fmt.Println("Part 2:", "Total group badge priority =>", parts.BadgePriorityTotal(input))
}
