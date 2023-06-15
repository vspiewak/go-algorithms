package heap

import . "github.com/vspiewak/go-algorithms/constraints"

type MinHeap[T Ordered] struct {
	items []T
}

func NewMinHeap[T Ordered]() *MinHeap[T] {
	return &MinHeap[T]{
		items: []T{},
	}
}

func (h *MinHeap[T]) leftChildIndex(parentIndex int) int {
	return 2*parentIndex + 1
}

func (h *MinHeap[T]) rightChildIndex(parentIndex int) int {
	return 2*parentIndex + 2
}

func (h *MinHeap[T]) parentIndex(childIndex int) int {
	return (childIndex - 1) / 2
}

func (h *MinHeap[T]) hasLeftChild(index int) bool {
	return h.leftChildIndex(index) < len(h.items)
}

func (h *MinHeap[T]) hasRightChild(index int) bool {
	return h.rightChildIndex(index) < len(h.items)
}

func (h *MinHeap[T]) hasParent(index int) bool {
	return index > 0
}

func (h *MinHeap[T]) leftChild(index int) T {
	return h.items[h.leftChildIndex(index)]
}

func (h *MinHeap[T]) rightChild(index int) T {
	return h.items[h.rightChildIndex(index)]
}

func (h *MinHeap[T]) parent(index int) T {
	return h.items[h.parentIndex(index)]
}

func (h *MinHeap[T]) swap(indexOne int, indexTwo int) {
	itemOne := h.items[indexOne]
	h.items[indexOne] = h.items[indexTwo]
	h.items[indexTwo] = itemOne
}

func (h *MinHeap[T]) ToSlice() []T {
	ret := make([]T, len(h.items))
	copy(ret, h.items)
	return ret
}

func (h *MinHeap[T]) Peek() (T, error) {
	var t T
	if len(h.items) == 0 {
		return t, ErrMinHeapEmpty
	}
	return h.items[0], nil
}

func (h *MinHeap[T]) Poll() (T, error) {
	var t T
	if len(h.items) == 0 {
		return t, ErrMinHeapEmpty
	}
	item := h.items[0]
	h.items[0] = h.items[len(h.items)-1]
	h.items = h.items[:len(h.items)-1]
	h.heapifyDown()
	return item, nil
}

func (h *MinHeap[T]) heapifyDown() {
	index := 0
	for h.hasLeftChild(index) {
		smallerChildIndex := h.leftChildIndex(index)
		if h.hasRightChild(index) && h.rightChild(index) < h.leftChild(index) {
			smallerChildIndex = h.rightChildIndex(index)
		}
		if h.items[index] < h.items[smallerChildIndex] {
			break
		} else {
			h.swap(smallerChildIndex, index)
			index = smallerChildIndex
		}
	}

}

func (h *MinHeap[T]) Add(value T) {
	h.items = append(h.items, value)
	h.heapifyUp()
}

func (h *MinHeap[T]) heapifyUp() {
	index := len(h.items) - 1
	for h.hasParent(index) && h.parent(index) > h.items[index] {
		h.swap(h.parentIndex(index), index)
		index = h.parentIndex(index)
	}
}
