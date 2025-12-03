package parts

import (
	"testing"
)

func Benchmark_Day3_Part1(b *testing.B) {
	for b.Loop() {
		SuppliesPriorityTotal("../input.txt")
	}
}

func Benchmark_Day3_Part2(b *testing.B) {
	for b.Loop() {
		BadgePriorityTotal("../input.txt")
	}
}

func Benchmark_Day3_Part2_ParallelGroups(b *testing.B) {
	for b.Loop() {
		BadgePriorityTotalParallelGroups("../input.txt")
	}
}

func Benchmark_Day3_Part2_Channels(b *testing.B) {
	for b.Loop() {
		BadgePriorityTotalChannels("../input.txt")
	}
}
