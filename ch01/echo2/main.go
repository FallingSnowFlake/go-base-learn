package main

import (
	"fmt"
	"os"
)

func main() {

	// 声明变量的方式
	//s := ""
	//var s string
	//var s = ""
	//var s string = ""
	// 在实践中，通常应该使用前两种形式之一：用显式初始化表明初始值是重要的，用隐式初始化表明初始值无关紧要。

	s, sep := "", ""

	// 在循环的每次迭代中，range 都会生成一对值：索引和该索引处元素的值
	// 下划线 _（空白标识符）。凡是语法需要变量名但程序逻辑不需要的地方，都可以使用空白标识符，例如当我们只需要元素值时，丢弃不需要的循环索引。
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
