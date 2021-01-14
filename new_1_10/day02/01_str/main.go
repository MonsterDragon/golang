package main

import (
	"fmt"
	"strings"
)

// Go语言中字符串使用双引号包裹的
// Go语言中单引号包裹的是字符

func main(){
	path := "C:\\Users\\admin\\.cargo\\bin"
	fmt.Println(path)

	s1 := `
		东临碣石
		以观沧海
	`
	fmt.Println(s1)
	// ``原样输出
	// 字符串相关操作

	name := "理想"
	world := "dss"
	fmt.Println(name + world)
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)

	// 分割
	ret := strings.Split(path, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(s1,"东"))

	// 前缀
	fmt.Println(strings.HasPrefix(name, "理"))
	fmt.Println(strings.HasSuffix(name, "想"))

	// 索引
	s3 := "aaabbbcccddd"
	fmt.Println(strings.Index(s3,"a"))
	fmt.Println(strings.LastIndex(s3, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))
}
