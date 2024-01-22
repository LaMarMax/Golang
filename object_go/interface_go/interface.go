package interface_go

import "fmt"

// 本质是一个父类指针，用于定义接口，底层是指针指向函数列表
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物颜色
	GetType() string  //获取动物类别
}

// 具体的类
type Cat struct {
	//无需在子类中继承interface，只需要实现interface的方法，就可以实现多态
	//这样就可以用interface的指针去指向Cat对象
	color string //猫的颜色
}

// 重写interface接口中的方法
// Cat类重写了interface中的所有方法，就代表该类实现了当前的接口
// 若Cat类没有完全重写interface中的所有方法，则interface指针无法指向该类
func (c *Cat) Sleep() {
	fmt.Println("Cat is sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "Cat"
}

type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("Dog is sleep")
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "Dog"
}

// 使用父类指针来接受子类对象
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("type = ", animal.GetType())
}

// arg这个参数可以接受任意的类型
func allType(arg interface{}) {
	//fmt.Println(arg)

	// //interface{}改如何区分 此时引用的底层数据类型到底是什么
	// value, ok := arg.(string)
	// if ok {
	// 	fmt.Println("arg is string type, value = ", value)
	// } else {
	// 	fmt.Println("arg is not string type")
	// }

	switch t := arg.(type) {
	case string:
		fmt.Println("arg is string type, value = ", t)
	case int:
		fmt.Println("arg is int type, value = ", t)
	case float32:
		fmt.Println("arg is float32 type, value = ", t)
	case float64:
		fmt.Println("arg is float64 type, value = ", t)
	default:
		fmt.Println("arg is unexpected type")
	}
}

type Book struct {
	name   string
	author string
}

func Interface() {
	book1 := Book{"shijie", "zhangsan"}
	allType(book1)
	allType("nihao")
	allType(12.5848)

	// var animal AnimalIF     //接口的数据类型，也就是父类的指针
	// animal = &Cat{"yellow"} //通过匿名对象的指针给接口赋值，
	// animal.Sleep()          //调用的就是Cat的Sleep()方法

	// animal = &Dog{"black"}
	// fmt.Println(animal.GetColor()) //调用的就是Dog的GetColor()方法，多态的现象

	// c1 := Cat{"yellow"}
	// d1 := Dog{"white"}
	// showAnimal(&c1)
	// showAnimal(&d1)
}
