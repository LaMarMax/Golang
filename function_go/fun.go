package main

import "fmt"

// func swap(x, y string) (string, string) {
// 	return y, x
// }

// func prt(a string, b int) (r1 int, r2 int) {
// 	fmt.Println("a = ", a)
// 	fmt.Println("b = ", b)

// 	fmt.Println(r1, r2)

// 	//给有名称的返回值变量赋值
// 	r1 = 1000
// 	r2 = 2000
// 	return
// }

func swap(x, y *int) {
	var temp int

	temp = *x
	*x = *y
	*y = temp
}

func main() {
	// a, b := swap("golang", "c++")

	//a, b := prt("golang", 123)
	//fmt.Println(a, b)

	a := 100
	// b := 200
	// fmt.Println("交换前a的值：", a)
	// fmt.Println("交换前b的值：", b)

	// swap(&a, &b)
	// fmt.Println("交换后a的值：", a)
	// fmt.Println("交换后b的值：", b)

	var p *int = &a
	var pp **int = &p
	fmt.Println(&p)
	fmt.Println(pp)

	//fmt.Println(&a)
}
