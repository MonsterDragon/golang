package main

import "fmt"

func main() {
	var arr[3][4] int = [3][4]int{{1, 2, 3, 4}, {2, 3, 4, 5}, {3, 4, 5, 6}}

	for i:=0; i<3; i++ {
		for j:=0; j<4;j++ {
			fmt.Print(arr[i][j])
		}
		fmt.Println()
	}
	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))
}
