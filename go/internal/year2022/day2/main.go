//nolint:forbidigo // this package is outdated and not worth updating
package main

import (
	"fmt"

	"github.com/kellen-miller/aoc/go/internal/year2022/day2/parts"
)

func main() {
	input := "internal/2022/day2/input.txt"

	fmt.Println("--- Day 2: Rock Paper Scissors ---")

	fmt.Println("Part 1:", "Total score =>", parts.TotalScore(input))

	fmt.Println("Part 2:", "Set round result =>", parts.SetRoundResult(input))
}
