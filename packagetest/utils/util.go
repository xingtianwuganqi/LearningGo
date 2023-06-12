package utils

import "fmt"

func init() {
	fmt.Println("util文件下init方法被调用")
}

func init() {
	fmt.Println("util文件下另一个init函数")
}

func Count() {
	fmt.Println("util文件下Count方法被调用")
}
