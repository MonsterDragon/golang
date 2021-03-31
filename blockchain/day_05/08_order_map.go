package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	rand.Seed(time.Now().UnixNano())

	var scoreMap = make(map[string]int, 200)

	for i := 0; i < 100; i++{
		key := fmt.Sprintf("std%02d", i)
		value := rand.Intn(100)
		scoreMap[key] = value
	}

	//var keys = make([]string, 0, 200)
	//for key := range scoreMap{
	//	keys = append(keys, key)
	//}
	//fmt.Println(keys)
	//
	//sort.Strings(keys)
	//for _, key := range keys{
	//	fmt.Println(key, scoreMap[key])
	//}
	for key := range scoreMap{
		fmt.Println(key)
	}
}
