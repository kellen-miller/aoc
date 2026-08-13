package structures

import (
	"container/heap"
	"slices"
	"testing"
)

func TestHeapKeepsLargestValuesInMinHeap(t *testing.T) {
	h := NewHeap[int](true, 3)
	for _, value := range []int{10, 20, 30, 40, 5} {
		heap.Push(h, []int{value})
	}

	if got := popValues(h); !slices.Equal(got, []int{20, 30, 40}) {
		t.Fatalf("got values %v, want %v", got, []int{20, 30, 40})
	}
}

func TestHeapKeepsSmallestValuesInMaxHeap(t *testing.T) {
	h := NewHeap[int](false, 3)
	for _, value := range []int{10, 20, 30, 5, 40} {
		heap.Push(h, []int{value})
	}

	if got := popValues(h); !slices.Equal(got, []int{20, 10, 5}) {
		t.Fatalf("got values %v, want %v", got, []int{20, 10, 5})
	}
}

func TestHeapPushMaintainsOrderingWithoutCapacity(t *testing.T) {
	h := NewHeap[int](true, 0)
	for _, value := range []int{4, 1, 5, 2, 3} {
		heap.Push(h, []int{value})
	}

	if got := popValues(h); !slices.Equal(got, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("got values %v, want %v", got, []int{1, 2, 3, 4, 5})
	}
}

func popValues(h *Heap[int]) []int {
	values := make([]int, 0, h.Len())
	for h.Len() > 0 {
		values = append(values, heap.Pop(h).([]int)[0])
	}

	return values
}
