package slice_go

import "fmt"

func Slice() {
	//arr := [5]int{1, 2, 3}
	s1 := []int{4, 5, 6, 7, 8}

	//s1 := arr[1:]
	// s2 := make([]int, 3)
	s2 := []int{1, 2, 3}
	copy(s2, s1[2:5])
	s2[0] = 100

	printSlice(s1)
	printSlice(s2)

	//fmt.Println(arr)
	//printSlice(s1)

	//arr1 := []int{1, 2, 3, 4}
	// changeSlice(arr1)
	//printSlice(arr1)

	// var numbers []int

	// //允许追加空切片
	// numbers = append(numbers, 0)
	// fmt.Println("len = ", len(numbers), "cap = ", cap(numbers), "numbers = ", numbers)

	// //向切片追加一个元素
	// numbers = append(numbers, 1)
	// fmt.Println("len = ", len(numbers), "cap = ", cap(numbers), "numbers = ", numbers)

	// //同时添加多个元素
	// numbers = append(numbers, 2, 3, 4)
	// fmt.Println("len = ", len(numbers), "cap = ", cap(numbers), "numbers = ", numbers)

	// //创建切片，是numbers的两倍容量
	// numbers1 := make([]int, len(numbers), cap(numbers)*2)
	// //拷贝numbers的数据到numbers1中
	// copy(numbers1, numbers)
	// fmt.Println("len = ", len(numbers1), "cap = ", cap(numbers1), "numbers = ", numbers1)

	// s1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	// //打印原始切片
	// fmt.Println("s1 = ", s1)

	// //打印子切片，从索引1（包含）到索引4（不包含）
	// fmt.Println("s1[1:4]", s1[1:4])

	// //默认下限为0
	// fmt.Println("s1[:5]", s1[:5])

	// //默认上限为len(s1)
	// fmt.Println("s1[4:]", s1[4:])

	//切片定义
	// var slice1 []int
	// var slice2 = make([]int, 5)
	// slice3 := make([]int, 3, 10)

	//切片初始化
	//arr := [3]int{1, 2, 3}
	//s1 := []int{4, 5, 6}
	//s3 := arr[0:2]
	//s4 := arr[1:]
	//s5 := arr[:3]
	//s6 := s1[1:]
	//s7 := make([]int, 5, 10)

	// for i := 0; i < len(s3); i++ {
	// 	fmt.Println(s3[i])
	// }
	// for i := 0; i < len(s6); i++ {
	// 	fmt.Println(s6[i])
	// }

	// var s1 []int

	// fmt.Println("s1 len:", len(s1), "s1 cap:", cap(s1), "s1 = ", s1)

	// if s1 == nil {
	// 	fmt.Println("切片是空的")
	// }
}

func printSlice(arr []int) {
	fmt.Println("len = ", len(arr), "cap = ", cap(arr), "numbers = ", arr)
}

func changeSlice(arr []int) {
	arr[2] = 10
	fmt.Println(arr)
}
