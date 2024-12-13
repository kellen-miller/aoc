package parts

import (
	"testing"
)

func Benchmark_Day3_Part1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SuppliesPriorityTotal("../input.txt")
	}
}

func Benchmark_Day3_Part2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BadgePriorityTotal("../input.txt")
	}
}

func Benchmark_Day3_Part2_ParallelGroups(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BadgePriorityTotalParallelGroups("../input.txt")
	}
}

func Benchmark_Day3_Part2_Channels(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BadgePriorityTotalChannels("../input.txt")
	}
}
