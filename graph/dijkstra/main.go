package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	MaxUint  = ^uint(0)
	MinUint  = 0
	MaxInt   = int(^uint(0) >> 1)
	MinInt   = -(MaxInt - 1)
	MaxInt64 = int64(^uint64(0) >> 1)
	MinInt64 = -(MaxInt64 - 1)
)

func er(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var sc = bufio.NewScanner(os.Stdin)

func nextInt() int {
	sc.Scan()
	ret, err := strconv.Atoi(sc.Text())
	er(err)
	return ret
}

type Edge struct {
	from int
	to   int
	c    int64
}

type Node struct {
	id   int
	dist int64
}

type PriorityQueue []*Node

// Len is the number of elements in the collection.
func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i int, j int) bool {
	return p[i].dist < p[j].dist
}

// Swap swaps the elements with indexes i and j.
func (p PriorityQueue) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x interface{}) {
	*p = append(*p, x.(*Node))
}

func (p *PriorityQueue) Pop() interface{} {
	n := len(*p)
	res := (*p)[n-1]
	(*p)[n-1] = nil
	*p = (*p)[0 : n-1]
	return res
}

func dijkstra(n, m int, es [][]*Edge) []int64 {
	pq := PriorityQueue(make([]*Node, 0))
	d := make([]int64, n)
	for i := 0; i < n; i++ {
		d[i] = MaxInt64
	}

	heap.Push(&pq, &Node{id: 0, dist: 0})
	for len(pq) > 0 {
		v := heap.Pop(&pq).(*Node)
		if d[v.id] <= v.dist {
			continue
		}
		d[v.id] = v.dist
		for _, e := range es[v.id] {
			heap.Push(&pq, &Node{
				id:   e.to,
				dist: v.dist + e.c,
			})
		}
	}

	return d
}

// https://atcoder.jp/contests/typical-algorithm/tasks/typical_algorithm_d
func main() {
	sc.Split(bufio.ScanWords)

	n, m := nextInt(), nextInt()
	es := make([][]*Edge, n)
	for i := 0; i < n; i++ {
		es[i] = make([]*Edge, 0)
	}
	for i := 0; i < m; i++ {
		e := &Edge{
			from: nextInt(),
			to:   nextInt(),
			c:    int64(nextInt()),
		}
		es[e.from] = append(es[e.from], e)
		// er := &Edge{
		// 	from: e.to,
		// 	to:   e.from,
		// 	c:    e.c,
		// }
		// es[er.from] = append(es[er.from], er)
	}
	d := dijkstra(n, m, es)
	fmt.Println(d[n-1])
}
