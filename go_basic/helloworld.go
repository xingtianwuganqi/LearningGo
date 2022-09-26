package main // 同文件下包名应该一致

import (
	"fmt"
	"math/rand"
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
	//constactValue()

	// 基础数据类型
	//dataType()

	// 条件语句
	//choseValue()

	// goto语句
	//gotoValue()

	// 数组
	//arrayValue()

	// map
	mapValue()
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

	//iota
	/*
		iota，特殊常量，可以认为是一个可以被编译器修改的常量
		每当定义一个const，iota的初始值为0
		每当定义一个常量，iota就会自动累加1
		直到下一个const出现，清零
	*/
	const (
		h = iota
		i = iota
		j = iota
	)

	const (
		k = iota
		l
	)

	const (
		m = iota // iota 0
		n = 1    // iota 1
		o = "q"  // iota 2
		p        //p和上一个常量的值一样， iota 自增到3
		q = iota //iota 自增到4
	)

	fmt.Println(h, i, j)
	fmt.Println(k, l)

	fmt.Println(m, n, o, p, q)
}

func dataType() {
	/*
		布尔类型
		数值类型
			整型：int int8 int16 ...
				uint8 uint16 ...
			浮点型：float32 float64
			complex64：32位实数和虚数 complex128：64位实数和虚数
			byte：类似uint8
			rune：类似int32

		字符串
		array
		slice
		map
		function
		pointer
		struct
		interface
		channel
	*/

	var b1 bool
	b1 = true
	var b2 bool = false
	fmt.Printf("%T,%t\n", b1, b1)
	fmt.Printf("%T,%t\n", b2, b2)

	var i1 int8 = -100
	var i2 int16 = 10000
	var i3 int32 = 199999999
	fmt.Println(i1, i2, i3)

	var i4 uint8 = 100
	fmt.Println(i4)

	// int和int64是不同类型
	var i5 int = 100
	fmt.Printf("%T,%d\n", i5, i5)

	/*
		字符串：
		1.概念：多个byte的集合，理解为一个字符序列
		2.语法：使用双引号
		也可使用 ``
		3.编码问题
		4.转义字符
		\n, \", \t
	*/

	var string1 string
	string1 = "wangergou"
	fmt.Printf("%T,%s\n", string1, string1)

	string2 := `hello world`
	fmt.Printf("%T,%s\n", string2, string2)

	string3 := `hel"lo wo"rld`
	fmt.Printf("%T,%s\n", string3, string3)

	// 基本数据类型转换
	/*
		数据类型转换：Type(Value)
		go 静态语言，定义、赋值、运算必须类型一致
		注意点：兼容类型可以转换
		常数：在有需要的时候自动转换
		变量：需要手动转换
	*/
	var value1 int8 = 10
	var value2 int16 = int16(value1)
	fmt.Println(value2)

	f1 := 3.14
	var f2 int = int(f1)
	fmt.Println(f1, f2)

	/*
		运算符
		1.算术运算符 + - * % ++ --
		2.关系运算符 > < >= <= == !=
		3.逻辑运算 && || !
		4.位运算符：
			将数值，转为二进制后，按位操作
			按位&：
				对应位的值如果都是1才为1，有一个为0就位0
			按位|：
				对应位的值如果都是0才为0，有一个为1就为1

			异或^:
				二元：a^b
					对应位的值不同为1，相同为0
				一元：^a
					按位取反
						1---->0, 0---->1
			位清空: &^
				对于 a &^ b
					对于b上的每个数值
					如果为0,则去a对应位上的数值
					如果为1，则结果为就取a
		5：赋值运算符：
			=、+=、-=、*=、/=、%=，
			<<= （左移位并赋值运算符）
			>>= （左移位并赋值运算符）
			&= 按位与赋值运算符
			|= 按位或赋值运算符
	*/

	/*
		输入和输出：
			fmt包：输入，输出

			输出：
				Print() // 打印
				Printf() // 格式化打印
				Println() // 换行打印

			格式化打印占位符:
			%v: 原样输出
			%T:打印类型
			%t:Bool类型
			%s:字符串
			%f：浮点
			%d：10进制的整数
			%b：二进制的整数
			%o：8进制
			%x，%X：16进制
				%x：0-9，a-f
				%X：0-9，A-F
			%c：打印字符
			%p：打印地址

		输入：
			Scanln()
	*/
	var scan1 int
	var scan2 float64
	fmt.Println("请输入一个整数，一个浮点类型")
	fmt.Scanln(&scan1, &scan2)
	fmt.Scanf("%d,%f", &scan1, &scan2)
	fmt.Printf("x:%d,y:%f", scan1, scan2)

}

