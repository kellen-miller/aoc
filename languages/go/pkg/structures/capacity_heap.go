package structures

import (
	"cmp"
)

type CapacityHeap[T cmp.Ordered] struct {
	*MinHeap[T]
	Capacity int
}

func (h *CapacityHeap[T]) Push(val T) {
	if h.Len() < h.Capacity {
		h.MinHeap.Push(val)
		return
	}

	if h.Capacity == 0 {
		return
	}

	if peek := h.Peek(); peek < val {
		h.Pop()
		h.MinHeap.Push(val)
	}
}
