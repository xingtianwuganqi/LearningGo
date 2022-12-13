package utils

import "fmt"

func init() {
	fmt.Println("init方法被调用")
}

func init() {
	fmt.Println("另一个init函数")
}

func Count() {
	fmt.Println("Count方法被调用")
}
