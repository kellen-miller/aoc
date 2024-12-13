package structures

import (
	"github.com/ugurcsen/gods-generic/trees/binaryheap"
	"github.com/ugurcsen/gods-generic/utils"
)

type CapacityHeap[T utils.ComparableNumber] struct {
	Capacity int
	*binaryheap.Heap[T]
}

func (h *CapacityHeap[T]) Push(val T) {
	if h.Size() < h.Capacity {
		h.Heap.Push(val)
	} else {
		peek, notEmpty := h.Peek()
		if notEmpty && utils.NumberComparator[T](peek, val) < 0 {
			h.Heap.Pop()
			h.Heap.Push(val)
		}
	}
}
