package structures

import (
	"slices"
	"testing"
)

func TestHeapKeepsLargestValuesInMinHeap(t *testing.T) {
	h := NewHeap[int](func(a, b int) bool { return a < b }, 3)
	for _, value := range []int{10, 20, 30, 40, 5} {
		h.Push(value)
	}

	if got := popValues(h); !slices.Equal(got, []int{20, 30, 40}) {
		t.Fatalf("got values %v, want %v", got, []int{20, 30, 40})
	}
}

func TestHeapKeepsSmallestValuesInMaxHeap(t *testing.T) {
	h := NewHeap[int](func(a, b int) bool { return a > b }, 3)
	for _, value := range []int{10, 20, 30, 5, 40} {
		h.Push(value)
	}

	if got := popValues(h); !slices.Equal(got, []int{20, 10, 5}) {
		t.Fatalf("got values %v, want %v", got, []int{20, 10, 5})
	}
}

func TestHeapPushMaintainsOrderingWithoutCapacity(t *testing.T) {
	h := NewHeap[int](func(a, b int) bool { return a < b }, 0)
	for _, value := range []int{4, 1, 5, 2, 3} {
		h.Push(value)
	}

	if got := popValues(h); !slices.Equal(got, []int{1, 2, 3, 4, 5}) {
		t.Fatalf("got values %v, want %v", got, []int{1, 2, 3, 4, 5})
	}
}

func TestHeapSupportsCustomValues(t *testing.T) {
	type task struct {
		name     string
		priority int
	}

	h := NewHeap[task](func(a, b task) bool {
		return a.priority < b.priority
	}, 0)
	for _, value := range []task{
		{name: "low", priority: 1},
		{name: "high", priority: 3},
		{name: "medium", priority: 2},
	} {
		h.Push(value)
	}

	values := make([]string, 0, h.Len())
	for h.Len() > 0 {
		values = append(values, h.Pop().name)
	}

	if want := []string{"low", "medium", "high"}; !slices.Equal(values, want) {
		t.Fatalf("got values %v, want %v", values, want)
	}
}

func popValues(h *Heap[int]) []int {
	values := make([]int, 0, h.Len())
	for h.Len() > 0 {
		values = append(values, h.Pop())
	}

	return values
}
