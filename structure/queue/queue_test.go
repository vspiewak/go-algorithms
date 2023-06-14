package queue

import (
	"reflect"
	"testing"
)

func TestQueue(t *testing.T) {

	t.Run("Test Add & Remove", func(t *testing.T) {

		q := NewQueue[int]()
		q.Add(1)
		q.Add(2)
		q.Add(3)

		one := q.Remove()
		two := q.Remove()
		three := q.Remove()

		expected := []int{1, 2, 3}
		actual := []int{one, two, three}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Peek", func(t *testing.T) {

		q := NewQueue[int]()
		q.Add(1)
		one := q.Peek()
		q.Add(2)
		two := q.Peek()
		q.Add(3)
		three := q.Peek()

		expected := []int{1, 1, 1}
		actual := []int{one, two, three}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test IsEmpty", func(t *testing.T) {

		q := NewQueue[int]()
		one := q.IsEmpty()
		q.Add(1)
		two := q.IsEmpty()
		q.Add(2)
		three := q.IsEmpty()
		q.Remove()
		q.Remove()
		four := q.IsEmpty()

		expected := []bool{true, false, false, true}
		actual := []bool{one, two, three, four}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
