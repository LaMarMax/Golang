package struct_tag

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Golang允许向结构体字段后面绑定标签，使用kv格式
// 主要作用是其他的包再倒入当前结构体属性的时候，来判断该属性对其他包的作用，起到说明的作用
type Resume struct {
	Name string `info:"name" doc:"我的名字"`
	Sex  string `info:"sex"`
}

func findTag(input interface{}) {
	//通过reflect来得到标签
	t := reflect.TypeOf(input).Elem() // 得到当前结构体的全部元素
	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info") // 通过标签的key找到标签的value
		tagDoc := t.Field(i).Tag.Get("doc")
		fmt.Println("info:", tagInfo, "doc", tagDoc)
	}
}

// 可以使用encoding/json将数据转换为json文件格式
// 如果需要转换为json文件，则需要给结构体字段打上json标签
type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"price"`
	Actors []string `json:"actors"`
}

func StructTag() {
	var re Resume
	findTag(&re)
}

func ConvertToJson() {
	movie1 := Movie{"让子弹飞", 2009, 50, []string{"姜文，葛优，周润发"}}

	//编码的过程 struct --> json
	jsonStr, err := json.Marshal(movie1)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	}

	fmt.Printf("jsonStr = %s\n", jsonStr)

	//解码的过程 jsonStr --> struct
	my_movie := Movie{}
	err = json.Unmarshal(jsonStr, &my_movie)
	if err != nil {
		fmt.Println("json unmarshal error", err)
		return
	}
	fmt.Println(my_movie)
}
