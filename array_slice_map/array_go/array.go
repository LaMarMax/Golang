package array_go

import "fmt"

func Array() {
	var arr1 [10]int
	// arr2 := []int{1, 2, 3}

	// fmt.Println(arr1[0])

	// for i := 0; i < len(arr2); i++ {
	// 	fmt.Println(arr2[i])
	// }

	// for i, num := range arr1 {
	// 	fmt.Println(i, num)
	// }

	printArr(arr1)
}

func printArr(myArr [10]int) {
	for i, num := range myArr {
		fmt.Println(i, num)
	}
}
