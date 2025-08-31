package rat

import "fmt"


type Rational struct {
	n, d int
}

// 1. 构造
func MakeRat(n, d int) Rational {
	return  Rational{n, d};
}

// 2. 选择
func (r Rational) Numer() int {
	return r.n;
}

func (r Rational) Denom() int {
	return r.d;
}

// 3. 加减乘除
func (r Rational) Add(other Rational) (Rational) {
	numer := r.Numer()*other.Denom() + other.Numer()*r.Denom() // 修正
	denom := r.Denom() * other.Denom()
	return MakeRat(numer, denom)
}

func (r Rational) Sub(other Rational) (Rational) {
	numer := r.Numer()*other.Denom() - other.Numer()*r.Denom() // 修正
	denom := r.Denom() * other.Denom()
	return MakeRat(numer, denom)
}

func (r Rational) Mul(other Rational) (Rational) {
	numer := r.Numer() * other.Numer()
	denom := r.Denom() * other.Denom()
	return MakeRat(numer, denom)
}
// 4. print
func (r Rational) PrintRat() {
	fmt.Printf("%d / %d\n", r.n, r.d);
}


