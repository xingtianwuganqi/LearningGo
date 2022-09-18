package main // 同文件下包名应该一致

import (
	"fmt"
)

// import(
// 	. "fmt" // 不需要写名
// 	f "fmt" // 别名
// )

func main() {
	fmt.Println("hello,world!")

	// 变量
	//variable()

	// 常量
	constactValue()
}

// go的源码文件 bin src pkg

// go的常用命令
/*
go run 运行命令源码文件
go build 用于测试编译，不会产生结果
go install 编译并安装代码包或者源文件
go get 从远程代码库下载并安装代码库
go get -x 可以看到过程
go list
*/

/*
命名：
1.以字母开头或下划线开头
2.区分大小写
3.包里大写开头相当于public，小写相当于private

包命名：包名应该用小写，不使用下划线和混合大小写
文件命名：小写单词，使用下划线分隔各个单词
结构体命名：驼峰命名法
接口命名：
变量命名：
常量：大写，用下划线分隔
*/

// 变量
func variable() {
	var name string = "张三"
	fmt.Println(name)

	var num int
	num = 10
	fmt.Println(num)

	num = 20
	fmt.Println(num)

	// 类型推断
	var num2 = 16
	fmt.Printf("类型%T, 数值是%d", num2, num2)

	// 简短定义
	sum := 100
	fmt.Println(sum)

	// 多个变量同时赋值
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	// 类型推断
	var d, e, f = 3, 4.4, "go"
	fmt.Println(d, e, f)

	// 集合类型
	var (
		student = "李晓华"
		age     = 18
		sex     = "女"
	)
	fmt.Println(student, age, sex)

	// 变量的内存分析和注意事项
	/*
		1.变量需先定义才能使用
		2.变量类型和赋值的必须一致
		3.同一个作用域内，变量名不能冲突
		4.简短定义方式，左边的变量至少有一个是新的
		5.简短定义方式，不能定义全局变量
		6.变量的零值：默认值
			整型：0
			浮点型：0
			字符：""
			切片：[]
			字典：map[]
		7.不使用的变量会报错

	*/
	var num3 int
	num3 = 100
	fmt.Printf("num的数值是%d,地址是%p\n", num3, &num3)
	//num的数值是100,地址是0x140000a8020

	num3 = 200
	fmt.Printf("num的数值是%d,地址是%p\n", num3, &num3)
	//num的数值是100,地址是0x140000a8020

	//num, name := 1000, "李晓华"
	//fmt.Println(num, name)
	//./helloworld.go:97:12: no new variables on left side of :=

	num, name, tag := 1000, "李晓华", 2
	fmt.Println(num, name, tag)

	var m int //整数：默认是0
	fmt.Println(m)

	var s1 string
	fmt.Println(s1)

	var s2 []int // 默认是空
	fmt.Println(s2)

	var s3 map[string]string
	fmt.Println(s3)
}

func constactValue() {
	const test string = "abc"
	// 隐式定义：
	//const b = "abc"

	const PI = 3.14
	fmt.Println(test, PI)

	const a, b, c = 1, false, "str"
	fmt.Println(a, b, c)

	const (
		MALE    = 0
		FEMALE  = 1
		UNKNOWN = 3
	)

	// 一组常量中，如果某个常量没有初始值，默认和上一行一致
	const (
		x int = 100
		y
		z string = "go"
	)
	fmt.Println(x, y, z)

	// 枚举类型：使用常量组作为枚举类型，一组相关数值的数据
	const (
		SPRING = 1
		SUMMER = 2
		AUTUMN = 3
		WINTER = 4
	)

	/*
		// 常量只可以是布尔型，数字型（整数型，浮点型和复数）和字符串型
		// 不曾使用的常量，在编译的时候，是不会报错的
		// 显示指定类型的时候，必须确保常量左右值类型一致，需要时可做类型转换，
		这与变量不同，变量可以是不同的类型值
	*/
}
