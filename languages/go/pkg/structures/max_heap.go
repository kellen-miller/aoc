package structures

import (
	"cmp"
	"container/heap"
)

type MaxHeap[T cmp.Ordered] struct {
	data genericSlice[T]
}

// NewMaxHeap initializes the heap.
func NewMaxHeap[T cmp.Ordered]() *MaxHeap[T] {
	h := &MaxHeap[T]{data: genericSlice[T]{
		data: make([]T, 0),
	}}
	h.data.less = func(i, j int) bool { return h.data.data[i] > h.data.data[j] }
	heap.Init(&h.data)
	return h
}

// Push adds an element to the heap safely.
func (h *MaxHeap[T]) Push(val T) {
	heap.Push(&h.data, val)
}

// Pop removes and returns the maximum element.
func (h *MaxHeap[T]) Pop() T {
	return heap.Pop(&h.data).(T)
}

// Peek returns the maximum element without removing it.
func (h *MaxHeap[T]) Peek() T {
	return h.data.data[0]
}

// Len returns the current count of elements.
func (h *MaxHeap[T]) Len() int {
	return h.data.Len()
}

// Values returns all elements in the heap.
func (h *MaxHeap[T]) Values() []T {
	res := make([]T, len(h.data.data))
	copy(res, h.data.data)
	return res
}
