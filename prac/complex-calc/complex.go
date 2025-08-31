package complex

import (
	"fmt"
	"math"
)

// ======================================================================
// 1. The Abstract Contract: Complex Interface
// ======================================================================

// Complex 接口是 SICP 通用操作的直接体现。
// 它定义了一个抽象的复数，以及它必须支持的所有行为。
// 任何实现了这些方法的类型，都可以被当作一个 Complex 来使用。
type Complex interface {
	Real() float64      // 获取实部
	Imag() float64      // 获取虚部
	Magnitude() float64 // 获取模长
	Angle() float64     // 获取幅角 (弧度)

	Add(other Complex) Complex // 加法
	Mul(other Complex) Complex // 乘法
	String() string             // 字符串表示 (实现 fmt.Stringer)
}

// ======================================================================
// 2. Concrete Implementation 1: Rectangular (直角坐标)
// ======================================================================

// Rectangular 结构体是第一种具体的数据表示法
type Rectangular struct {
	real, imag float64
}

// NewRectangular 是 Rectangular 类型的构造函数
func NewRectangular(r, i float64) Complex {
	return &Rectangular{real: r, imag: i}
}

// --- 为 *Rectangular 实现 Complex 接口 ---

func (r *Rectangular) Real() float64      { return r.real }
func (r *Rectangular) Imag() float64      { return r.imag }
func (r *Rectangular) Magnitude() float64 { return math.Sqrt(r.real*r.real + r.imag*r.imag) }
func (r *Rectangular) Angle() float64     { return math.Atan2(r.imag, r.real) }
func (r *Rectangular) String() string     { return fmt.Sprintf("(%f + %fi)", r.real, r.imag) }

// Add 方法。它接收的是一个抽象的 Complex 接口。
// 它不关心 other 的具体类型，只通过接口调用其方法。
func (r *Rectangular) Add(other Complex) Complex {
	// 加法在直角坐标系下最简单
	return NewRectangular(r.Real()+other.Real(), r.Imag()+other.Imag())
}

// Mul 方法
func (r *Rectangular) Mul(other Complex) Complex {
	// (a+bi)(c+di) = (ac-bd) + (ad+bc)i
	real := r.Real()*other.Real() - r.Imag()*other.Imag()
	imag := r.Real()*other.Imag() + r.Imag()*other.Real()
	return NewRectangular(real, imag)
}

// ======================================================================
// 3. Concrete Implementation 2: Polar (极坐标)
// ======================================================================

// Polar 结构体是第二种具体的数据表示法
type Polar struct {
	magnitude, angle float64
}

// NewPolar 是 Polar 类型的构造函数
func NewPolar(m, a float64) Complex {
	return &Polar{magnitude: m, angle: a}
}

// --- 为 *Polar 实现 Complex 接口 ---

func (p *Polar) Real() float64      { return p.magnitude * math.Cos(p.angle) }
func (p *Polar) Imag() float64      { return p.magnitude * math.Sin(p.angle) }
func (p *Polar) Magnitude() float64 { return p.magnitude }
func (p *Polar) Angle() float64     { return p.angle }
func (p *Polar) String() string     { return fmt.Sprintf("(%f∠%f rad)", p.magnitude, p.angle) }

// Add 方法
func (p *Polar) Add(other Complex) Complex {
	// 加法在极坐标下很复杂，但在直角坐标下很简单。
	// 所以我们利用接口，将两者都转换为直角坐标进行计算。
	// 这完美体现了 SICP 的思想：为操作选择最合适的表示法。
	real := p.Real() + other.Real()
	imag := p.Imag() + other.Imag()
	return NewRectangular(real, imag)
}

// Mul 方法
func (p *Polar) Mul(other Complex) Complex {
	// 乘法在极坐标下非常简单：模长相乘，幅角相加。
	return NewPolar(p.Magnitude()*other.Magnitude(), p.Angle()+other.Angle())
}

