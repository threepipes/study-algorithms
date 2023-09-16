package number

type Num struct {
	n int64
	d int64
}

func gcd(a, b int64) int64 {
	if a < b {
		return gcd(b, a)
	}
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func sign(a int64) int64 {
	if a < 0 {
		return -1
	}
	return 1
}

func New(i int64) Num {
	return Num{n: i, d: 1}
}

func (n *Num) normalize() {
	sn := sign(n.n)
	sd := sign(n.d)
	n.n *= sn
	n.d *= sd
	g := gcd(n.n, n.d)
	n.n /= g
	n.d /= g
	n.n *= sn * sd
}

func (n Num) Add(m Num) Num {
	r := Num{
		n: n.n*m.d + m.n*n.d,
		d: n.d * m.d,
	}
	r.normalize()
	return r
}

func (n Num) Sub(m Num) Num {
	r := Num{
		n: n.n*m.d - m.n*n.d,
		d: n.d * m.d,
	}
	r.normalize()
	return r
}

func (n Num) Mul(m Num) Num {
	r := Num{
		n: n.n * m.n,
		d: n.d * m.d,
	}
	r.normalize()
	return r
}

func (n Num) Div(m Num) Num {
	r := Num{
		n: n.n * m.d,
		d: n.d * m.n,
	}
	r.normalize()
	return r
}
