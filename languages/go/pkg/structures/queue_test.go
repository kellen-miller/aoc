package structures

import (
	"slices"
	"testing"
)

func TestQueueIsFIFO(t *testing.T) {
	q := NewQueue[int]()
	for _, value := range []int{1, 2, 3} {
		q.Enqueue(value)
	}

	if got := q.Size(); got != 3 {
		t.Fatalf("got size %d, want %d", got, 3)
	}

	values := make([]int, 0, q.Size())
	for !q.Empty() {
		value, ok := q.Dequeue()
		if !ok {
			t.Fatal("Dequeue reported an empty queue")
		}
		values = append(values, value)
	}

	if want := []int{1, 2, 3}; !slices.Equal(values, want) {
		t.Fatalf("got values %v, want %v", values, want)
	}
}

func TestQueuePeekDoesNotRemove(t *testing.T) {
	q := NewQueue[string]()
	q.Enqueue("first")
	q.Enqueue("second")

	value, ok := q.Peek()
	if !ok || value != "first" {
		t.Fatalf("got peek (%q, %t), want (%q, true)", value, ok, "first")
	}

	if got := q.Size(); got != 2 {
		t.Fatalf("got size %d after Peek, want %d", got, 2)
	}
}

func TestQueueSupportsCustomValues(t *testing.T) {
	type task struct {
		name string
	}

	q := NewQueue[task]()
	q.Enqueue(task{name: "first"})
	q.Enqueue(task{name: "second"})

	first, firstOK := q.Dequeue()
	second, secondOK := q.Dequeue()
	if !firstOK || !secondOK || first.name != "first" || second.name != "second" {
		t.Fatalf("got values (%+v, %t), (%+v, %t)", first, firstOK, second, secondOK)
	}
}

func TestQueueEmptyOperations(t *testing.T) {
	q := NewQueue[int]()

	if value, ok := q.Peek(); ok || value != 0 {
		t.Fatalf("got Peek (%d, %t) on empty queue", value, ok)
	}
	if value, ok := q.Dequeue(); ok || value != 0 {
		t.Fatalf("got Dequeue (%d, %t) on empty queue", value, ok)
	}
}

func TestDequeSupportsBothEnds(t *testing.T) {
	d := NewDeque[int]()
	d.PushBack(2)
	d.PushFront(1)
	d.PushBack(3)

	if value, ok := d.PopFront(); !ok || value != 1 {
		t.Fatalf("got front (%d, %t), want (1, true)", value, ok)
	}
	if value, ok := d.PopBack(); !ok || value != 3 {
		t.Fatalf("got back (%d, %t), want (3, true)", value, ok)
	}
	if value, ok := d.PopFront(); !ok || value != 2 {
		t.Fatalf("got remaining value (%d, %t), want (2, true)", value, ok)
	}
}

func TestStackIsLIFO(t *testing.T) {
	s := NewStack[string]()
	s.Push("first")
	s.Push("second")

	if value, ok := s.Pop(); !ok || value != "second" {
		t.Fatalf("got value (%q, %t), want (%q, true)", value, ok, "second")
	}
	if value, ok := s.Pop(); !ok || value != "first" {
		t.Fatalf("got value (%q, %t), want (%q, true)", value, ok, "first")
	}
}
