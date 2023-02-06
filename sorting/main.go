package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{3, 5, 1, 5, 9, 10, -1}))
	fmt.Println(quickSort([]int{3, 5, 1, 5, 9, 10}))
	fmt.Println(quickSort([]int{1, 2, 3, 4, 5}))
	fmt.Println(quickSort([]int{5, 4, 3, 2, 1}))
	fmt.Println(quickSort([]int{1}))
	fmt.Println(quickSort([]int{1, 1, 1, 1, 1}))
	fmt.Println(quickSort([]int{}))
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

func quickSort(a []int) []int {
	res := make([]int, len(a))
	copy(res, a)
	quickSortInner(res, 0, len(res)-1)
	return res
}

func quickSortInner(a []int, left, right int) {
	if right-left < 1 {
		return
	}
	l := left
	r := right
	piv := a[(l+r)/2]
	// fmt.Printf("piv:%v, arr:%v\n", piv, a[left:right])
	for l < r {
		for l < r && a[l] < piv {
			l++
		}
		for l < r && a[r] > piv {
			r--
		}
		if l < r {
			a[l], a[r] = a[r], a[l]
			l++
			r--
		}
	}
	if l == r {
		if a[r] > piv {
			r--
		}
	}
	quickSortInner(a, left, r)
	quickSortInner(a, r+1, right)
}
