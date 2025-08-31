package main

// 接口
type Number interface {
	Add(other Number) Number
	Mul(other Number) Number
	String() string
}




// complex