package main

/*
defer 传入的函数不是在退出代码块的作用域时执行的，它只会在当前函数和方法返回之前被调用。
*/

import (
	"fmt"
)

func main() {

	{	defer fmt.Println("aa")
		fmt.Println("bb")
	}

	fmt.Println("main ends")
	/*
	bb
	main ends
	aa
	*/
}
