package main

import "fmt"

func main() {
	fmt.Println(mergeSort([]int{3, 5, 1, 5, 9, 10, -1}))
	fmt.Println(mergeSort([]int{3, 5, 1, 5, 9, 10}))
}

func mergeSort(a []int) []int {
	// split the array into two parts
	if len(a) <= 1 {
		return a
	}
	// call mergeSort recursively to sort the former array
	fm := mergeSort(a[:len(a)/2])
	// to sort the latter
	lt := mergeSort(a[len(a)/2:])

	// merge two arrays into one
	res := make([]int, len(a))
	var fi, li int
	for i := 0; i < len(a); i++ {
		if fi >= len(fm) {
			res[i] = lt[li]
			li++
		} else if li >= len(lt) {
			res[i] = fm[fi]
			fi++
		} else if fm[fi] < lt[li] {
			res[i] = fm[fi]
			fi++
		} else {
			res[i] = lt[li]
			li++
		}
	}

	// return the result
	return res
}
