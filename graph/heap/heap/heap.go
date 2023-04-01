package heap

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
Heap is a data structure that gives
- Heapify: O(NlogN) Make a heap from an array
- Insersion: O(logN) Insert a element into the heap
- Deletion: O(logN) Delete a top element of the heap
- Peek: O(1) Check a top element of the heap

Let `i` as the current node, then left child is 2i+1 and right child is 2i+2
*/
type Heap[T constraints.Ordered] interface {
	Push(T)
	Empty() bool
	Pop() (T, error)
	Peek() (T, error)
}

type minHeap[T constraints.Ordered] struct {
	arr  []T
	size int
}

var _ Heap[int] = (*minHeap[int])(nil)

// NewMinHeap makes a new empty minHeap with the given capacity
func NewMinHeap[T constraints.Ordered](cap int) *minHeap[T] {
	arr := make([]T, 0, cap)
	return &minHeap[T]{
		arr:  arr,
		size: 0,
	}
}

func (h *minHeap[T]) heapifyNode(p int) (int, error) {
	if p > h.size/2-1 {
		return 0, fmt.Errorf("heapifyNode: `p` is larger than len(a)/2-1")
	}
	smallest := p
	left := p*2 + 1
	right := p*2 + 2
	if h.arr[left] < h.arr[p] {
		smallest = left
	}
	if right < h.size && h.arr[right] < h.arr[p] {
		smallest = right
	}
	h.arr[p], h.arr[smallest] = h.arr[smallest], h.arr[p]
	return smallest, nil
}

// Heapify makes a new minHeap with the given slice
func Heapify[T constraints.Ordered](a []T) *minHeap[T] {
	h := minHeap[T]{
		arr:  a,
		size: len(a),
	}
	// begin with a node whose index is len/2-1 := the last non-leaf node
	// swap the node with smaller child
	for i := len(a)/2 - 1; i >= 0; i-- {
		_, err := h.heapifyNode(i)
		if err != nil {
			panic(err)
		}
	}
	return &h
}

func (h *minHeap[T]) Push(v T) {
	h.arr = append(h.arr, v)
	h.size++
	for i := h.size/2 - 1; i >= 0; i = (i - 1) / 2 {
		_, err := h.heapifyNode(i)
		if err != nil {
			panic(err)
		}
		if i == 0 {
			break
		}
	}
}

func (h *minHeap[T]) Empty() bool {
	return h.size == 0
}

func (h *minHeap[T]) Pop() (T, error) {
	if h.Empty() {
		return *new(T), fmt.Errorf("Pop: heap is empty")
	}
	res := h.arr[0]
	h.size--
	h.arr[0] = h.arr[h.size]
	h.arr = h.arr[:h.size]
	if h.size <= 1 {
		return res, nil
	}
	for i := 0; i < h.size; {
		nxt, err := h.heapifyNode(i)
		if err != nil {
			panic(err)
		}
		if i == nxt || nxt > h.size/2-1 {
			break
		}
		i = nxt
	}
	return res, nil
}

func (h *minHeap[T]) Peek() (T, error) {
	if h.Empty() {
		return *new(T), fmt.Errorf("Peek: heap is empty")
	}
	return h.arr[0], nil
}
