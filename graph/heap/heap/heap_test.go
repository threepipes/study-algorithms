package heap_test

import (
	"graph/heap/heap"
	"reflect"
	"testing"
)

func ignoreErr[T any](t T, err error) T {
	return t
}

func TestHeap(t *testing.T) {
	tests := []struct {
		name     string
		commands func() []int
		want     []int
	}{
		{
			name: "normal case",
			commands: func() []int {
				res := make([]int, 0)
				h := heap.NewMinHeap[int](5)
				h.Push(6)
				h.Push(4)
				h.Push(8)
				res = append(res, ignoreErr(h.Pop()))
				h.Push(-1)
				res = append(res, ignoreErr(h.Pop()))
				h.Push(10)
				res = append(res, ignoreErr(h.Pop()))
				res = append(res, ignoreErr(h.Pop()))
				res = append(res, ignoreErr(h.Pop()))
				return res
			},
			want: []int{4, -1, 6, 8, 10},
		},
		{
			name: "same values",
			commands: func() []int {
				res := make([]int, 0)
				h := heap.NewMinHeap[int](5)
				h.Push(7)
				h.Push(7)
				h.Push(7)
				res = append(res, ignoreErr(h.Pop()))
				h.Push(7)
				h.Push(-1)
				h.Push(7)
				res = append(res, ignoreErr(h.Pop()))
				h.Push(7)
				h.Push(10)
				res = append(res, ignoreErr(h.Pop()))
				res = append(res, ignoreErr(h.Pop()))
				res = append(res, ignoreErr(h.Pop()))
				return res
			},
			want: []int{7, -1, 7, 7, 7},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.commands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEmpty(t *testing.T) {
	h := heap.Heapify([]float32{1.0, 0.2, 4.0})
	res := make([]float32, 0)
	want := []float32{0.2, 1.0, 4.0}
	for !h.Empty() {
		res = append(res, ignoreErr(h.Pop()))
	}
	if !reflect.DeepEqual(res, want) {
		t.Errorf("result = %v, want %v", res, want)
	}
	_, err := h.Pop()
	if err == nil {
		t.Errorf("Pop() didn't raise an error even though the heap is empty")
	}
}
