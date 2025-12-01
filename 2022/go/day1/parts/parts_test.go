package parts

import (
	"testing"
)

func Benchmark_Day1_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MostCalories("../input.txt")
	}
}

func Benchmark_Day1_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Top3Calories("../input.txt")
	}
}
