package queue

type node[T any] struct {
	value T
	next  *node[T]
}

func newNode[T any](value T) *node[T] {
	return &node[T]{
		value: value,
		next:  nil,
	}
}

type Queue[T any] struct {
	head *node[T]
	tail *node[T]
}

func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{}
}

func (q *Queue[T]) IsEmpty() bool {
	return q.head == nil
}

func (q *Queue[T]) Peek() T {
	return q.head.value
}

func (q *Queue[T]) Add(value T) {
	node := newNode(value)
	if q.tail != nil {
		q.tail.next = node
	}
	q.tail = node

	if q.head == nil {
		q.head = node
	}
}

func (q *Queue[T]) Remove() T {
	value := q.head.value
	q.head = q.head.next
	if q.head == nil {
		q.tail = nil
	}
	return value
}
