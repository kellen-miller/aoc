package parts

import (
	"testing"
)

func Benchmark_Day1_Part1(b *testing.B) {
	for b.Loop() {
		MostCalories("../input.txt")
	}
}

func Benchmark_Day1_Part2(b *testing.B) {
	for b.Loop() {
		Top3Calories("../input.txt")
	}
}
