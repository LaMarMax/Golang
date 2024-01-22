package reflect_go

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Sex  string
	Age  int
}

func (usr User) ReflectCallFuncHasArgs(name string, age int) {
	fmt.Println("ReflectCallFuncHasArgs name:", name, "age", age)
}

func (usr User) ReflectCallFuncNoArgs() {
	fmt.Println("ReflectCallFuncNoArgs")
}

func Reflect() {
	//User1 := User{"zhangsan", "male", 16}
	//doFiledAndMethod(User1)

	//var num float64 = 12.4485
	//changeValue(User1)

	//callFunc(User1)
	convert()
}

func doFiledAndMethod(input interface{}) {
	//获取type
	inputType := reflect.TypeOf(input)
	fmt.Println("input type: ", inputType)
	//获取value
	inputValue := reflect.ValueOf(input)
	fmt.Println("input value: ", inputValue)
	//通过type获取里面的字段
	//1.获取interface的reflect.Type，通过Type得到NumField，进行遍历
	//2.得到每个field，数据类型
	//3.通过field中有一个Interface()方法，得到对应的value
	for i := 0; i < inputType.NumField(); i++ {
		field := inputType.Field(i)              //取出第i个字段
		value := inputValue.Field(i).Interface() //取出第i个字段的值

		fmt.Println(field.Name, field.Type, value)
	}

	//通过type获取里面的方法
	for i := 0; i < inputType.NumMethod(); i++ {
		m := inputType.Method(i)
		fmt.Println(m.Name, m.Type)
	}
}

func changeValue(input interface{}) {
	// 通过reflect.ValueOf获取num中的reflect.Value，注意，参数必须是指针才能修改其值
	pointer := reflect.ValueOf(&input)
	//Elem返回v持有的接口保管的值的Value封装，或者v持有的指针指向的值的Value封装。
	//如果v的Kind不是Interface或Ptr会panic；如果v持有的值为nil，会返回Value零值。
	newValue := pointer.Elem()

	fmt.Println("pointer: type of pointer: ", pointer.Type())
	fmt.Println("pointer: settability of pointer: ", pointer.CanSet())

	fmt.Println("newValue: type of pointer: ", newValue.Type())
	fmt.Println("newValue: settability of pointer: ", newValue.CanSet())

	//重新赋值
	u := User{"lisi", "male", 25}
	val := reflect.ValueOf(u)
	newValue.Set(val)

	fmt.Println("new value: ", input)
}

func callFunc(input interface{}) {
	//获取Value对象
	getValue := reflect.ValueOf(input)
	getType := reflect.TypeOf(input)

	//获取方法
	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Println("name:", method.Name, "type:", method.Type)
	}

	//有参调用
	//获取方法,一定要指定参数为正确的方法名
	method1 := getValue.MethodByName("ReflectCallFuncHasArgs")
	//参数列表
	args := []reflect.Value{reflect.ValueOf("xiaozhang"), reflect.ValueOf(18)}
	//调用
	method1.Call(args)

	//无参调用
	method2 := getValue.MethodByName("ReflectCallFuncNoArgs")
	//无参，也要传入空的args
	args = make([]reflect.Value, 0)
	//或者 args2 = []reflect.Value{}
	method2.Call(args)
}

func convert() {
	var num float64 = 1.2345

	pointer := reflect.ValueOf(&num)
	value := reflect.ValueOf(num)

	// 可以理解为“强制转换”，但是需要注意的时候，转换的时候，如果转换的类型不完全符合，则直接panic
	// Golang 对类型要求非常严格，类型一定要完全符合
	// 如下两个，一个是*float64，一个是float64，如果弄混，则会panic
	convertPointer := pointer.Interface().(*float64)
	convertValue := value.Interface().(float64)
	convertInterface := value.Interface()

	fmt.Printf("%T\n", convertPointer)
	fmt.Println(convertPointer)

	fmt.Printf("%T\n", convertValue)
	fmt.Println(convertValue)

	val, ok := convertInterface.(float64)
	if ok {
		fmt.Printf("convertInterface is %T type\n", val)
	}
}
