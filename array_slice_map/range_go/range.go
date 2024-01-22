package range_go

import "fmt"

func Range() {
	arr1 := [5]int{1, 2, 3, 4}
	for i, num := range arr1 {
		fmt.Println(i, num)
	}

	for _, num := range arr1 {
		fmt.Println(num)
	}

	slice1 := arr1[:3]
	for i, num := range slice1 {
		fmt.Println(i, num)
	}

	map1 := make(map[int]string)
	map1[0] = "xiaoming"
	map1[1] = "xiaogang"
	map1[2] = "xiangzhao"
	for key, value := range map1 {
		fmt.Println(key, value)
	}
}
