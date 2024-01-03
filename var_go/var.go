package main

import "fmt"

var x, y int

// 这种分解的写法,一般用于声明全局变量
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "Golang"

//这种不带声明格式的只能在函数体内声明
//g, h := 123, "需要在func函数体内实现"

// var a int = 100
// var b = "abc"

// c := 1.2

func main() {
	// fmt.Println("a = ", a)
	// fmt.Println("b = ", b)
	// fmt.Println("c = ", c)

	//默认值
	// var a int
	// fmt.Printf("a = %d\n", a)
	// fmt.Printf("type of a = %T\n", a)

	// var b int = 10
	// fmt.Printf("b = %d\n", b)
	// fmt.Printf("type of b = %T\n", b)

	// var c = "golang"
	// fmt.Printf("c = %s\n", c)
	// fmt.Printf("type of c = %T\n", c)

	// d := 20.15
	// fmt.Printf("d = %f\n", d)
	// fmt.Printf("type of d = %T\n", d)

	g, h := 333, "需要在函数体内实现"

	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	//不能对g变量再次做初始化声明
	//g := 400

	_, value := 7, 5 ////实际上7的赋值被废弃，变量 _  不具备读特性
	//fmt.Println(_)
	fmt.Println(value)

	var (
		vv int  = 15
		cc bool = false
	)
	fmt.Println(vv, cc)
}
