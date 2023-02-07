package main

import (
	"fmt"
)

type Operator uint8

const (
	OperatorUnknown Operator = iota
	OperatorAnd
	OperatorOr
	OperatorXor
)

type Expression struct {
	vals []bool
	ops  []Operator
}

func main() {
	fmt.Println(solve("1^0|0|1", false))
	fmt.Println(solve("0&0&0&1^1|0", true))
}

func solve(s string, expectedResult bool) (uint64, error) {
	ex, err := tokenize(s)
	if err != nil {
		return 0, fmt.Errorf("failed to tokenize: %v", err)
	}
	cache = make([][]*ResultCount, len(ex.vals))
	for i := 0; i < len(ex.vals); i++ {
		cache[i] = make([]*ResultCount, len(ex.vals))
	}
	res, err := count(ex, 0, len(ex.ops))
	if err != nil {
		return 0, fmt.Errorf("failed to count: %v", err)
	}
	if expectedResult {
		return res.trueCount, nil
	} else {
		return res.falseCount, nil
	}
}

type ResultCount struct {
	falseCount uint64
	trueCount  uint64
}

var cache [][]*ResultCount

func calcCases(l, r *ResultCount, op Operator) (fls uint64, trs uint64) {
	switch op {
	case OperatorAnd:
		fls = l.falseCount*r.falseCount + l.falseCount*r.trueCount + l.trueCount*r.falseCount
		trs = l.trueCount * r.trueCount
	case OperatorOr:
		fls = l.falseCount * r.falseCount
		trs = l.falseCount*r.trueCount + l.trueCount*r.trueCount + l.trueCount*r.falseCount
	case OperatorXor:
		fls = l.falseCount*r.falseCount + l.trueCount*r.trueCount
		trs = l.falseCount*r.trueCount + l.trueCount*r.trueCount
	}
	return
}

func count(ex Expression, left int, right int) (*ResultCount, error) {
	if left > right {
		return nil, fmt.Errorf("wrong argument: left=%v, right=%v", left, right)
	}
	if cache[left][right] != nil {
		return cache[left][right], nil
	}
	if left == right {
		if ex.vals[left] {
			cache[left][right] = &ResultCount{0, 1}
		} else {
			cache[left][right] = &ResultCount{1, 0}
		}
		return cache[left][right], nil
	}
	res := &ResultCount{0, 0}
	cache[left][right] = res
	for i := left; i < right; i++ {
		cntL, err := count(ex, left, i)
		if err != nil {
			return nil, err
		}
		cntR, err := count(ex, i+1, right)
		if err != nil {
			return nil, err
		}
		fls, trs := calcCases(cntL, cntR, ex.ops[i])
		res.falseCount += fls
		res.trueCount += trs
	}
	return cache[left][right], nil
}

func toBool(c rune) (bool, error) {
	if c == '1' {
		return true, nil
	} else if c == '0' {
		return false, nil
	}
	return false, fmt.Errorf("to bool: wrong value %v", c)
}

func toOperator(c rune) (Operator, error) {
	switch c {
	case '&':
		return OperatorAnd, nil
	case '|':
		return OperatorOr, nil
	case '^':
		return OperatorXor, nil
	}
	return OperatorUnknown, fmt.Errorf("to operator: wrong value %v", c)
}

func tokenize(s string) (ex Expression, err error) {
	vals := make([]bool, (len(s)+1)/2)
	ops := make([]Operator, (len(s)-1)/2)
	ex = Expression{
		vals: vals,
		ops:  ops,
	}
	for i, c := range s {
		if i%2 == 0 {
			vals[i/2], err = toBool(c)
			if err != nil {
				return
			}
		} else {
			ops[i/2], err = toOperator(c)
			if err != nil {
				return
			}
		}
	}
	return
}
