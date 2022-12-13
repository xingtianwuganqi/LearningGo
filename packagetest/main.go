package main

import (
	"packagetest/pk1"
)

func main() {
	//utils.Count()
	//utils.MyTest()
	//pk1.MyTest()
	pk1.MyTest()
}

/*
main函数、init函数
go语言的保留函数，init() 用户初始化信息，mian()用于作为程序的入口
相同点：
两个函数在定义时不能有任何的参数和返回值
该函数只能由go程序自动调用，不可以被引用
不同点：
init可以应用于任意包中，且可以重复定义多个
main函数只能用于main包中，且只能定义一个。

两个函数的执行顺序：
在main包中go文件默认总是会被执行。
对同一个go文件中的init函数，调用顺序从上到下。
对同一个package中的不同文件，将文件名按字符串进行"从大到小"排序
之后顺序调用各文件中的init函数

对于不同package中的不同文件，如果不相互依赖的话，按照main包中import的顺序调用其包中的init函数

如果package存在依赖，调用顺序为最后被依赖的最先被初始化，例如：导入顺序main -> A -> B -> C
则初始化顺序为 C->B->A->main，一次执行对应的init方法，main包总是被最后一个初始化，因为它总是依赖别的包

避免出现循环 import，例如：A->B->C->A;
一个包被其他多个包import，但只能被初始化一次
*/
/*
关于包的使用：
1.一个目录下的统计文件归属一个包，package的声明要一致
2.package声明的包和对应的目录名可以不一致，但习惯上还是写成一致的
3.包可以嵌套
4.同包下的函数不需要导入包，其他的包不能使用
5.main包，main()函数所在的包，其他包不能使用
6.导入包的时候，路径要从src下开始写
*/

/*
包（packge）是多个 Go 源码的集合，是一种高级的代码复用方案，Go 语言为我们提供了 很多内置包，如:fmt、strconv、strings、sort、eros、time、encoding/json、os、io 等
Golang 中的包可以分为三种：1、系统内置包 2、自定义包 3、第三方包 系统内置包: Golang
语言给我们提供的内置包，引入后可以直接使用，如 fmt、srconv、strings、sort、eros、time、encoding/json、os、io 等
自定义包：开发者自己写的包

包管理工具 go mod
在 Golang1.11版本之前如果我们要自定义包的话必须把项目放在 GOPATH 目录。Go1.11版本之后无需手动配置环境变量，使用 go mod 管理项目，也不需要非得把项目放到 GOPATH
指定目录下，你可以在你磁盘的任何位置新建一个项目 , Go1.13以后可以彻底不要 GOPATH了。


在项目下 go mod init packagetest
go: modules disabled by GO111MODULE=off; see 'go help modules'

# 5. 创建模块
go mod init "github.com/qq1060656096/hellomod"
# 创建模块失败会提示: "go: modules disabled inside GOPATH/src by GO111MODULE=auto; see 'go help modules'"

# 为什么创建模块失败
# 因为GO111MODULE默认值是auto, 在 GOPATH/src 之外的目录才开启模块支持
# 我们有2中方式解决以上问题
#   第1种: 在 GOPATH/src 之外的目录创建模块
#   第2种: 直接设置GO111MODULE=on 模块支持

# 这里我们直接使用第2种
# 使用命令，将GO111MODULE改为on
export GO111MODULE=on
go mod init "github.com/qq1060656096/hellomod"
# 创建模块成功会提示"go: creating new go.mod: module github.com/qq1060656096/hellomod"
# 模块创建后里面会有一个go.mod文件

# 6. 查看go.mod文件的内容
$ cat go.mod
module github.com/qq1060656096/hellomod
# 里面只有一行, 就定义的模块名字
*/
