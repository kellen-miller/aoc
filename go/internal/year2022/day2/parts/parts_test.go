package parts

import (
	"testing"
)

func Benchmark_Day2_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		TotalScore("../input.txt")
	}
}

func Benchmark_Day2_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SetRoundResult("../input.txt")
	}
}
