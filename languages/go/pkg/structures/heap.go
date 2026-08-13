package structures

import (
	"container/heap"
)

type Heap[T any] struct {
	data     []T
	less     func(T, T) bool
	capacity int
}

func NewHeap[T any](less func(T, T) bool, capacity int) *Heap[T] {
	h := &Heap[T]{
		data:     make([]T, 0),
		less:     less,
		capacity: capacity,
	}
	heap.Init(h)
	return h
}

func (h *Heap[T]) Len() int {
	return len(h.data)
}

func (h *Heap[T]) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h *Heap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *Heap[T]) Push(x any) {
	val := x.(T)
	if h.capacity > 0 && h.Len() >= h.capacity {
		if !h.less(h.data[0], val) {
			return
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

func (h *Heap[T]) Peek() (T, bool) {
	if len(h.data) == 0 {
		var zero T
		return zero, false
	}

	return h.data[0], true
}

func (h *Heap[T]) Values() []T {
	return h.data
}
