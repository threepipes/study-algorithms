package main

import (
	"fmt"
	"maketen/number"
	"strconv"
)

type set struct {
	v map[number.Num]string
}

func newSet() set {
	v := make(map[number.Num]string)
	return set{v}
}

func (s set) add(a number.Num, anc string) {
	s.v[a] = anc
}

func (s set) merge(l set, r set) {
	for kl, sl := range l.v {
		for kr, sr := range r.v {
			s.add(kl.Add(kr), fmt.Sprintf("(%s + %s)", sl, sr))
			s.add(kl.Sub(kr), fmt.Sprintf("(%s - %s)", sl, sr))
			s.add(kl.Mul(kr), fmt.Sprintf("(%s * %s)", sl, sr))
			s.add(kl.Div(kr), fmt.Sprintf("(%s / %s)", sl, sr))
		}
	}
}

func calc(ns []int64, ans int64) string {
	nl := len(ns)
	space := make([][]set, nl)
	for i := 0; i < nl; i++ {
		space[i] = make([]set, nl-i)
		for j := 0; j < nl-i; j++ {
			space[i][j] = newSet()
		}
	}
	for i, sp := range space[0] {
		sp.add(number.New(ns[i]), strconv.FormatInt(ns[i], 10))
	}
	for i := 1; i < nl; i++ {
		spl := i + 1
		for offset, sp := range space[i] {
			for left := 1; left < spl; left++ {
				right := spl - left
				lsp := space[left-1][offset]
				rsp := space[right-1][offset+left]
				sp.merge(lsp, rsp)
			}
		}
	}
	ansn := number.New(ans)
	return space[nl-1][0].v[ansn]
}

func main() {
	fmt.Println("13 =", calc([]int64{8, 3, 6, 7, 13, 7, 2, 11}, 13))
	fmt.Println("12 =", calc([]int64{8, 3, 6, 13, 7, 2, 11}, 12))
	fmt.Println("11 =", calc([]int64{3, 6, 13, 9, 7, 5}, 12))
}
