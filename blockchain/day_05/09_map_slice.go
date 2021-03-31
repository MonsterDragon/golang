package main

import "fmt"

func main() {
	var mapslice = make([]map[string]string, 3) //[{}, {}, {}]
	for index, value := range mapslice {
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	fmt.Println("after init")
	mapslice[0] = make(map[string]string, 10)
	mapslice[0]["name"] = "laowang"
	mapslice[0]["password"] = "123456"
	mapslice[0]["address"] = "hongqidagai"
	for index, value := range  mapslice{
		fmt.Printf("index:%d value:%v\n", index, value)
	}
	//index:0 value:map[]
	//index:1 value:map[]
	//index:2 value:map[]
	//after init
	//index:0 value:map[address:hongqidagai name:laowang password:123456]
	//index:1 value:map[]
	//index:2 value:map[]

	}
