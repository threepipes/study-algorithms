package main

import (
	"fmt"
	"math/bits"
	"strconv"
)

func makeMask(i, j uint) int {
	return ^0<<j | ^(^0 << i)
}

func InsertBits(n, m int, i, j uint) int {
	if i > j {
		j, i = i, j
	}
	msk := makeMask(i, j)
	return n&msk | m<<i
}

func max(a, b, c int) int {
	res := a
	if res < b {
		res = b
	}
	if res < c {
		res = c
	}
	return res
}

func BestReversedSize(n uint) int {
	var usedSec, unusedSec, mx int
	for i := 0; i < bits.UintSize; i++ {
		if n&1 == 1 {
			usedSec += 1
			unusedSec += 1
		} else {
			usedSec = unusedSec + 1
			unusedSec = 0
		}
		n >>= 1
		mx = max(mx, usedSec, unusedSec)
	}
	return mx
}

func main() {
	fmt.Println(strconv.FormatInt(-1, 2))
	fmt.Println(strconv.FormatInt(-1<<2, 2))
	fmt.Println(strconv.FormatInt(-1<<8, 2))
	fmt.Println(strconv.FormatInt(^(-1 << 3), 2))
}
