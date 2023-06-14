package linkedlist

import "fmt"

type node[T comparable] struct {
	value T
	next  *node[T]
}

func newNode[T comparable](value T) *node[T] {
	return &node[T]{
		value: value,
		next:  nil,
	}
}

type LinkedList[T comparable] struct {
	head *node[T]
}

func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) Append(value T) {
	node := newNode(value)

	if l.head == nil {
		l.head = node
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = node
	}
}

func (l *LinkedList[T]) Prepend(value T) {
	node := newNode(value)
	node.next = l.head
	l.head = node
}

func (l *LinkedList[T]) Delete(value T) {
	if l.head == nil {
		return
	}
	if l.head.value == value {
		l.head = l.head.next
	}

	for current := l.head; current.next != nil; current = current.next {
		if current.next.value == value {
			current.next = current.next.next
		}
	}

}

func (l *LinkedList[T]) ToSlice() []T {
	slice := []T{}
	for current := l.head; current != nil; current = current.next {
		slice = append(slice, current.value)
	}
	return slice
}

/*
	func (l *LinkedList[T]) Display() {
		for current := l.head; current != nil; current = current.next {
			fmt.Print(current.value, " ")
		}
		fmt.Print("\n")
	}
*/
func (l *LinkedList[T]) Display() {
	fmt.Println(l.ToSlice())
}
