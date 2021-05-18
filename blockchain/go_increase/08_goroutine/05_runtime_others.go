package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOROOT()) // 返回goroot路径
}
