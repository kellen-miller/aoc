package parts

import (
	"container/heap"
	"strconv"

	"github.com/kellen-miller/aoc/2022/go/shared"
	"github.com/kellen-miller/aoc/languages/go/pkg/io"
	"github.com/kellen-miller/aoc/languages/go/pkg/structures"
)

const (
	capacity = 3
)

func Top3Calories(input string) ([]int, int) {
	if input == "" {
		input = shared.Input
	}

	sc, closeFile := io.GetScanner(input)
	defer closeFile()

	var (
		elfHeap     = structures.NewHeap[int](true, capacity)
		currentCals int
	)

	for sc.Scan() {
		cals, err := strconv.Atoi(sc.Text())
		currentCals += cals

		if err != nil {
			heap.Push(elfHeap, []int{currentCals})
			currentCals = 0
		}
	}

	var sum int
	for _, val := range elfHeap.Values() {
		sum += val[0]
	}

	res := make([]int, len(elfHeap.Values()))
	for i, v := range elfHeap.Values() {
		res[i] = v[0]
	}

	return res, sum
}
