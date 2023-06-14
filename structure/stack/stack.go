package stack

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

type Stack[T any] struct {
	top *node[T]
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) IsEmpty() bool {
	return s.top == nil
}

func (s *Stack[T]) Peek() T {
	return s.top.value
}

func (s *Stack[T]) Push(value T) {
	node := newNode(value)
	node.next = s.top
	s.top = node
}

func (s *Stack[T]) Pop() T {
	value := s.top.value
	s.top = s.top.next
	return value
}
