package parts

import (
	"testing"
)

func Benchmark_Day2_Part1(b *testing.B) {
	for b.Loop() {
		TotalScore("../input.txt")
	}
}

func Benchmark_Day2_Part2(b *testing.B) {
	for b.Loop() {
		SetRoundResult("../input.txt")
	}
}
