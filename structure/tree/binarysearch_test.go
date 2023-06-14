package tree

import (
	"reflect"
	"testing"
)

func TestBinarySearchTree(t *testing.T) {

	/**
	 *            10
	 *         /      \
	 *        5        15
	 *         \     /    \
	 *          8   12    18
	 */

	t.Run("Test Equal", func(t *testing.T) {

		actual := NewBinarySearchTree[int]()
		expected := NewBinarySearchTree[int]()

		if !actual.Equal(expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

		actual.root = newNode[int](10)
		expected.root = newNode[int](10)

		if !actual.Equal(expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

		actual.root.left = newNode[int](5)
		expected.root.left = newNode[int](5)

		if !actual.Equal(expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

		actual.root.right = newNode[int](15)
		expected.root.right = newNode[int](15)

		if !actual.Equal(expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

		actual.root.right.left = newNode[int](12)
		expected.root.right.left = newNode[int](12)

		if !actual.Equal(expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

		actual.root.right.right = newNode[int](18)
		expected.root.right.right = newNode[int](18)

		if !actual.Equal(expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Insert", func(t *testing.T) {

		actual := NewBinarySearchTree[int]()
		actual.Insert(10)
		actual.Insert(5)
		actual.Insert(15)
		actual.Insert(8)
		actual.Insert(12)
		actual.Insert(18)

		expected := NewBinarySearchTree[int]()
		expected.root = &node[int]{value: 10}
		expected.root.left = &node[int]{value: 5}
		expected.root.right = &node[int]{value: 15}
		expected.root.left.right = &node[int]{value: 8}
		expected.root.right.left = &node[int]{value: 12}
		expected.root.right.right = &node[int]{value: 18}

		if !actual.Equal(expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test Contains", func(t *testing.T) {

		bst := NewBinarySearchTree[int]()
		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(15)
		bst.Insert(8)
		bst.Insert(12)
		bst.Insert(18)

		actual := []bool{
			bst.Contains(1), bst.Contains(2), bst.Contains(3), bst.Contains(4), bst.Contains(5),
			bst.Contains(6), bst.Contains(7), bst.Contains(8), bst.Contains(9), bst.Contains(10),
			bst.Contains(11), bst.Contains(12), bst.Contains(13), bst.Contains(14), bst.Contains(15),
			bst.Contains(16), bst.Contains(17), bst.Contains(18), bst.Contains(19), bst.Contains(20),
		}
		expected := []bool{
			false, false, false, false, true,
			false, false, true, false, true,
			false, true, false, false, true,
			false, false, true, false, false,
		}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test ToSliceInOrder", func(t *testing.T) {

		bst := NewBinarySearchTree[int]()
		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(15)
		bst.Insert(8)
		bst.Insert(12)
		bst.Insert(18)

		actual := bst.ToSliceInOrder()
		expected := []int{5, 8, 10, 12, 15, 18}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test ToSlicePreOrder", func(t *testing.T) {

		bst := NewBinarySearchTree[int]()
		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(15)
		bst.Insert(8)
		bst.Insert(12)
		bst.Insert(18)

		actual := bst.ToSlicePreOrder()
		expected := []int{10, 5, 8, 12, 15, 18}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

	t.Run("Test ToSlicePostOrder", func(t *testing.T) {

		bst := NewBinarySearchTree[int]()
		bst.Insert(10)
		bst.Insert(5)
		bst.Insert(15)
		bst.Insert(8)
		bst.Insert(12)
		bst.Insert(18)

		actual := bst.ToSlicePostOrder()
		expected := []int{5, 8, 12, 15, 18, 10}

		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("actual: %v, expected: %v", actual, expected)
		}

	})

}
