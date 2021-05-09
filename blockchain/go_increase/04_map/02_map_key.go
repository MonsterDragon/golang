package main

import "fmt"

func main() {
	// 判断map中key是否存在
	var  m1 map[int]string = map[int]string{1:"Luffy", 2:"Nami", 3:"zoro"}
	for k, v := range m1 {
		fmt.Println(k, v)
	}

	// range返回的key/value 可以省略value
	for K := range m1{
		fmt.Println(K)
	}

	// 判断key是否存在
	if v, ok := m1[1]; ok == true {
		fmt.Println(v)
	}
}
