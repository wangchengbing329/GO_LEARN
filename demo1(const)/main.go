package main

import "fmt"

var age int

// 批量申明
var (
	name   string
	gender int
)

// 初始化多个值
var name1, name2 = "xiaowang", "xiaoming"

func main() {
	age = 10
	// 在函数内部使用短变量声明
	n := 10
	m := 200
	fmt.Println(n, m)
	// 同时声明多个常量
	const (
		pi = 3.141592653
		e  = 2.71828
	)

	// 如果忽略值则代表此变量值与上一个一样
	const (
		n1 = 100
		n2
		n3
	)
	fmt.Println(n1, n2, n3)

	// iota 用于一般用于枚举
	const (
		a1 = iota
		a2
		a3
		a4
		a5
	)
	const n6 = iota
	fmt.Println(a1, a2, a3, a4, a5, n6)

	// 算平常用的kb
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB = 1 << (10 * iota)
		GB = 1 << (10 * iota)
		TB = 1 << (10 * iota)
		PB = 1 << (10 * iota)
	)

	// 多个iota 定义在一行
	const (
		a, b = iota + 1, iota + 2 //1,2
		c,   //2,3
		h //3,4
	)

	fmt.Println(a, b, c, h)

	const (
		q    = iota
		r, s = iota + 1, iota + 2
	)
	fmt.Println(q, r, s)
}
