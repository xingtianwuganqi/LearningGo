package main

import "fmt"

func main() {

	// 方法
	methodTest()
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
