package calculator

import (
	"calculator/ast"
	mt "calculator/math"
	"fmt"
)

type Result struct {
	Num   int64
	Denom int64
	Sign  int64
}

func NewResult(v int64) *Result {
	return &Result{
		Num:   v,
		Denom: 1,
		Sign:  1,
	}
}

func (r *Result) String() string {
	sig := ""
	if r.Sign < 0 {
		sig = "- "
	}
	if r.Denom != 1 {
		return fmt.Sprintf("%v%v / %v", sig, r.Num, r.Denom)
	} else {
		return fmt.Sprintf("%v%v", sig, r.Num)
	}
}

func (r *Result) Valid() bool {
	return !(r.Num != 0 && r.Denom == 0) && r.Sign*r.Sign == 1
}

func (r *Result) normalize() {
	if r.Num < 0 {
		r.Sign *= -1
		r.Num *= -1
	}
	if r.Denom < 0 {
		r.Sign *= -1
		r.Denom *= -1
	}
	// TODO: validation
	l := mt.LCD(r.Num, r.Denom)
	if l > 1 {
		r.Num /= l
		r.Denom /= l
	}
}

func (r *Result) AddTo(a *Result) {
	r.normalize()
	a.normalize()
	lcd := mt.LCD(r.Denom, a.Denom)
	rMul := a.Denom / lcd
	aMul := r.Denom / lcd
	r.Num = r.Sign*(r.Num*rMul) + a.Sign*(a.Num*aMul) // may cause overflow
	r.Denom = rMul * r.Denom                          // may cause overflow
	r.Sign = 1
}

func (r *Result) MultTo(a *Result) {
	r.normalize()
	a.normalize()
	r.Num *= a.Num
	r.Denom *= a.Denom
	r.Sign *= a.Sign
}

func (r *Result) ReverseSign() {
	r.Sign *= -1
}

func (r *Result) DivTo(a *Result) {
	r.normalize()
	a.normalize()
	r.Num *= a.Denom
	r.Denom *= a.Num
	r.Sign *= a.Sign
}

func (r *Result) AsFloat64() (float64, error) {
	if !r.Valid() {
		return 0, fmt.Errorf("result is invalid: %v", r)
	}
	return float64(r.Num) / float64(r.Denom) * float64(r.Sign), nil
}

func Calculate(e *ast.Expression) (*Result, error) {
	r, err := expression(e)
	if err != nil {
		return nil, fmt.Errorf("calculation error: %w", err)
	}
	r.normalize()
	return r, nil
}

func expression(e *ast.Expression) (*Result, error) {
	if !e.Valid() {
		return nil, fmt.Errorf("invalid expression found: %v", e)
	}
	r, err := term(e.Ts[0])
	if err != nil {
		return nil, fmt.Errorf("expression: %w", err)
	}
	if e.Ops[0].Op == ast.OperatorAddLikeKindSub {
		r.ReverseSign()
	}
	for i := 1; i < len(e.Ts); i++ {
		tr, err := term(e.Ts[i])
		if err != nil {
			return nil, fmt.Errorf("expression(term[%v]): %w", i, err)
		}
		if e.Ops[i].Op == ast.OperatorAddLikeKindSub {
			tr.ReverseSign()
		}
		r.AddTo(tr)
	}
	return r, nil
}

func term(t *ast.Term) (*Result, error) {
	if !t.Valid() {
		return nil, fmt.Errorf("invalid term found: %v", t)
	}
	r, err := factor(t.Fs[0])
	if err != nil {
		return nil, fmt.Errorf("term: %w", err)
	}
	for i := 1; i < len(t.Fs); i++ {
		fr, err := factor(t.Fs[i])
		if err != nil {
			return nil, fmt.Errorf("term(factor[%v]): %w", i, err)
		}
		if t.Ops[i].Op == ast.OperatorMultLikeKindMult {
			r.MultTo(fr)
		} else if t.Ops[i].Op == ast.OperatorMultLikeKindDiv {
			r.DivTo(fr)
		} else {
			return nil, fmt.Errorf("term(factor[%v]): invalid operator found", i)
		}
	}
	return r, nil
}

func factor(f *ast.Factor) (*Result, error) {
	if !f.Valid() {
		return nil, fmt.Errorf("invalid factor found: %v", f)
	}
	if f.Exp != nil {
		return expression(f.Exp)
	} else {
		return number(f.Num)
	}
}

func number(n *ast.Number) (*Result, error) {
	return NewResult(n.Value), nil
}
