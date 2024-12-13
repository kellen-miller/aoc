package parts

import (
	"strconv"

	"github.com/kellen-miller/aoc/go/internal"
	"github.com/kellen-miller/aoc/go/pkg/io"
	"github.com/kellen-miller/aoc/go/pkg/structures"
	"github.com/ugurcsen/gods-generic/trees/binaryheap"
)

const (
	capacity = 3
)

func Top3Calories(input string) ([]int, int) {
	if input == "" {
		input = internal.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		elfHeap = &structures.CapacityHeap[int]{
			Capacity: capacity,
			Heap:     binaryheap.NewWithNumberComparator[int](),
		}
		currentCals int
	)

	for sc.Scan() {
		cals, err := strconv.Atoi(sc.Text())
		currentCals += cals

		if err != nil {
			elfHeap.Push(currentCals)
			currentCals = 0
		}
	}

	var sum int
	for _, val := range elfHeap.Values() {
		sum += val
	}

	return elfHeap.Values(), sum
}
