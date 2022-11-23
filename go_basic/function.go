package main

import (
	"fmt"
	//"math/rand"
	//"sort"
	//"strconv"
	//"strings"
)

func main() {
	// 函数返回值
	//returnSum()
	//returnSum2()
	//per, area := rectangle(100, 100)
	//fmt.Println(per, area)
	//rectangle2(200, 300)

	// 延迟函数
	//deferTest()

	// 匿名函数
	//func3()

	// 高阶函数，回调函数
	//res1 := oper(10, 20, add)
	//fmt.Println(res1)
	//
	//res2 := oper(20, 20, func(a, b int) int {
	//	return a * b
	//})
	//fmt.Println(res2)

	// 闭包
	//res1 := increment()
	//fmt.Println(res1)
	//v1 := res1()
	//fmt.Println(v1)
	//v2 := res1()
	//fmt.Println(v2)

	//指针数组
	//pointerTest()

	// 指针作为参数
	a := 10
	fmt.Println(a) // 打印10
	pointer1(&a)   // 在内部修改了a内存保存的数值
	fmt.Println(a) // 打印100
}

// 函数
func addFunction(n int) int {
	/*
		概念：具有特定功能的代码，可以被多次调用执行
		1.可以避免重复使用
		2增加函数的扩展性
		main: 程序的入口，是一个特殊的函数
		3.注意事项：
			a.函数必须先定义
			b.函数名不能冲突
			c.main函数系统自动调用

	*/
	//return 10 + 2
	sum := 0
	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	return sum
}

func getAdd(a int, b int) int {
	sun := a + b
	for i := 0; i < 10; i++ {
		sun += i
	}
	return sun
}

func changeParam(nums ...int) int {
	/*
		1.如果函数的参数是可变参数，同时还有其他参数，那么可变参数要放在列表的最后
		2.一个函数的参数列表中最多只能有一个可变参数
	*/
	/*
		参数的传递
		a.值传递：传递的是数据的副本，修改数据，对于原始的数据没有影响，值类型的数据，默认都是值传递，基础类型，array, struct
		b.引用传递：传递的事数据地址，引导多个变量指向同一块内存，所有的引用类型都是引用传递：slice、map、chan
	*/
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
	}
	return sum
}

// 定义函数时，直接指定要返回的对象
func returnSum() (sum int) {
	//for i := 0; i < 100; i++ {
	//	sum += 1
	//}
	fmt.Println(sum)
	return
}

func returnSum2() (sum int) {
	fmt.Println(sum) // sum有默认值
	for i := 0; i < 100; i++ {
		sum += 1
	}
	fmt.Println(sum)
	return // 代表返回sum
}

// 返回多个值
func returnParam() (string, int) {
	return "hello", 3
}

/*
return语句:
一个函数的定义有返回值，那么函数中必须使用return语句，将结果返回给调用处
函数的返回的结果，必须和函数定义的一致：类型，个数，顺序

1.将函数的结果返回给调用处
2.同时结束了改函数的执行

空白标识符，专门用户舍弃数据：_

注意点：
1.一个函数定义了返回值，必须使用return语句将结果返回给调用处，return后的数据必须和函数定义的一致：个数，类型，顺序
2.可以使用_舍弃多余的返回值
3.如果一个函数定义了有返回值，那么函数中有分、循环，那么要保证无论执行了哪个分支，都要有return语句被执行
4.如果函数没有返回值，也可以使用return结束函数
*/

func rectangle(len, wid float64) (float64, float64) {
	perimeter := (len + wid) * 2
	area := len * wid
	return perimeter, area
}

func rectangle2(len, wid float64) (perimeter float64, area float64) {
	perimeter = (len + wid) * 2
	area = len * wid
	return
}

// defer 函数
/*
延迟函数，在go 语言中使用defer关键字来延迟一个函数或者方法的执行
1.defer函数或者方法：一个函数或方法的执行被延迟

2.defer的用法：
A: 对象.close(),临时文件的删除
文件.open()
defer close()

B：go语言中关于异常的处理，使用pannic()和recover()
panic()用户引发恐慌，导致程序中断执行
recover() 函数用于恢复程序的执行，recover() 语法上要求必须在defer中执行。

3.多个defer，按照栈的形式，先进后出，先延迟的后执行，后延迟的先执行

4.defer 函数传递参数的时候：defer函数调用时，就已经传递了参数数据，只是暂时不执行函数中的代码而已。

5.defer 函数注意点：
defer函数：
在外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
当执行外围函数中的return语句时，只有其中所有的寒池函数都执行完毕后，外围函数才真正返回。
当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数执行完毕后，该运行时恐慌才会真正被扩展至调用函数。
*/