func choseValue() {
	/*
			if语句的其他写法：
			if 初始化语句；条件 {

			}

			switch:
			1.注意点
			1.可以作用于其他类型上，case后的数值必须和switch作用的变量类型一致
			2.case 是无序的
			3.case后的数值是唯一的
			4.default语句是可选操作

			2.省略switch后的变量，相当于直接作用在true上
			switch {

			}

			3.case后可跟随多个数值
			case "A", "B"

			4.switch后可以多一条初始化语句

		5.switch 中的break和fallthrough
		break: 强制结束case，意味switch结束
		fallthrough：用于穿透switch，连接作用，匹配到case后，直接执行下一个case
		fallthrough：应该位于某个case的最后一行
	*/
	if num := 4; num > 0 { // num作用域只能在if中
		fmt.Printf("正数：%d\n", num)
	}

	num := 5
	switch num {

	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三")
	case 5:
		fmt.Println("第五")
		fallthrough
	default:
		fmt.Println("执行default")
	}

	switch language := "golang"; language {
	case "golang":
		fmt.Println(language)
	}

	/*
					1.for 循环
					for init; condition; post {

					}

					for 表达式1; 表达式2; 表达式3 {

					}
					表达式3在for循环后执行
				2.同时省略表达式1，和表达式3
				相当于while语句

				3.同时省略3个表达式：
				相当于while(true)
				注意点：当for循环中，省略了2，相当于直接作用于true上

			4.其他写法：for循环同时省略几个都可以

		break,continue;
		break:结束全部循环
		continue：结束当前循环
		注意点：多层循环嵌套，break，continue，默认结束里层循环
		如果想结束某层循环，可以给for循环起名
	*/

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d,hello world\n", i)
	}

	j := 1
	for j <= 5 {
		fmt.Println(j)
		j++
	}

	// 省略3个表达式，死循环
	for {
		j++
	}

	// for循环,
out:
	for h := 1; h <= 5; h++ {
		for k := 1; k <= 5; k++ {
			if k == 2 {
				// 结束外层循环
				break out
				//continue out
			}
		}
	}
}

func gotoValue() {
	/*
		goto语句：跳转到指定的行

	*/
	var a = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Printf("a的值为%d\n", a)
		a++
	}

	// 随机数
	num1 := rand.Int()
	fmt.Println(num1)
}

