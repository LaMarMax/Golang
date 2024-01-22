package main

// import (
// 	"expression/Defer"
// 	_ "expression/For"
// 	_ "expression/Select"
// 	_ "expression/Switch"
// 	"fmt"
// )

// func main() {
// 	Defer.Recover(10)
// 	fmt.Println("程序继续执行...")
// }

import "fmt"

func handlePanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered:", r)
	}
}

func performTask() {
	defer handlePanic()

	fmt.Println("Performing some task...")
	panic("Oops! Something went wrong!")
	fmt.Println("Task completed.")
}

func main() {
	performTask()
	fmt.Println("Main function continues.")
}
