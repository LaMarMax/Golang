package main

import (
	"fmt"
	"goroutine_go/channel_go"
	"time"
)

// 子goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Println("new Goroutine : i =", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine
func main() {
	channel_go.SingleChannel()

	// go func(a int, b int) bool {
	// 	fmt.Println("a =", a, "b =", b)
	// 	return true
	// }(10, 20)

	// //用go创建承载一个形参为空，返回值为空的一个函数
	// go func() {
	// 	defer fmt.Println("A.defer")

	// 	func() {
	// 		defer fmt.Println("B.defer")
	// 		//退出当前goroutine
	// 		runtime.Goexit()
	// 		fmt.Println("B")
	// 	}()
	// 	fmt.Println("A")
	// }()

	//死循环
	// for {
	// 	time.Sleep(1 * time.Second)
	// }

	// //创建一个go程，去执行newTask流程
	// go newTask()
	// i := 0
	// for {
	// 	i++
	// 	fmt.Println("main Goroutine : i =", i)
	// 	time.Sleep(1 * time.Second)
	// }
}