func arrayValue() {
	/*
			语法：var 数据名 [长度] 数据类型
			var 数据名 = [长度] 数据类型{元素1，元素2...}
			数据名 := [...]数据类型{元素...}

			长度和容量
			len()
			cap()

			基本数据类型也是值传递
			数组是值传递

			值类型：理解为存储的数值本身
		将数据传递给其他变量，传递的是数据的副本
		int,float,string,bool,array

		引用类型：理解为存储的数据的内存地址
		slice,map

	*/
	var array = [4]int{1, 2, 3, 4}
	fmt.Println(array)
	fmt.Println(len(array)) // 实际存储的数据量
	fmt.Println(cap(array)) // 能存储的最大数据量

	a := [...]int{1, 2, 3, 4}
	fmt.Println(a)

	// for循环遍历 for_range
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	for index, value := range array {
		fmt.Printf("%d，%d\n", index, value)
	}

	arr := [5]int{15, 23, 8, 10, 7}

	// 冒泡排序
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)

	// 二维数组, 3个一维数组，每个一维数组长度为4
	twoArr := [3][4]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{10, 11, 12, 13},
	}
	fmt.Println(twoArr)
	fmt.Println(twoArr[0][2])

	// 遍历
	for k := 0; k < len(twoArr); k++ {
		for l := 0; l < len(twoArr[k]); l++ {
			fmt.Println(twoArr[k][l])
		}
	}

	for _, arr := range twoArr {
		for _, val := range arr {
			fmt.Println(val)
		}
	}

	// 声明切片
	/*
		引用类型，指向一个底层数组
		make()
		func make(t Type, size ...IntegerType) Type
		第一个参数：类型
			slice,map,chan
		第二个参数：长度len
			实际存储元素的数量
		第三个参数：容量cap
		最多存储的元素数量


		切片Slice：
		1.每一个切片引用了一个底层数组，
		2.切片本身不存储任何数据，都是这个底层数组存储，所有修改切片也就是修改这个数组中的数据
		3.当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容（成倍增长）
		4.切片一旦扩容，就是重新指向一个新的底层数组
	*/
	var slice []int
	fmt.Println(slice)

	var slice2 = make([]int, 3, 10)
	fmt.Println(slice2)

	s4 := make([]int, 0, 5)
	s4 = append(s4, 1, 2)
	fmt.Println(s4)
	s4 = append(s4, 5, 7, 4)

	slice2[0] = 1
	s4 = append(s4, slice2...)
	s4 = append(s4, []int{15, 16}...)

	for _, value := range s4 {
		fmt.Printf("%d\n", value)
	}

	s5 := s4[:5]
	fmt.Println(s5)
	s6 := s4[1:6]
	fmt.Println("s6", s6)
	s7 := s4[1:]
	fmt.Println(s7)

	copy(s6, slice2) // 将s7的元素，拷贝到s6中
	fmt.Println("s6", s6)

}

func mapValue() {
	/*
		map：映射，是一种专门用于存储键值对的 集合，属于引用类型，
		存储特点：
		1.存储的是无序的键值对
		2.键不能重复，且和value值是一一对应的，map中的key不能重复，如果重复，那么新的value会覆盖原来的

		语法结构：
		1.创建map
		var map1 map[key类型]value类型,nil map,无法直接使用，只有声明，没有初始化
		var map2 = make(map[key类型]value类型)
		var map3 = map[key]value{key:value,key,value...}

		2.添加/修改
		map[key] = value
		key不存在，添加数据
		key存值，修改数据

		3.获取
		map[key] -->value
		value,ok := map[key]
		根据key获取对应的value
		如果key存在，value就是对应的值，ok为true
		如果key不存在，value是值类型的默认值，ok为false

		4.删除数据
		delete(map,key)
		如果key存在，直接删除
		如果key不存在，删除失败，对map没什么影响

		5.长度
		len(map)
	*/
	var map1 map[string]int // 只有声明，没有初始化
	fmt.Println(map1)
	//map1["hello"] = 1,报错

	var map2 = make(map[int]string)
	fmt.Println(map2)

	var map3 = map[string]int{"go": 1}
	fmt.Println(map3)

	if map1 == nil {
		map1 = make(map[string]int)
		fmt.Println(map1)
	}
	//存值
	map1["hello"] = 1
	fmt.Println(map1)
	fmt.Println(map1["hello"])

	// 判断key是否存在
	v1, ok := map1["hello"]
	if ok {
		fmt.Println(v1)
	}

	// 修改
	map1["hello"] = 3
	fmt.Println(map1)

	delete(map1, "hello")
	fmt.Println(map1)

	// 长度
	fmt.Println(len(map1))
}
