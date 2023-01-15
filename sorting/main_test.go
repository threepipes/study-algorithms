package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func test_mergeSort(t *testing.T, arr []int) bool {
	a1 := make([]int, len(arr))
	copy(a1, arr)
	got := mergeSort(a1)

	a2 := make([]int, len(arr))
	copy(a2, arr)
	sort.Ints(a2)
	return reflect.DeepEqual(got, a2)
}

func Test_mergeSort(t *testing.T) {
	rand.Seed(15)
	for i := 0; i < 1000; i++ {
		sz := rand.Intn(100)
		arr := make([]int, sz)
		for j := 0; j < len(arr); j++ {
			arr[j] = rand.Int()
		}
		if !test_mergeSort(t, arr) {
			t.Error("Failed. arr=", arr)
		}
	}
}
