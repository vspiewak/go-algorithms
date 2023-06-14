package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {

	t.Run("Test Push & Pop", func(t *testing.T) {

		s := NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		one := s.Pop()
		two := s.Pop()
		three := s.Pop()

		expected := []int{3, 2, 1}
		actual := []int{one, two, three}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Peek", func(t *testing.T) {

		s := NewStack[int]()

		s.Push(1)
		one := s.Peek()

		s.Push(2)
		two := s.Peek()

		s.Push(3)
		three := s.Peek()

		expected := []int{1, 2, 3}
		actual := []int{one, two, three}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test IsEmpty", func(t *testing.T) {

		s := NewStack[int]()
		one := s.IsEmpty()

		s.Push(1)
		two := s.IsEmpty()

		s.Push(2)
		three := s.IsEmpty()

		s.Pop()
		s.Pop()
		four := s.IsEmpty()

		expected := []bool{true, false, false, true}
		actual := []bool{one, two, three, four}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
