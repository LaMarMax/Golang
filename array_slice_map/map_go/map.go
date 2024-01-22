package map_go

import "fmt"

func Map() {
	map1 := make(map[int]string)
	map1[1] = "cpp"
	map1[2] = "java"
	map1[3] = "golang"

	fmt.Println(map1)
	fmt.Println("-------------")
	changeMap(map1)
	fmt.Println(map1)

	// language := make(map[string]map[string]string)
	// language["php"] = make(map[string]string, 2)
	// language["php"]["id"] = "1"
	// language["php"]["desc"] = "php是世界上最美的语言"

	// language["golang"] = make(map[string]string, 2)
	// language["golang"]["id"] = "2"
	// language["golang"]["desc"] = "golang抗并发非常good"
	// language["golang"]["name"] = "golang"

	// fmt.Println(language)

	// //增
	// language["cpp"] = make(map[string]string)
	// language["cpp"]["id"] = "3"
	// fmt.Println(language)

	// //删
	// delete(language, "cpp")
	// delete(language["php"], "desc")
	// fmt.Println(language)

	// //查
	// val, key := language["golang"]
	// if key {
	// 	fmt.Println(key, val)
	// } else {
	// 	fmt.Println("no key")
	// }

	// //改
	// language["golang"]["name"] = "GOLANG"
	// fmt.Println(language)

	// var map1 map[string]string
	// //在使用map前，需要先make，make的作用就是给map分配数据空间
	// map1 = make(map[string]string, 10)
	// map1["one"] = "php"
	// map1["two"] = "golang"
	// map1["three"] = "java"
	// fmt.Println(map1)

	// map2 := make(map[string]string)
	// map2["one"] = "php"
	// map2["two"] = "golang"
	// map2["three"] = "java"
	// fmt.Println(map2)

	// map3 := map[string]string{
	// 	"one":   "php",
	// 	"two":   "golang",
	// 	"three": "java", // 每一个键值对后面都需要加,
	// }
	// fmt.Println(map3)
}

func changeMap(myMap map[int]string) {
	myMap[0] = "python"
}
