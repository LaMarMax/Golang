package main

import (
	"fmt"
)

const (
	UNKNOW = 0
	FEMALE = 1
	MALE   = 2
)

// const (
// 	a = "golang"
// 	b = len(a)
// 	c = unsafe.Sizeof(a)
// )

const (
	CategoryBooks    = iota // 0
	CategoryHealth          // 1
	CategoryClothing        // 2
)

type Allergen int

const (
	IgEggs         Allergen = 1 << iota // 1 << 0 which is 00000001
	IgChocolate                         // 1 << 1 which is 00000010
	IgNuts                              // 1 << 2 which is 00000100
	IgStrawberries                      // 1 << 3 which is 00001000
	IgShellfish                         // 1 << 4 which is 00010000
)

type ByteSize float64

const (
	_           = iota             // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota) // 1 << (10 * 1)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	Apple, Banana = iota + 1, iota + 2
	Cherimoya, Durian
	Elderberry, Fig
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, j
)

func main() {
	// const LENGTH int = 10
	// const WIDTH int = 5
	// var area int
	// const a, b, c = 1, false, "golang"

	// area = LENGTH * WIDTH
	// fmt.Println("面积为：", area)

	//fmt.Println(a, b, c)

	//fmt.Println(IgEggs | IgChocolate | IgShellfish)
	//fmt.Println(KB, MB, GB)
	//fmt.Println(Apple, Banana, Cherimoya, Durian)
	fmt.Println("a = ", a, "b = ", b)
	fmt.Println("c = ", c, "d = ", d)
	fmt.Println("e = ", e, "f = ", f)
	fmt.Println("g = ", g, "h = ", h)
	fmt.Println("i = ", i, "j = ", j)
}