func deferTest() {
	//defer func1("hello")
	//fmt.Println("111")
	//defer func1("world")
	//fmt.Println("王二狗")

	a := 1
	fmt.Println(a)
	defer func2(a)
	a++
	fmt.Println(a)
}

func func1(s string) {
	fmt.Println(s)
}

func func2(a int) {
	fmt.Println(a)

}

/*
函数的类型：func(参数列表的数据类型)(返回值列表的数据类型)
*/
/*
函数的本质：函数也是一个对象，函数在go中也是复合类型，可以看作特殊的变量
函数名()：将函数进行调用，函数中的代码会全部执行，然后将return的结果返回给调用处
函数名：指向函数体的内存地址
*/
/*
匿名函数：没有名字的函数
定义一个匿名函数，直接进行调用，通常只能使用一次。
也可以使用匿名函数赋值给某个函数变量，那么就可以调用多次了。

匿名函数的用途：
go是支持函数式编程的
1.将匿名函数作为另一个函数的参数，回调函数
2.将匿名函数作为另一个函数的返回值，可以形成闭包函数。
*/

func func3() {
	// 匿名函数
	func() {
		fmt.Println("我是一个匿名函数")
	}()

	f2 := func() {
		fmt.Println("我是一个多次调用的匿名函数")
	}
	f2()

	// 匿名函数参数
	func(a, b int) {
		fmt.Println(a, b)
	}(1, 2)

	// 定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(1, 3)
	fmt.Println(res1)

	// 定义带返回值的匿名函数
	res2 := func(a, b int) int {
		return a + b
	}
	fmt.Println(res2(100, 200))
}

/*
高阶函数：根据go语言的数据类型特点，可以将一个函数作为另一个函数的参数

fun1(),fun2()
将fun1() 函数作为fun2这个函数的参数
fun2函数：就叫做高阶函数，接受了一个函数作为参数的函数，高阶函数

fun1函数：回调函数
作为另一个函数的参数的函数，叫做回调函数

*/

// 设计一个函数，用于求两个整数的加减乘除运算
func add(a, b int) int {
	return a + b
}

func oper(a, b int, fun func(int, int) int) int {
	fmt.Println(a, b, fun)
	res := fun(a, b)
	return res
}

/*
闭包：
go语言支持函数式编程：
支持将一个函数作为另一个函数的参数
也支持将一个函数作为另一个函数的返回值

闭包：
一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量，
（外层函数中的参数，或者外层函数中直接定义的变量）并且该外层函数的返回值就是
这个内存函数。 这个内存函数和外层函数的局部变量，统称为闭包结构。

局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
但是闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续调用。

*/
func increment() func() int {
	i := 0
	fun := func() int {
		i++
		return i
	}
	return fun
}

/*
指针：pointer
存储了另一个变量的内存地址的变量

数组指针：首先是一个指针，一个数组的地址
*[4]int

指针数组：首先是一个数组，存储的数据类型是指针
[4]*Type
*/
func pointerTest() {
	a := 10
	p1 := &a  // p1 存储的a的地址
	p2 := &p1 // p2 存储的p1的地址
	b := *p1  // 取p1的地址对应的值
	fmt.Println(p2, b)

}

/*
函数指针：一个指针，指向了一个函数的指针，因为go语言中，function，默认看做一个指针，没有*。
slice,map,function 都是引用类型
指针函数：一个函数，该函数的返回值是一个指针
*/

func pointerFun() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	return &arr
}

/*
指针作为参数
值传递：在指针作为参数后，在方法中修改会修改掉原参数的值
值类型在方法中传递是copy，不会修改原类型的数值
引用传递：不需要传递指针，引用类型在方法中被修改，会改变原类型的数值
*/

func pointer1(num *int) {
	fmt.Println(*num)
	*num = 100
	fmt.Println(*num)
}
