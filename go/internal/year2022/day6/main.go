//nolint:forbidigo // this package is outdated and not worth updating
package main

import (
	"fmt"

	"github.com/kellen-miller/aoc/go/internal/year2022/day6/parts"
)

func main() {
	input := "internal/2022/day6/input.txt"

	fmt.Println("--- Day 6: Tuning Trouble ---")

	fmt.Println("Part 1:", "Packet Start Marker =>", parts.StartOfPacket(input))

	fmt.Println("Part 2:", "Message Start Marker =>", parts.StartOfMessage(input))
}
