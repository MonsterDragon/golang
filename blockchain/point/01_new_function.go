package main

import "fmt"

func main() {
	var p *int
	p = new(int)
	*p = 57
	fmt.Println(*p) // new( )函数的作用就是C语言中的动态分配空间。但是在这里与C语言不同的地方，就是最后不需要关系该空间的释放。GO语言会自动释放。这也是比C语言使用方便的地方
}
