package main

import "fmt"

type Integer int

func (a Integer) Test(b Integer) Integer {
	return a + b //为整型绑定的函数，只不过在给整型绑定函数(方法)时，一定要通过type来指定一个别名，因为int类型是系统已经规定好了，无法直接绑定函数，所以只能通过别名的方式
}

func main() {
	var result Integer = 3
	r := result.Test(4)
	fmt.Println(r)
}
