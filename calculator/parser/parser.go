package parser

import (
	"fmt"
	"strconv"
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
	return len(e.Ts) == len(e.Ops)
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

func convertOperatorAddLike(s string) OperatorAddLikeKind {
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

func convertOperatorMultLike(s string) OperatorMultLikeKind {
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

// Parse make a AST of valid BNF by given tokenized tokens.
func Parse(tokens []string) (*Expression, error) {
	exp, rest, err := expression(tokens)
	if err != nil {
		return nil, fmt.Errorf("failed to parse: %w", err)
	}
	if len(rest) != 0 {
		return exp, fmt.Errorf("there are some unparsed tokens are left: %v", rest)
	}
	return exp, nil
}

func isTopOperatorAddLike(tokens []string) bool {
	if len(tokens) == 0 {
		return false
	}
	opk := convertOperatorAddLike(tokens[0])
	return opk != OperatorAddLikeKindUndefined
}

func expression(tokens []string) (*Expression, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty expression")
	}
	res := &Expression{
		Ts:  make([]*Term, 0),
		Ops: make([]*OperatorAddLike, 0),
	}

	// Process the first operator if exists
	if isTopOperatorAddLike(tokens) {
		op, next, err := opAdd(tokens)
		if err != nil {
			return nil, nil, fmt.Errorf("expression: %w", err)
		}
		tokens = next
		res.Ops = append(res.Ops, op)
	} else {
		res.Ops = append(res.Ops, &OperatorAddLike{OperatorAddLikeKindAdd})
	}

	// At least one term exists
	t, next, err := term(tokens)
	if err != nil {
		return nil, nil, fmt.Errorf("expression: %w", err)
	}
	tokens = next
	res.Ts = append(res.Ts, t)

	for isTopOperatorAddLike(tokens) {
		op, next, err := opAdd(tokens)
		if err != nil {
			return nil, nil, fmt.Errorf("expression: %w", err)
		}
		tokens = next
		res.Ops = append(res.Ops, op)

		t, next, err := term(tokens)
		if err != nil {
			return nil, nil, fmt.Errorf("expression: %w", err)
		}
		tokens = next
		res.Ts = append(res.Ts, t)
	}
	return res, tokens, nil
}

func isTopOperatorMultLike(tokens []string) bool {
	if len(tokens) == 0 {
		return false
	}
	opk := convertOperatorMultLike(tokens[0])
	return opk != OperatorMultLikeKindUndefined
}

func term(tokens []string) (*Term, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty term")
	}
	res := &Term{
		Fs:  make([]*Factor, 0),
		Ops: make([]*OperatorMultLike, 0),
	}
	res.Ops = append(res.Ops, nil)

	// At least one factor exists
	t, next, err := factor(tokens)
	if err != nil {
		return nil, nil, fmt.Errorf("term: %w", err)
	}
	tokens = next
	res.Fs = append(res.Fs, t)

	for isTopOperatorMultLike(tokens) {
		op, next, err := opMult(tokens)
		if err != nil {
			return nil, nil, fmt.Errorf("term: %w", err)
		}
		tokens = next
		res.Ops = append(res.Ops, op)

		t, next, err := factor(tokens)
		if err != nil {
			return nil, nil, fmt.Errorf("term: %w", err)
		}
		tokens = next
		res.Fs = append(res.Fs, t)
	}
	return res, tokens, nil
}

func factor(tokens []string) (*Factor, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty factor")
	}
	if tokens[0] != "(" {
		num, next, err := number(tokens)
		if err != nil {
			return nil, nil, fmt.Errorf("factor: %w", err)
		}
		return &Factor{Num: num}, next, err
	}
	exp, next, err := expression(tokens[1:])
	if err != nil {
		return nil, nil, fmt.Errorf("factor: %w", err)
	}
	if len(next) == 0 || next[0] != ")" {
		return nil, nil, fmt.Errorf("factor: no end bracket found")
	}
	return &Factor{Exp: exp}, next[1:], nil
}

func number(tokens []string) (*Number, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty number")
	}
	v, err := strconv.ParseInt(tokens[0], 10, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("number: %w", err)
	}
	return &Number{v}, tokens[1:], nil
}

func opAdd(tokens []string) (*OperatorAddLike, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty add like operator")
	}
	op := convertOperatorAddLike(tokens[0])
	if op == OperatorAddLikeKindUndefined {
		return nil, nil, fmt.Errorf("invalid add like operator: %v", tokens[0])
	}
	return &OperatorAddLike{op}, tokens[1:], nil
}

func opMult(tokens []string) (*OperatorMultLike, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty mult like operator")
	}
	op := convertOperatorMultLike(tokens[0])
	if op == OperatorMultLikeKindUndefined {
		return nil, nil, fmt.Errorf("invalid mutl like operator: %v", tokens[0])
	}
	return &OperatorMultLike{op}, tokens[1:], nil
}
