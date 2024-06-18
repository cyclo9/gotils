package collections

import utils "github.com/cyclo9/gotils"

type Queue[T any] struct {
	list []T
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: make([]T, 0),
	}
}

func (q *Queue[T]) Enqueue(item T) {
	q.list = append(q.list, item)
}

func (q *Queue[T]) Dequeue() (item T, ok bool) {
	if len(q.list) == 0 {
		return utils.ZeroValue[T](), false
	}

	item = q.list[0]
	q.list = q.list[1:]
	return item, true
}

func (q *Queue[T]) IsEmpty() bool {
	return len(q.list) == 0
}

func (q *Queue[T]) Size() int {
	return len(q.list)
}
