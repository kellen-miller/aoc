package structures

import "cmp"

type genericSlice[T cmp.Ordered] struct {
	data []T
	less func(i, j int) bool
}

func (s *genericSlice[T]) Len() int {
	return len(s.data)
}

func (s *genericSlice[T]) Less(i, j int) bool {
	return s.less(i, j)
}

func (s *genericSlice[T]) Swap(i, j int) {
	s.data[i], s.data[j] = s.data[j], s.data[i]
}

func (s *genericSlice[T]) Push(x any) {
	s.data = append(s.data, x.(T))
}

func (s *genericSlice[T]) Pop() any {
	old := s.data
	n := len(old)
	item := old[n-1]
	s.data = old[0 : n-1]
	return item
}
