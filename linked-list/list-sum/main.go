package main

import (
	"fmt"
	"strconv"
)

type ListItem struct {
	Val  uint8
	Next *ListItem
}

func New(value string) (*ListItem, error) {
	if len(value) == 0 {
		return nil, nil
	}
	d, err := strconv.Atoi(value[0:1])
	if err != nil {
		return nil, fmt.Errorf("invalid value found: %v", value[0])
	}
	nx, err := New(value[1:])
	if err != nil {
		return nil, err
	}
	res := &ListItem{
		Val:  uint8(d),
		Next: nx,
	}
	return res, nil
}

func (l *ListItem) String() string {
	if l.Next == nil {
		return fmt.Sprintf("%v", l.Val)
	}
	return fmt.Sprintf("%v -> %v", l.Val, l.Next)
}

func (l *ListItem) CalcLen() int {
	if l == nil {
		return 0
	}
	return 1 + l.Next.CalcLen()
}

func padding(l *ListItem, s uint) *ListItem {
	cur := l
	for i := 0; i < int(s); i++ {
		top := &ListItem{
			Val:  0,
			Next: cur,
		}
		cur = top
	}
	return cur
}

/*
- investigate which one is longer
- check how longer the longer one than shorter one => offset
- sum these items with recursive strategy
*/
func sum(a, b *ListItem) (*ListItem, error) {
	al := a.CalcLen()
	bl := b.CalcLen()
	if al < bl {
		a, b = b, a
		al, bl = bl, al
	}
	offset := al - bl
	b = padding(b, uint(offset))
	res, inc, err := sumRec(a, b)
	if inc == 1 {
		top := &ListItem{
			Val:  1,
			Next: res,
		}
		res = top
	} else if inc > 1 {
		return nil, fmt.Errorf("returned increment value isn't valid: %v", inc)
	}
	return res, err
}

func (l *ListItem) addDigit(v uint8) (inc uint8) {
	l.Val += v
	if l.Val >= 10 {
		inc = 1
		l.Val -= 10
	}
	return
}

func sumRec(lng, sht *ListItem) (res *ListItem, inc uint8, err error) {
	res = &ListItem{
		Val: lng.Val,
	}
	inc = res.addDigit(sht.Val)
	if lng.Next == nil || sht.Next == nil {
		if lng.Next != nil || sht.Next != nil {
			return nil, 0, fmt.Errorf("offset value may be wrong")
		}
		return
	}
	lower, lowInc, err := sumRec(lng.Next, sht.Next)
	res.Next = lower
	inc += res.addDigit(lowInc)
	return
}

func testSum(a, b string) string {
	al, err := New(a)
	if err != nil {
		panic(err)
	}
	bl, err := New(b)
	if err != nil {
		panic(err)
	}
	res, err := sum(al, bl)
	if err != nil {
		panic(err)
	}
	return res.String()
}

func main() {
	fmt.Println(testSum("1", "5"))
	fmt.Println(testSum("12345", "56789"))
	fmt.Println(testSum("12", "5"))
	fmt.Println(testSum("12", "345"))
	fmt.Println(testSum("1", "99999"))
}
