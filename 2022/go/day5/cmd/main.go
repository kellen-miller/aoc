//nolint:forbidigo // this package is outdated and not worth updating
package main

import (
	"fmt"

	"github.com/kellen-miller/aoc/2022/go/day5/parts"
)

func main() {
	input := "2022/go/day5/input.txt"

	fmt.Println("--- Day 5: Supply Stacks ---")

	fmt.Println("Part 1:", "Rearrange crates =>", parts.RearrangeCrates(input))

	fmt.Println("Part 2:", "Rearrange multiple crates =>", parts.RearrangeCratesMulti(input))
}
