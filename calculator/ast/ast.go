package ast

import (
	"fmt"
	"strings"
)

/*
The number of terms and operators are same.
Example:
9 - 8 => terms[9, 8] operators[+, -]
*/
type Expression struct {
	Ts  []*Term
	Ops []*OperatorAddLike
}

func (e *Expression) Valid() bool {
	return len(e.Ts) == len(e.Ops) && len(e.Ts) > 0
}

func (e *Expression) String() string {
	if !e.Valid() {
		return fmt.Sprintf("Not valid expression: terms=%v, ops=%v", e.Ts, e.Ops)
	}
	b := strings.Builder{}
	for i := 0; i < len(e.Ts); i++ {
		b.WriteString(fmt.Sprintf("%v%v", e.Ops[i], e.Ts[i]))
	}
	return b.String()
}

/*
The number of factors and operators are same.
The first item of operator must nil.
Example:
9 * 8 => factors[9, 8] operators[nil, *]
*/
type Term struct {
	Fs  []*Factor
	Ops []*OperatorMultLike
}

func (t *Term) Valid() bool {
	return len(t.Fs) == len(t.Ops) && (len(t.Ops) == 0 || t.Ops[0] == nil)
}

func (t *Term) String() string {
	if !t.Valid() {
		return fmt.Sprintf("Not valid term: factors=%v, ops=%v", t.Fs, t.Ops)
	}
	b := strings.Builder{}
	for i := 0; i < len(t.Fs); i++ {
		if i == 0 {
			b.WriteString(fmt.Sprint(t.Fs[i]))
		} else {
			b.WriteString(fmt.Sprintf("%v%v", t.Ops[i], t.Fs[i]))
		}
	}
	return b.String()
}

/*
Factor has one of Number or Expression.
Having both or nothing is invalid.
*/
type Factor struct {
	Num *Number
	Exp *Expression
}

func (t *Factor) Valid() bool {
	return !(t.Num != nil && t.Exp != nil) && !(t.Num == nil && t.Exp == nil)
}

func (t *Factor) String() string {
	if !t.Valid() {
		return fmt.Sprintf("Not valid factor: num=%v, exp=%v", t.Num, t.Exp)
	}
	if t.Num != nil {
		return t.Num.String()
	} else {
		return fmt.Sprintf("(%v)", t.Exp)
	}
}

type OperatorAddLikeKind uint8

const (
	OperatorAddLikeKindUndefined OperatorAddLikeKind = iota
	OperatorAddLikeKindAdd
	OperatorAddLikeKindSub
)

type OperatorAddLike struct {
	Op OperatorAddLikeKind
}

func (o *OperatorAddLike) String() string {
	switch o.Op {
	case OperatorAddLikeKindAdd:
		return "+"
	case OperatorAddLikeKindSub:
		return "-"
	default:
		return "<Undefined>"
	}
}

func ConvertOperatorAddLike(s string) OperatorAddLikeKind {
	switch s {
	case "+":
		return OperatorAddLikeKindAdd
	case "-":
		return OperatorAddLikeKindSub
	default:
		return OperatorAddLikeKindUndefined
	}
}

type OperatorMultLikeKind uint8

const (
	OperatorMultLikeKindUndefined OperatorMultLikeKind = iota
	OperatorMultLikeKindMult
	OperatorMultLikeKindDiv
)

type OperatorMultLike struct {
	Op OperatorMultLikeKind
}

func (o *OperatorMultLike) String() string {
	switch o.Op {
	case OperatorMultLikeKindMult:
		return "*"
	case OperatorMultLikeKindDiv:
		return "-"
	default:
		return "<Undefined>"
	}
}

func ConvertOperatorMultLike(s string) OperatorMultLikeKind {
	switch s {
	case "*":
		return OperatorMultLikeKindMult
	case "/":
		return OperatorMultLikeKindDiv
	default:
		return OperatorMultLikeKindUndefined
	}
}

type Number struct {
	Value int64
}

func (n *Number) String() string {
	return fmt.Sprint(n.Value)
}
