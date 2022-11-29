package main

import "fmt"

func main() {

	// 方法
	//methodTest()

	// 接口
	//interfaceType()

	// 接口嵌套
	manyInterface()
}

/*
方法：method
一个方法就是一个包含了接受者的函数，接受者可以是命名类型或者结构体类型的一个值或者是一个指针
所有给定类型的方法属于该类型的方法集

语法：
func (接受者) 方法名(参数列表) {返回值列表} {

}

总结：method，同函数类似，区别需要有接受者。（也就是调用者），接受者可以在方法内访问
对比函数：
A：意义
方法：某个类型的行为功能，需要指定的接受者调用
函数：一个独立功能的代码，可以直接调用
B：语法
方法：方法名可以相同，只能接受这不同
函数：命名不能冲突
*/

type Worker1 struct {
	name string
	age  int
	sex  string
}

type Cat struct {
	name string
	age  int
}

func methodTest() {
	w1 := Worker1{"哈哈", 18, "女"}
	w1.work()

	// 可以使用指针调用
	w2 := &Worker1{"张三", 17, "男"}
	w2.work()
	w2.rest()
	w1.rest()

	// 方法名可相同
	c1 := Cat{"mini", 3}
	c1.rest()
}

func (w Worker1) work() {
	fmt.Println(w.name, "在工作")
}

func (w *Worker1) rest() {
	fmt.Println(w.name, "在休息")
}

func (c Cat) rest() {
	fmt.Println(c.name, "在休息")
}

/*
接口：
面向对象世界中的接口一般定义为"接口定义对象的行为"。它表示让指定对象应该做什么，
实现这种行为的方法（实现细节）是针对对象的。

在go语言中，接口是一组方法签名，当类型为接口中的所有方法提供定义时，他被称为实现接口，
它与OOP非常相似。接口指定了类型应该具有的方法，类型决定了如何实现这些方法。
*/

/*
接口：interface
在go中，接口是一组方法签名
当某个类型为这个接口中的所有提供了方法的实现，它被称为实现接口。

go语言中，接口和类型的实现关系，是非侵入式

1.当需要接口类型的对象时，可以使用任意实现类对象代替
2.接口对象不能访问实现类中的属性

多态：一个事物的多种形态
go 通过接口模拟多态

就一个接口的实现：
1.看成实现本身的类型，能够访问实现类中的属性和方法
2.看成是对应的接口类型，那就只能够访问接口中的方法

接口的用法：
1.一个函数如果接受接口类型作为参数，
那么实际上可以传入改接口的任意实现类型对象作为参数

2.定义一个类型为接口类型，实际上可以赋值为任意实现类的对象

鸭子类型：

空接口：interface{}
不包含任何的方法，正因为如此，所有的类型都能实现这个空接口，
因此空接口可以存储任意类型的数值

*/

type USB interface {
	start()
	end()
}

type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标，准备就绪，可以开始工作")
}

func (m Mouse) end() {
	fmt.Println(m.name, "鼠标结束工作")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "优盘开始行动")
}

func (f FlashDisk) end() {
	fmt.Println(f.name, "优盘结束行动")
}

func interfaceTest(usb USB) {
	usb.start()
	usb.end()
}

func interfaceType() {
	m1 := Mouse{"罗技小黑"}
	fmt.Println(m1.name)

	f1 := FlashDisk{"闪存"}
	fmt.Println(f1.name)

	interfaceTest(m1)
	interfaceTest(f1)

	var usb USB
	usb = m1
	usb.start()

	var d1 Dog = Dog{"小黑"}
	fmt.Println(d1.name)

	// 空接口
	emptyInterface(d1)
	emptyInterface("3.14")
	emptyInterface(6)

	/*
		空接口的使用
	*/
	// map，key字符串，value任意类型
	map1 := make(map[string]interface{})
	map1["name"] = "李晓华"
	map1["age"] = 30
	map1["friend"] = Mouse{name: "王二狗"}
	fmt.Println(map1)

	// 切片，存储任意类型的数据
	slice1 := make([]interface{}, 0, 10)
	slice1 = append(slice1, map1, 100, d1)
	fmt.Println(slice1)
}

type A interface {
}

type Dog struct {
	name string
}

// 接口A是空接口，理解为代表了任意类型
func emptyInterface(a A) {
	fmt.Println(a)
}

// 也可以直接这样写
func emptyTest(a interface{}) {
	fmt.Println(a)
}

/*
接口的嵌套，接口多继承
*/

type B interface {
	test1()
}

type C interface {
	test2()
}

type D interface {
	B
	C
	test3()
}

type InterTest struct { // 如果想实现接口D，那接口B，C中的方法都要实现
	name string
}

func (c InterTest) test1() {
	fmt.Println("test1")
}

func (c InterTest) test2() {
	fmt.Println("test2")
}

func (c InterTest) test3() {
	fmt.Println("test3")
}

func manyInterface() {
	var cat InterTest = InterTest{}
	cat.test1()
	cat.test2()
	cat.test3()

	fmt.Println("-----------")
	var b1 B = cat
	b1.test1()
	fmt.Println("-----------")
	var c1 C = cat
	c1.test2()
	fmt.Println("-----------")
	var d1 D = cat
	d1.test1()
	d1.test2()
	d1.test3()
}
