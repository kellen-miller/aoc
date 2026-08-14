package structures

type Deque[T any] struct {
	data []T
	head int
	size int
}

func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{}
}

func (d *Deque[T]) PushFront(value T) {
	d.ensureCapacity()
	d.head = (d.head - 1 + len(d.data)) % len(d.data)
	d.data[d.head] = value
	d.size++
}

func (d *Deque[T]) PushBack(value T) {
	d.ensureCapacity()
	d.data[(d.head+d.size)%len(d.data)] = value
	d.size++
}

func (d *Deque[T]) PopFront() (T, bool) {
	if d.Empty() {
		var zero T
		return zero, false
	}

	value := d.data[d.head]
	var zero T
	d.data[d.head] = zero
	d.head = (d.head + 1) % len(d.data)
	d.size--
	if d.size == 0 {
		d.head = 0
	}

	return value, true
}

func (d *Deque[T]) PopBack() (T, bool) {
	if d.Empty() {
		var zero T
		return zero, false
	}

	index := (d.head + d.size - 1) % len(d.data)
	value := d.data[index]
	var zero T
	d.data[index] = zero
	d.size--
	if d.size == 0 {
		d.head = 0
	}

	return value, true
}

func (d *Deque[T]) PeekFront() (T, bool) {
	if d.Empty() {
		var zero T
		return zero, false
	}

	return d.data[d.head], true
}

func (d *Deque[T]) PeekBack() (T, bool) {
	if d.Empty() {
		var zero T
		return zero, false
	}

	return d.data[(d.head+d.size-1)%len(d.data)], true
}

func (d *Deque[T]) Size() int {
	return d.size
}

func (d *Deque[T]) Empty() bool {
	return d.size == 0
}

func (d *Deque[T]) ensureCapacity() {
	if d.size < len(d.data) {
		return
	}

	capacity := 8
	if len(d.data) > 0 {
		capacity = len(d.data) * 2
	}

	data := make([]T, capacity)
	for i := 0; i < d.size; i++ {
		data[i] = d.data[(d.head+i)%len(d.data)]
	}

	d.data = data
	d.head = 0
}
