package structures

import "cmp"

type Heap[T cmp.Ordered] struct {
	data  [][]T
	isMin bool
}

func NewHeap[T cmp.Ordered](isMin bool) *Heap[T] {
	return &Heap[T]{
		data:  make([][]T, 0),
		isMin: true,
	}
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
	h.data = append(h.data, x.([]T))
}

func (h *Heap[T]) Pop() any {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[:n-1]
	return x
}
