package main

import "fmt"

func main() {
	var sliceMap = make(map[string][]string, 3) // {'a': []}
	fmt.Println(sliceMap)
	fmt.Println("after init")
	key := "china"
	value, ok := sliceMap[key]
	if !ok {
		value = make([]string, 0, 2)
	}
	value = append(value, "beijing", "shanghai")
	sliceMap[key] = value
	fmt.Println(sliceMap)
}
