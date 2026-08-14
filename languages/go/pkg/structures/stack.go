package structures

type Stack[T any] struct {
	data Deque[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(value T) {
	s.data.PushBack(value)
}

func (s *Stack[T]) Pop() (T, bool) {
	return s.data.PopBack()
}

func (s *Stack[T]) Peek() (T, bool) {
	return s.data.PeekBack()
}

func (s *Stack[T]) Size() int {
	return s.data.Size()
}

func (s *Stack[T]) Empty() bool {
	return s.data.Empty()
}
