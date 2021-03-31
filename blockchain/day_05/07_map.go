package main

import "fmt"

func main()  {
	// map[KeyType]valueType
	m := make(map[int]string, 1)
	m[100] = "aa"
	fmt.Println(m)
	fmt.Println(m[100])
	st, ok := m[100]
	if ok{
		fmt.Println(st)
	} else {
		fmt.Println("null")
	}

	scoreMap := make(map[string]int)
	scoreMap["zhangsan"] = 90
	scoreMap["xiaoming"] = 100
	scoreMap["wangwu"] = 60

	for k, v := range scoreMap{
		fmt.Println(k, v)
	}

	delete(scoreMap, "zhangsan")
	fmt.Println(scoreMap)
}
