package main

import "fmt"

func main() {
	arr := [...]int{4, 2, 8, 0, 5, 7, 1, 3, 9}

	// 选择排序
	//for i:=0; i<len(arr)-1; i++{
	//	for j:=i+1; j<len(arr); j++{
	//		if arr[i] > arr[j]{
	//			arr[i], arr[j] = arr[j], arr[i]
	//		}
	//	}
	//}

	for i:=0; i<len(arr)-1; i++{
		for j:=0; j<len(arr)-i-1; j++{
			if arr[j] > arr[j+1]{
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}


	fmt.Println(arr)
}
