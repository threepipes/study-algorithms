package main

type Item struct {
	c int64
}

type PriorityQueue []*Item

// Len is the number of elements in the collection.
func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i int, j int) bool {
	return p[i].c < p[j].c
}

// Swap swaps the elements with indexes i and j.
func (p PriorityQueue) Swap(i int, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *PriorityQueue) Push(x any) {
	*p = append(*p, x.(*Item))
}

func (p *PriorityQueue) Pop() any {
	n := len(*p)
	res := (*p)[n-1]
	(*p)[n-1] = nil
	*p = (*p)[0 : n-1]
	return res
}
