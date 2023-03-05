package main

import (
	"fmt"
	"strings"
)

const rawBoard = `xwxx
x...
x.w.
b.b.`

const (
	h = 4
	w = 4

	block = 0
	road  = 1

	OOB = uint8((1 << 8) - 1)
	EOP = ^uint32(0)
)

var (
	dy = []int8{-2, -2, -1, -1, 1, 1, 2, 2}
	dx = []int8{-1, 1, -2, -2, -2, 2, -1, 1}
)

func outOf(mp []uint8, y int8, x int8) bool {
	if y < 0 || y >= h || x < 0 || x >= w {
		return true
	}
	return mp[y*w+x] == block
}

func applyDiff(mp []uint8, idx int8, i int) uint8 {
	y := idx/w + dy[i]
	x := idx%w + dx[i]
	if outOf(mp, y, x) {
		return OOB
	}
	return uint8(y*w + x)
}

const mask = (1 << 5) - 1

func stateToPos(state uint32) (w1, w2, b1, b2 uint8) {
	return uint8(state & mask), uint8((state >> 5) & mask), uint8((state >> 10) & mask), uint8((state >> 15) & mask)
}

func posToState(w1, w2, b1, b2 uint8) uint32 {
	if w1 < w2 {
		w1, w2 = w2, w1
	}
	if b1 < b2 {
		b1, b2 = b2, b1
	}
	return uint32(w1) | (uint32(w2) << 5) | (uint32(b1) << 10) | (uint32(b2) << 15)
}

type Queue struct {
	qu  []uint32
	top int
}

func NewQueue(cap int) *Queue {
	return &Queue{
		qu:  make([]uint32, 0, cap),
		top: 0,
	}
}

func (q *Queue) empty() bool {
	return q.top >= len(q.qu)
}

func (q *Queue) pop() uint32 {
	if q.empty() {
		panic("The queue has no items.")
	}
	res := q.qu[q.top]
	q.top++
	return res
}

func (q *Queue) push(v uint32) {
	q.qu = append(q.qu, v)
}

func main() {
	solve()
}

const maxStateNum = 1 << 20

func solve() bool {
	board := make([]uint8, h*w)
	rows := strings.Split(rawBoard, "\n")
	whiteKts := make([]uint8, 0, 2)
	blackKts := make([]uint8, 0, 2)
	for i, row := range rows {
		for j, c := range row {
			idx := i*w + j
			switch c {
			case 'x':
				board[idx] = block
			case '.':
				board[idx] = road
			case 'w':
				board[idx] = road
				whiteKts = append(whiteKts, uint8(idx))
			case 'b':
				board[idx] = road
				blackKts = append(blackKts, uint8(idx))
			}
		}
	}
	paths := make([][]uint8, h*w)
	for idx, p := range paths {
		if board[idx] == 'x' {
			continue
		}
		for i := 0; i < len(dy); i++ {
			nid := applyDiff(board, int8(idx), i)
			if nid == OOB {
				continue
			}
			p = append(p, nid)
		}
	}
	iniSt := posToState(whiteKts[0], whiteKts[1], blackKts[0], blackKts[1])
	goalSt := posToState(blackKts[0], blackKts[1], whiteKts[0], whiteKts[1])
	// fmt.Printf("inital: %20s\n", strconv.FormatInt(int64(iniSt), 2))
	// fmt.Printf("goal  : %20s\n", strconv.FormatInt(int64(goalSt), 2))
	var prev [maxStateNum]uint32
	prev[iniSt] = EOP
	qu := NewQueue(maxStateNum)
	qu.push(iniSt)
	pos := [4]uint8{}
	for !qu.empty() {
		st := qu.pop()
		pos[0], pos[1], pos[2], pos[3] = stateToPos(st)
		for pid, p := range pos {
		loop:
			for i := 0; i < len(dy); i++ {
				nid := applyDiff(board, int8(p), i)
				for j := 0; j < 4; j++ {
					if pos[j] == nid {
						continue loop
					}
				}
				pos[pid] = nid
				nst := posToState(pos[0], pos[1], pos[2], pos[3])
				if prev[nst] == 0 {
					prev[nst] = st
					if nst == goalSt {
						displayAnswer(board, prev[:], nst)
						return true
					}
					qu.push(nst)
				}
				pos[pid] = p
			}
		}
	}

	return false
}

func boardToStr(board []byte) string {
	bld := strings.Builder{}
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			bld.WriteByte(board[i*w+j])
		}
		bld.WriteByte('\n')
	}
	return bld.String()
}

func stateToStr(board []byte, st uint32) string {
	w1, w2, b1, b2 := stateToPos(st)
	board[w1], board[w2], board[b1], board[b2] = 'w', 'w', 'b', 'b'
	res := boardToStr(board)
	board[w1], board[w2], board[b1], board[b2] = '.', '.', '.', '.'
	return res
}

func displayAnswer(board []uint8, prev []uint32, lastSt uint32) {
	current := lastSt
	ans := make([]uint32, 0)
	for current != EOP {
		ans = append(ans, current)
		current = prev[current]
	}

	bb := make([]byte, len(board))
	for i := 0; i < h*w; i++ {
		switch board[i] {
		case block:
			bb[i] = ' '
		case road:
			bb[i] = '.'
		}
	}
	for _, a := range ans {
		fmt.Println(stateToStr(bb, a))
	}
}
