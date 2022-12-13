package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	//errorTest()
	//errorTest1()

	// 自定义错误信息
	textError()
}

/*
error: 内置的数据类型，内置的接口。 一种类型，和int,float一样
使用go语言提供好的包：
errors包下的函数：New(),创建一个error对象
	err1 := errors.New("创建错误对象")

fmt包下的Errorf()函数：

*/

func errorTest() {
	f, err := os.Open("text.txt")
	if err != nil {
		fmt.Println(err)
		// error的表示
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println(ins.Op)
			fmt.Println(ins.Path)
			fmt.Println(ins.Err)
		}
	}
	fmt.Println(f)
}

func errorTest1() {
	err1 := errors.New("创建错误对象")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1)

	err2 := fmt.Errorf("错误信息")
	fmt.Println(err2)
	fmt.Printf("%T", err2)
}

/*
自定义错误类型
*/

type areaError struct {
	msg   string
	radis float64
}

// 实现error接口，就是实现Error方法
func (a *areaError) Error() string {
	return fmt.Sprintf("error: 错误信息%s, %.2f", a.msg, a.radis)
}

// 自定义错误方法，用来打印错误的半径
func (a *areaError) ErrorType() string {
	return fmt.Sprintf("error: 错误信息%.2f", a.radis)
}

func circleArea(radis float64) (float64, error) {
	if radis < 0 {
		return 0, &areaError{"半径非法", radis}
	}
	return math.Pi * radis * radis, nil
}

func textError() {
	radis := -3.0
	area, err := circleArea(radis)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError); ok {
			fmt.Println(err.msg, err.radis)
			fmt.Println(err.ErrorType())
		}
		return
	}
	fmt.Println(area)
}

/*
go语言使用panic(), recover(), 实现程序中的极特殊的异常处理
panic():让当前的程序进入恐慌，终端程序的执行
recover(): 让程序恢复，必须在defer函数中执行
*/
