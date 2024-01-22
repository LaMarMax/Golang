package main

import (
	"fmt"
	"object_go/interface_go"
)

// 定义一个结构体
type Book struct {
	title  string
	author string
}

type Hero struct {
	Name  string
	Ad    int
	Level int
}

func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) SetName(newName string) {
	hero.Name = newName
}

type Human struct {
	name string
	sex  string
}

func (human *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (human *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type Superman struct {
	Human //表示Superman类继承了Human类的方法

	level int
}

// 重定义父类方法
func (superman *Superman) Eat() {
	fmt.Println("Superman.Eat()...")
}

// 定义子类的新方法
func (superman *Superman) Fly() {
	fmt.Println("Superman.Fly()...")
}

func main() {
	interface_go.Interface()

	// h1 := Human{name: "zhangsan", sex: "male"}
	// h1.Eat()
	// h1.Walk()

	// //定义一个子类对象
	// s1 := Superman{Human{"蝙蝠侠", "male"}, 3}
	// s1.Eat()  //子类重写的成员函数
	// s1.Walk() //父类成员函数
	// s1.Fly()  //子类独有成员函数

	// var s2 Superman
	// s2.name = "钢铁侠"
	// s2.sex = "male"
	// s2.level = 4

	// hero1 := Hero{Name: "zhangsan", Ad: 25, Level: 3}
	// fmt.Println(hero1)

	// name1 := hero1.GetName()
	// fmt.Println(name1)

	// hero1.SetName("超人")
	// fmt.Println(hero1)

	//可以先定义后初始化
	// var book1 Book
	// book1.title = "Golang"
	// book1.author = "李四"

	// //也可以直接在{}中初始化
	// book2 := Book{title: "c++", author: "王五"}

	// fmt.Println(book1)
	// fmt.Println(book2)

	// changeBook(&book1)
	// fmt.Println(book1)
}

func changeBook(book *Book) {
	book.author = "张三"
}
