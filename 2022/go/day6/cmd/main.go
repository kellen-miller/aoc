//nolint:forbidigo // this package is outdated and not worth updating
package main

import (
	"fmt"

	"github.com/kellen-miller/aoc/2022/go/day6/parts"
)

func main() {
	input := "2022/go/day6/input.txt"

	fmt.Println("--- Day 6: Tuning Trouble ---")

	fmt.Println("Part 1:", "Packet Start Marker =>", parts.StartOfPacket(input))

	fmt.Println("Part 2:", "Message Start Marker =>", parts.StartOfMessage(input))
}
