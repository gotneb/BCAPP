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
	var k = big.NewInt(0)

	n.Set(k.Rem(n, p1))
	if n.Sign() < 0 {
		return k.Add(p1, n)
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
	// k, _a, _b, _c are "dumb" variables. Only made for divide big expressions into small blocks for my own understanding
	var (
		k  = new(big.Int)
		_a = new(big.Int)
		_b = new(big.Int)
		_c = new(big.Int)
		m  = new(big.Int)
		p0 = new(big.Int)
		p1 = new(big.Int)
	)
	//m = modp (((3*(x*x) + A)*inverse ((2*y), p)), p);
	m.Set(ModP(k.Mul(Inverse(k.Mul(y, big.NewInt(2)), P), k.Add(k.Mul(x, k.Mul(x, big.NewInt(3))), A)), P))

	//point[0] = modp ((((m*m)) - (2*x)), p);
	_a.Set(k.Mul(m, m))
	_b.Set(k.Mul(x, big.NewInt(2)))
	p0.Set(ModP(k.Sub(_a, k.Mul(x, big.NewInt(2))), P))

	_a.Set(k.Mul(m, k.Mul(x, big.NewInt(2))))
	_b.Set(k.Sub(k.Mul(m, ModP(k.Mul(m, m), P)), _a))
	_c.Set(k.Sub(y, k.Mul(m, x)))

	//point[1] = modp ((-( (m * modp ((m*m), p)) - (m*(2*x)) + (y - (m*x)))), p);
	p1.Set(ModP(k.Sub(big.NewInt(0), k.Add(_b, _c)), P))

	point = [2]*big.Int{p0, p1}
	return point
}
