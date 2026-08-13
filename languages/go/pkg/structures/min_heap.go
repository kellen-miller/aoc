package structures

import (
	"cmp"
	"container/heap"
)

type MinHeap[T cmp.Ordered] struct {
	data genericSlice[T]
}

// NewMinHeap initializes the heap.
func NewMinHeap[T cmp.Ordered]() *MinHeap[T] {
	h := &MinHeap[T]{
		data: genericSlice[T]{
			data: make([]T, 0),
		},
	}
	h.data.less = func(i, j int) bool { return h.data.data[i] < h.data.data[j] }
	heap.Init(&h.data)
	return h
}

// Push adds an element to the heap safely.
func (h *MinHeap[T]) Push(val T) {
	heap.Push(&h.data, val)
}

// Pop removes and returns the minimum element.
func (h *MinHeap[T]) Pop() T {
	return heap.Pop(&h.data).(T)
}

// Peek returns the minimum element without removing it.
func (h *MinHeap[T]) Peek() T {
	return h.data.data[0]
}

// Len returns the current count of elements.
func (h *MinHeap[T]) Len() int {
	return h.data.Len()
}

// Values returns all elements in the heap.
func (h *MinHeap[T]) Values() []T {
	res := make([]T, len(h.data.data))
	copy(res, h.data.data)
	return res
}
