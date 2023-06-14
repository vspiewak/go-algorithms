package linkedlist

import (
	"reflect"
	"testing"
)

func TestLinkedList(t *testing.T) {

	t.Run("Test Append", func(t *testing.T) {

		l := NewLinkedList[int]()
		l.Append(1)
		l.Append(2)
		l.Append(3)

		expected := []int{1, 2, 3}
		actual := l.ToSlice()

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Prepend", func(t *testing.T) {

		l := NewLinkedList[int]()
		l.Prepend(1)
		l.Prepend(2)
		l.Prepend(3)

		expected := []int{3, 2, 1}
		actual := l.ToSlice()

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Delete", func(t *testing.T) {

		l := NewLinkedList[int]()
		l.Append(1)
		l.Append(2)
		l.Append(3)
		l.Delete(2)

		expected := []int{1, 3}
		actual := l.ToSlice()

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
