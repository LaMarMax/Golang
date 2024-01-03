package InitLib2

import "fmt"

// lib2提供的API
func Lib2Test() {
	fmt.Println("lib2Test()...")
}

func init() {
	fmt.Println("lib2")
}
