package structures

import (
	"cmp"
	"container/heap"
)

type Heap[T cmp.Ordered] struct {
	data     [][]T
	isMin    bool
	capacity int
}

func NewHeap[T cmp.Ordered](isMin bool, capacity int) *Heap[T] {
	h := &Heap[T]{
		data:     make([][]T, 0),
		isMin:    isMin,
		capacity: capacity,
	}
	heap.Init(h)
	return h
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Less(i, j int) bool {
	if h.isMin {
		return h.data[i][0] < h.data[j][0]
	}

	return h.data[i][0] > h.data[j][0]
}

func (h *Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) Push(x any) {
	val := x.([]T)
	if h.capacity > 0 && h.Len() >= h.capacity {
		if h.isMin {
			if val[0] <= h.data[0][0] {
				return
			}
		} else {
			if val[0] >= h.data[0][0] {
				return
			}
		}

		heap.Pop(h)
	}

	h.data = append(h.data, val)
	heap.Fix(h, h.Len()-1)
}

func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	item := old[n-1]
	h.data = old[0 : n-1]
	return item
}

func (h *Heap[T]) Peek() []T {
	if len(h.data) == 0 {
		return nil
	}

	return h.data[0]
}

func (h *Heap[T]) Values() [][]T {
	return h.data
}
