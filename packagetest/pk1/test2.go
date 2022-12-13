package pk1

import "fmt"
import "packagetest/pk2"

func init() {
	fmt.Println("pk1下test2的init方法调用")
	pk2.MyTest()
}

func MyTest() {
	fmt.Println("pk1下Mytest方法被调用")
}
