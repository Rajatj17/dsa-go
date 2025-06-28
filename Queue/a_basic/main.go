package basic

type Queue[T any] struct {
	items []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) Dequeue() T {
	top := q.items[0]

	q.items = q.items[1:]

	return top
}

func (q *Queue[T]) Peek() T {
	return q.items[0]
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.items) <= 0
}
