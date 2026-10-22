package main

import (
	"fmt"
	"os"
)

func main() {
	// 如果参数很多的情况下，使用 Join 可以更加简单以及高效
	//fmt.Println(strings.Join(os.Args[1:], " "))

	// 如果我们不关心格式，只想查看值，也许是为了调试，我们可以让 Println 为我们格式化结果：
	// 这条语句的输出类似于我们从 strings.Join 中得到的结果，但会带有外层括号。 任何切片都可以用这种方式打印。
	fmt.Println(os.Args[1:], " ")
}
