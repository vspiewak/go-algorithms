package tree

import (
	. "github.com/vspiewak/go-algorithms/constraints"
)

type node[T Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

func newNode[T Ordered](value T) *node[T] {
	return &node[T]{value: value, left: nil, right: nil}
}

type BinarySearchTree[T Ordered] struct {
	root *node[T]
}

func NewBinarySearchTree[T Ordered]() *BinarySearchTree[T] {
	return &BinarySearchTree[T]{}
}

func (bst *BinarySearchTree[T]) IsEmpty() bool {
	return bst.root == nil
}

func (bst *BinarySearchTree[T]) Equal(toCompare *BinarySearchTree[T]) bool {
	if bst.root == nil {
		return toCompare.root == nil
	}
	//return bst.root.equal(toCompare.root)
	return equal[T](bst.root, toCompare.root)
}

func (bst *BinarySearchTree[T]) Insert(value T) {
	if bst.root == nil {
		bst.root = newNode[T](value)
		return
	}
	bst.root.insert(value)
}

func (bst *BinarySearchTree[T]) Contains(value T) bool {
	if bst.root == nil {
		return false
	}
	return bst.root.contains(value)
}

func (bst *BinarySearchTree[T]) ToSliceInOrder() []T {
	if bst.root == nil {
		return []T{}
	}
	return bst.root.toSliceInOrder()
}

func (bst *BinarySearchTree[T]) ToSlicePreOrder() []T {
	if bst.root == nil {
		return []T{}
	}
	return bst.root.toSlicePreOrder()
}

func (bst *BinarySearchTree[T]) ToSlicePostOrder() []T {
	if bst.root == nil {
		return []T{}
	}
	return bst.root.toSlicePostOrder()
}

func (node *node[T]) contains(value T) bool {

	if node.value == value {
		return true
	}

	if value < node.value {
		if node.left == nil {
			return false
		} else {
			return node.left.contains(value)
		}
	} else {
		if node.right == nil {
			return false
		} else {
			return node.right.contains(value)
		}
	}

}

func equal[T Ordered](x *node[T], y *node[T]) bool {
	if x == y {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	if x.value != y.value {
		return false
	}
	return equal[T](x.left, y.left) && equal[T](x.right, y.right)
}

func (node *node[T]) equal(toCompare *node[T]) bool {

	if toCompare == nil {
		return false
	}

	if node.value != toCompare.value {
		return false
	}

	if node.left == nil {
		if toCompare.left != nil {
			return false
		}
	} else {
		if !node.left.equal(toCompare.left) {
			return false
		}
	}

	if node.right == nil {
		if toCompare.right != nil {
			return false
		}
	} else {
		if !node.right.equal(toCompare.right) {
			return false
		}
	}

	return true

}

func (node *node[T]) insert(value T) {
	if value <= node.value {
		if node.left == nil {
			node.left = newNode[T](value)
		} else {
			node.left.insert(value)
		}
	} else {
		if node.right == nil {
			node.right = newNode[T](value)
		} else {
			node.right.insert(value)
		}
	}
}

func (node *node[T]) toSliceInOrder() []T {

	values := []T{}

	if node.left != nil {
		values = append(values, node.left.toSliceInOrder()...)
	}

	values = append(values, node.value)

	if node.right != nil {
		values = append(values, node.right.toSliceInOrder()...)
	}

	return values
}

func (node *node[T]) toSlicePreOrder() []T {

	values := []T{}

	values = append(values, node.value)

	if node.left != nil {
		values = append(values, node.left.toSliceInOrder()...)
	}

	if node.right != nil {
		values = append(values, node.right.toSliceInOrder()...)
	}

	return values
}

func (node *node[T]) toSlicePostOrder() []T {

	values := []T{}

	if node.left != nil {
		values = append(values, node.left.toSliceInOrder()...)
	}

	if node.right != nil {
		values = append(values, node.right.toSliceInOrder()...)
	}

	values = append(values, node.value)

	return values
}
