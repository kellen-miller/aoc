package main

import (
	"fmt"

	"github.com/kellen-miller/aoc/go/internal/year2022/day1/parts"
)

func main() {
	input := "internal/2022/day1/input.txt"

	fmt.Println("--- Day 1: Calorie Counting ---")

	fmt.Println("Part 1:", "Most calories being carried =>", parts.MostCalories(input))

	vals, sum := parts.Top3Calories(input)
	fmt.Println("Part 2:", fmt.Sprintf("Sum of top 3 calories carried = sum(%v) => %d\n", vals, sum))
}
