package nibblelinx

import (
	"math/big"
)

var P, Gx, Gy, NOrder *big.Int
var A, B *big.Int

var point, pointNULL [2]*big.Int

func Init() {
	P, _ = new(big.Int).SetString("115792089237316195423570985008687907853269984665640564039457584007908834671663", 10)
	NOrder, _ = new(big.Int).SetString("115792089237316195423570985008687907852837564279074904382605163141518161494337", 10)
	Gx, _ = new(big.Int).SetString("55066263022277343669578718895168534326250603453777594175500187360389116729240", 10)
	Gy, _ = new(big.Int).SetString("32670510020758816978083085130507043184471273380659243275938904335757337482424", 10)
	A, _ = new(big.Int).SetString("0", 10)
	B, _ = new(big.Int).SetString("7", 10)

	pointNULL[0] = new(big.Int).SetInt64(0)
	pointNULL[1] = new(big.Int).SetInt64(0)
}

func ModP(n, p1 *big.Int) *big.Int {
	n = n.Rem(n, p1)
	if n.Sign() < 0 {
		return p1.Add(p1, n)
	}
	return n
}

func Inverse(r, p *big.Int) *big.Int {
	// Operações feitas em big.Int são salvas no objeto de referência, K salvas alterações nele mesmo sem causar efetos colaterais
	var (
		k    = big.NewInt(0)
		t    = big.NewInt(1)
		aux  = big.NewInt(0)
		newR = big.NewInt(0)
		newT = big.NewInt(0)
	)
	aux.Set(p)
	newR.Set(r)

	newT.Set(k.Sub(big.NewInt(0), k.Div(aux, r)))

	r.Set(newR)
	newR.Set(k.Rem(aux, r))
	for newR.Sign() != 0 {
		aux.Set(t)

		//t.Set(newT)
		newT.Set(k.Sub(aux, k.Mul(t.Set(newT), k.Div(r, newR))))
		aux.Set(r)

		r.Set(newR)
		newR.Set(k.Rem(aux, newR))
	}

	if t.Sign() < 0 {
		return k.Add(t, p)
	}
	return t
}

func DoubleP(x, y *big.Int) [2]*big.Int {
	var (
		k  = new(big.Int)
		m  = new(big.Int)
		a  = new(big.Int)
		b  = new(big.Int)
		i3 = new(big.Int)
		i1 = new(big.Int)
		i2 = new(big.Int)
		p0 = new(big.Int)
		p1 = new(big.Int)
	)
	a.Set(Inverse(k.Mul(y, big.NewInt(2)), P))
	b.Set(k.Mul(x, k.Mul(x, big.NewInt(3))))

	m.Set(ModP(k.Mul(a, k.Add(b, A)), P))

	a.Set(k.Mul(m, m))
	b.Set(k.Mul(x, big.NewInt(2)))

	p0.Set(ModP(k.Sub(a, b), P))

	i3.Set(k.Mul(m, k.Mul(x, big.NewInt(2))))
	i1.Set(k.Sub(k.Mul(ModP(k.Mul(m, m), P), m), i3))
	i2.Set(k.Sub(y, k.Mul(m, x)))

	p1.Set(ModP(k.Sub(big.NewInt(0), k.Add(i1, i2)), P))

	return [2]*big.Int{p0, p1}
}
