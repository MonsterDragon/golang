package main

import (
	"fmt"
)

func test(arr [10]int){
	for i:=0;i<len(arr);i++{
		fmt.Println(arr[i])
	}
}

func BubbleSort(arr [10]int) [10]int{
	for i:=0; i<len(arr) - 1; i++{
		for j:=0; j<len(arr)-i-1; j++{
			if (arr[j] > arr[j+1]){
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
	return arr
}

func main() {
	var arr[10] int = [10]int{1, 2, 3, 4, 5, 7:10}
	fmt.Println(arr)
	test(arr)

	var arr_2[10] int = [10]int{9, 1, 5, 6, 7, 3, 10, 2, 4, 8}
	BubbleSort(arr_2)
}
