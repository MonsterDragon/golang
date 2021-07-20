package main

// func main() {
// 	// 匿名函数的优越性在于可以直接使用函数内的变量，不必申明
// 	getSqrt := func(a float64) float64 {
// 		return math.Sqrt(a)
// 	}
// 	fmt.Println(getSqrt(4))
// }

func main() {
	// Golang匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送
	fn := func() { println("hello world!") }
	// fmt.Println输出到标准输出，而println输出至标准错误。而且println在输入的打印参数和换行会追加空格
	// println用于程序启动和调试时使用
	fn()

	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	// 定义一个func(x int) int 类型的切片

	println(fns[0](100))
	println(fns[1](100))

	// 匿名函数作为结构体的结构字段
	d := struct {
		fn func() string // 匿名函数作为匿名结构体的字段
	}{
		fn: func() string { return "hello world!" }, // 对匿名结构体的匿名字段进行赋值
	}
	// 匿名结构体：没有名字的结构体，在创建匿名结构体时。同时创建对象
	// 变量名:=struct {
	// 	定义字段
	// }{
	// 	字段进行赋值
	// }

	println(d.fn())

	// 匿名函数在channel中进行传递
	fc := make(chan func() string, 2)
	fc <- func() string { return "hello world!" }

	println((<-fc)())
}
