package main

import "fmt"

func main() {
	// 结构体
	//strText()

	// 结构体的数据类型
	//strDataType()

	//匿名结构体
	//noNameStr()

	// 模拟面向对象
	oopTest()
}

/*
结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合
结构体成员是由一系列的成员变量构成，这些成员变量也被称为字段
*/

type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func strText() {
	var p1 Person
	fmt.Println(p1)
	p1.name = "李"
	p1.age = 10
	p1.sex = "男"
	p1.address = "北京"
	fmt.Println(p1)

	p2 := Person{}
	p2.name = "张"
	p2.age = 9
	p2.sex = "女"
	fmt.Println(p2)

	p3 := Person{name: "赵", age: 19, sex: "男"}
	fmt.Println(p3)

	p4 := Person{"刘", 19, "男", "北京"}
	fmt.Println(p4)
}

/*
数据类型：
值类型：string, int, float, bool, array, struct
引用类型： slice, map, function, pointer

通过指针：
new(), 创建的对象不是nil，而是空指针
指向了新分配的类型的内存空间，里面存储的零值

// 使用内置函数new(), go语言中国专门用于创建某种类型的指针的函数
	// make只能创建slice，map， channel，并返回一个有初始值（非零）的T类型。
*/

func strDataType() {
	p1 := Person{"王二狗", 20, "男", "上海"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1)
	p2 := p1
	fmt.Println(p2)
	fmt.Printf("%p,%T\n", &p2, p2)

	//2. 定义结构体指针
	var pp1 *Person
	pp1 = &p1
	fmt.Println(pp1)
	fmt.Printf("%p, %T\n", pp1, pp1)
	fmt.Println(*pp1)

	//(*pp1).name = "王五"
	pp1.name = "网吧"
	fmt.Println(pp1)
	fmt.Println(p1)

	// 使用内置函数new(), go语言中国专门用于创建某种类型的指针的函数
	// make只能创建slice，map， channel，并返回一个有初始值（非零）的T类型。
	pp2 := new(Person) // 创建一个结构体指针
	pp2.sex = "女"
	fmt.Println(pp2)
}

/*
匿名结构体
没有名字的结构体，在创建匿名结构体时，同时创建对象
s2 := struct {
		name string
		age  int
	}{"李四", 20}

匿名字段：
type Worker struct {
	string
	int
}
默认使用数据类型作为名字，通过s2.string 访问值
匿名字段的类型不能重复，否则会冲突
*/

func noNameStr() {
	s1 := Student{name: "张三", age: 30}
	fmt.Println(s1.name, s1.age)

	s2 := struct {
		name string
		age  int
	}{"李四", 20}
	fmt.Println(s2.name)

	w2 := Worker{"李晓华", 20}
	fmt.Println(w2.string, w2.int)
}

type Worker struct {
	string
	int
}

type Student struct {
	name string
	age  int
}

/*
go 语言中的oop（面向对象）
go语言的结构体嵌套
1.模拟的继承性：is - a
type A struct {
	field
}

type B struct {
	A // 匿名字段
}
2.模拟的聚合关系： has - a
type C struct {
	field
}

type D struct {
	c C // 字段名字，字段类型
}
*/

// 1.定义父类

type Person1 struct {
	name string
	age  int
}

// 2. 定义子类

type Student1 struct {
	Person1
	school string
}

/*
方法的模拟继承
*/
// 父类的方法
func (p Person1) eat() {
	fmt.Println(p.name, "吃东西")
}

func (s Student1) book() {
	fmt.Println(s.name, "书本")
}

func (s Student1) eat() {
	fmt.Println(s.name, "吃了点东西")
}

func oopTest() {
	// 1.创建父类的对象
	p1 := Person1{name: "张三", age: 30}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)

	// 2.创建子类对象
	s1 := Student1{Person1{"李四", 17}, "school"}
	fmt.Println(s1)

	s2 := Student1{Person1: Person1{name: "王五", age: 19}, school: "heh"}
	fmt.Println(s2)

	var s3 Student1
	s3.Person1.name = "王五"
	s3.Person1.age = 19
	fmt.Println(s3)

	//提升字段，可直接访问Person1的age，就像访问父类的对象一样
	s3.age = 30
	fmt.Println(s3)
	fmt.Println(s2.age)
	/*
		s3.Person1.age -> s3.age
		Student1结构体将Person结构体作为一个匿名字段了
		那么Person1中的字段，对于Student1来讲，就是提升字段
		Student1对象直接访问Person1中的字段
	*/

	// 模拟方法继承
	p1.eat()
	s1.eat() // 子类对象，访问父类方法

	// 方法的扩展
	s1.book() // 子类对象，访问自己新增的方法
	p1.eat()

	// 如果存在方法的重写，子类对象访问重写的方法
	s1.eat()
}
