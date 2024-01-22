package main

//import "fmt"

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

// func swap(x, y *int) {
// 	var temp int

// 	temp = *x
// 	*x = *y
// 	*y = temp
// }

// func main() {
// 	// a, b := swap("golang", "c++")

// 	//a, b := prt("golang", 123)
// 	//fmt.Println(a, b)

// 	a := 100
// 	// b := 200
// 	// fmt.Println("交换前a的值：", a)
// 	// fmt.Println("交换前b的值：", b)

// 	// swap(&a, &b)
// 	// fmt.Println("交换后a的值：", a)
// 	// fmt.Println("交换后b的值：", b)

// 	var p *int = &a
// 	var pp **int = &p
// 	fmt.Println(&p)
// 	fmt.Println(pp)

// 	//fmt.Println(&a)
// }

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	/* 声明函数变量 */
// 	res := func(x float64) float64 {
// 		return math.Sqrt(x)
// 	}(9)

// 	fmt.Println(res)
// }

// 	/* 使用函数 */
// 	//fmt.Println(getSquareRoot(9))
// }

// import "fmt"

// // getSequence()函数返回值是一个匿名函数
// func getSequence() func() int {
// 	i := 0
// 	//返回一个匿名函数，该函数引用了函数体外部的i变量
// 	//这样i和匿名函数形成一个整体，构成闭包
// 	return func() int {
// 		i += 1
// 		return i
// 	}
// }

// func main() {
// 	// nextNumber 为一个函数，函数中 i 为 0
// 	// nextNumber返回的是一个匿名函数，赋值给变量nextNumber
// 	nextNumber := getSequence()

// 	fmt.Println(nextNumber())
// 	fmt.Println(nextNumber())
// 	fmt.Println(nextNumber())

// 	//再创建一个变量，并赋值getSequence返回的匿名函数
// 	//nextNumber1函数中的i还是从0开始的
// 	nextNumber1 := getSequence()

// 	fmt.Println(nextNumber1())
// 	fmt.Println(nextNumber1())
// 	fmt.Println(nextNumber1())
// }

// import (
// 	"fmt"
// 	"strings"
// )

// func makeSuffix(suffix string) func(string) string {
// 	return func(fileName string) string {
// 		result := fileName
// 		if !strings.HasSuffix(fileName, suffix) {
// 			result = fileName + suffix
// 		}
// 		return result
// 	}
// }

// func main() {
// 	f := makeSuffix(".jpg")
// 	fmt.Println(f("img1.jpg"))
// 	fmt.Println(f("img2"))
// }

import "fmt"

// 声明一个结构体，成员有radius
type Circle struct {
	radius float64
}

func main() {
	//定义一个结构体对象
	var c1 Circle
	c1.radius = 15.0
	fmt.Println("Area of c1 = ", c1.getArea())
}

// getArea()属于Circle这个结构体类型中的方法
func (c Circle) getArea() float64 {
	//可以直接调用Circle中的成员
	return 3.14 * c.radius * c.radius
}
