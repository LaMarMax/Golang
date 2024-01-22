package Defer

import "fmt"

func Defer() {
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")

	fmt.Println("0")
	fmt.Println("1")
}

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func caller...")
	return 0
}

func ReturnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func Recover(i int) {
	//定义10个元素的数组
	var arr [10]int
	//错误拦截要在产生错误前设置
	//这里使用匿名函数进行错误拦截，在进行defer调用
	//相当于这个匿名函数在Recover的生命周期结束后才被调用
	defer func() {
		//设置recover拦截错误信息
		err := recover()
		//产生panic异常，打印错误信息
		if err != nil {
			fmt.Println(err)
		}
	}()
	//根据函数参数为数组元素赋值
	//如果i的值超过数组下标 会报错误：数组下标越界
	arr[i] = 2
}
