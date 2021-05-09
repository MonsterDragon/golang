package main

import "fmt"

func main() {
	// 删除map的kv，使用delete()函数
	var  m1 map[int]string = map[int]string{1:"Luffy", 2:"Nami", 3:"zoro"}

	delete(m1, 1)
	fmt.Println(m1)
}
