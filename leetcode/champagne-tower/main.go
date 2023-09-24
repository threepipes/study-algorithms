package main

import "fmt"

func main() {
	fmt.Println(champagneTower(100000009, 33, 17))
}

func min(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func champagneTower(poured int, query_row int, query_glass int) float64 {
	gls := make([]float64, 1)
	gls[0] = float64(poured)
	for i := 0; i < query_row; i++ {
		ngls := make([]float64, i+2)
		for j := 0; j <= i; j++ {
			ngls[j] += max((gls[j]-1)/2, 0)
			ngls[j+1] += max((gls[j]-1)/2, 0)
		}
		gls = ngls
	}
	return min(1, gls[query_glass])
}
