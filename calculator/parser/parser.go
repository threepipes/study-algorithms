package parser

import (
	"calculator/ast"
	"fmt"
	"strconv"
)

// Parse make a AST of valid BNF by given tokenized tokens.
func Parse(tokens []string) (*ast.Expression, error) {
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
	opk := ast.ConvertOperatorAddLike(tokens[0])
	return opk != ast.OperatorAddLikeKindUndefined
}

func expression(tokens []string) (*ast.Expression, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty expression")
	}
	res := &ast.Expression{
		Ts:  make([]*ast.Term, 0),
		Ops: make([]*ast.OperatorAddLike, 0),
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
		res.Ops = append(res.Ops, &ast.OperatorAddLike{Op: ast.OperatorAddLikeKindAdd})
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
	opk := ast.ConvertOperatorMultLike(tokens[0])
	return opk != ast.OperatorMultLikeKindUndefined
}

func term(tokens []string) (*ast.Term, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty term")
	}
	res := &ast.Term{
		Fs:  make([]*ast.Factor, 0),
		Ops: make([]*ast.OperatorMultLike, 0),
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

func factor(tokens []string) (*ast.Factor, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty factor")
	}
	if tokens[0] != "(" {
		num, next, err := number(tokens)
		if err != nil {
			return nil, nil, fmt.Errorf("factor: %w", err)
		}
		return &ast.Factor{Num: num}, next, err
	}
	exp, next, err := expression(tokens[1:])
	if err != nil {
		return nil, nil, fmt.Errorf("factor: %w", err)
	}
	if len(next) == 0 || next[0] != ")" {
		return nil, nil, fmt.Errorf("factor: no end bracket found")
	}
	return &ast.Factor{Exp: exp}, next[1:], nil
}

func number(tokens []string) (*ast.Number, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty number")
	}
	v, err := strconv.ParseInt(tokens[0], 10, 64)
	if err != nil {
		return nil, nil, fmt.Errorf("number: %w", err)
	}
	return &ast.Number{Value: v}, tokens[1:], nil
}

func opAdd(tokens []string) (*ast.OperatorAddLike, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty add like operator")
	}
	op := ast.ConvertOperatorAddLike(tokens[0])
	if op == ast.OperatorAddLikeKindUndefined {
		return nil, nil, fmt.Errorf("invalid add like operator: %v", tokens[0])
	}
	return &ast.OperatorAddLike{Op: op}, tokens[1:], nil
}

func opMult(tokens []string) (*ast.OperatorMultLike, []string, error) {
	if len(tokens) == 0 {
		return nil, nil, fmt.Errorf("empty mult like operator")
	}
	op := ast.ConvertOperatorMultLike(tokens[0])
	if op == ast.OperatorMultLikeKindUndefined {
		return nil, nil, fmt.Errorf("invalid mutl like operator: %v", tokens[0])
	}
	return &ast.OperatorMultLike{Op: op}, tokens[1:], nil
}
