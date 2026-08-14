package structures

import (
	"container/heap"
)

type Heap[T any] struct {
	backingHeap backingHeap[T]
	capacity    int
}

func NewHeap[T any](less func(T, T) bool, capacity int) *Heap[T] {
	h := &Heap[T]{
		backingHeap: backingHeap[T]{
			data: make([]T, 0),
			less: less,
		},
		capacity: capacity,
	}

	heap.Init(&h.backingHeap)
	return h
}

func (h *Heap[T]) Len() int {
	return h.backingHeap.Len()
}

func (h *Heap[T]) Push(val T) {
	if h.capacity > 0 && h.Len() >= h.capacity {
		if !h.backingHeap.less(h.backingHeap.data[0], val) {
			return
		}

		heap.Pop(&h.backingHeap)
	}

	heap.Push(&h.backingHeap, val)
}

func (h *Heap[T]) Pop() T {
	return heap.Pop(&h.backingHeap).(T)
}

func (h *Heap[T]) Peek() (T, bool) {
	if h.Len() == 0 {
		var zero T
		return zero, false
	}

	return h.backingHeap.data[0], true
}

func (h *Heap[T]) Values() []T {
	return h.backingHeap.data
}

type backingHeap[T any] struct {
	data []T
	less func(T, T) bool
}

func (h *backingHeap[T]) Len() int {
	return len(h.data)
}

func (h *backingHeap[T]) Less(i, j int) bool {
	return h.less(h.data[i], h.data[j])
}

func (h *backingHeap[T]) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h *backingHeap[T]) Push(x any) {
	h.data = append(h.data, x.(T))
}

func (h *backingHeap[T]) Pop() any {
	old := h.data
	n := len(old)
	item := old[n-1]
	h.data = old[:n-1]

	return item
}
