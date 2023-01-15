package main

import (
	"container/heap"
	"reflect"
	"testing"
)

func assertEq(t *testing.T, got, expected any, mes string) {
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Failed: got=%v, expected=%v; %s", got, expected, mes)
	}
}

func TestPriorityQueue(t *testing.T) {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &Item{4})
	heap.Push(&pq, &Item{9})
	assertEq(t, len(pq), 2, "len 1")
	assertEq(t, heap.Pop(&pq).(*Item).c, int64(4), "pop 1")

	heap.Push(&pq, &Item{-1})
	assertEq(t, heap.Pop(&pq).(*Item).c, int64(-1), "pop 2")
	assertEq(t, heap.Pop(&pq).(*Item).c, int64(9), "pop 3")
	assertEq(t, len(pq), 0, "len 2")

}
