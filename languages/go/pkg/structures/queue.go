package structures

type Queue[T any] struct {
	data Deque[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) Enqueue(value T) {
	q.data.PushBack(value)
}

func (q *Queue[T]) Dequeue() (T, bool) {
	return q.data.PopFront()
}

func (q *Queue[T]) Peek() (T, bool) {
	return q.data.PeekFront()
}

func (q *Queue[T]) Size() int {
	return q.data.Size()
}

func (q *Queue[T]) Empty() bool {
	return q.data.Empty()
}
