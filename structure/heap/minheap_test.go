package heap

import (
	"reflect"
	"testing"
)

func TestMinHeap(t *testing.T) {

	t.Run("Test LeftRightParentIndex", func(t *testing.T) {

		h := NewMinHeap[int]()

		expected := []int{
			1, 2,
			3, 4, 0,
			5, 6, 0,
		}
		actual := []int{
			h.leftChildIndex(0), h.rightChildIndex(0),
			h.leftChildIndex(1), h.rightChildIndex(1), h.parentIndex(1),
			h.leftChildIndex(2), h.rightChildIndex(2), h.parentIndex(2),
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test HasLeftRightParent", func(t *testing.T) {

		h := NewMinHeap[int]()
		h.items = []int{
			10, 15, 20, 17, 25,
		}

		expected := []bool{
			true, true, false,
			true, true, true,
			false, false, true,
		}
		actual := []bool{
			h.hasLeftChild(0), h.hasRightChild(0), h.hasParent(0),
			h.hasLeftChild(1), h.hasRightChild(1), h.hasParent(1),
			h.hasLeftChild(2), h.hasRightChild(2), h.hasParent(2),
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test HasLeftRight", func(t *testing.T) {

		h := NewMinHeap[int]()
		h.items = []int{
			10, 15, 20, 17, 25,
		}

		expected := []int{
			15, 20,
			17, 25, 10,
			10,
		}
		actual := []int{
			h.leftChild(0), h.rightChild(0),
			h.leftChild(1), h.rightChild(1), h.parent(1),
			h.parent(2),
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test swap", func(t *testing.T) {

		h := NewMinHeap[int]()
		h.items = []int{1, 2}
		h.swap(0, 1)
		expected := []int{2, 1}
		actual := h.items

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test ToSlice", func(t *testing.T) {

		h := NewMinHeap[int]()
		h.items = []int{1, 2, 3}

		s := h.ToSlice()
		s[0] = 10

		actual := h.ToSlice()
		expected := []int{1, 2, 3}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Poll", func(t *testing.T) {

		h := NewMinHeap[int]()
		h.items = []int{10, 15, 20, 17, 25}

		first, _ := h.Poll()
		actual := h.items

		expected := []int{15, 17, 20, 25}

		if first != 10 || !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v, first: %v", actual, expected, first)
		}

	})

	t.Run("Test Add", func(t *testing.T) {

		h := NewMinHeap[int]()
		h.items = []int{10, 15, 20, 17}

		h.Add(8)

		actual := h.items
		expected := []int{8, 10, 20, 17, 15}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
