package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func sortTest(t *testing.T, arr []int, sortFunc func([]int) []int) bool {
	a1 := make([]int, len(arr))
	copy(a1, arr)
	got := sortFunc(a1)

	a2 := make([]int, len(arr))
	copy(a2, arr)
	sort.Ints(a2)
	return reflect.DeepEqual(got, a2)
}

func test_sort(t *testing.T, sortFunc func([]int) []int) {
	rand.Seed(15)
	for i := 0; i < 1000; i++ {
		sz := rand.Intn(1000)
		arr := make([]int, sz)
		for j := 0; j < len(arr); j++ {
			arr[j] = rand.Int()
		}
		if !sortTest(t, arr, sortFunc) {
			t.Error("Failed. arr=", arr)
		}
	}
}

func Test_mergeSort(t *testing.T) {
	test_sort(t, mergeSort)
}

func Test_quickSort(t *testing.T) {
	test_sort(t, quickSort)
}
