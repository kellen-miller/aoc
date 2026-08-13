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
		elfHeap     = structures.NewHeap[int](func(a, b int) bool { return a < b }, capacity)
		currentCals int
	)

	for sc.Scan() {
		cals, err := strconv.Atoi(sc.Text())
		currentCals += cals

		if err != nil {
			heap.Push(elfHeap, currentCals)
			currentCals = 0
		}
	}

	values := elfHeap.Values()
	var sum int
	for _, val := range values {
		sum += val
	}

	res := append([]int(nil), values...)
	return res, sum
}
