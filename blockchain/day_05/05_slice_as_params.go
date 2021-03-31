package main

import "fmt"

func BubbleSort(s []int) {
	for i:=0; i<len(s)-1;i++{
		for j:=0; j<len(s)-i-1;j++{
			if s[j] > s[j+1]{
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	fmt.Println(s)
}

func main() {
	s := []int{1, 3 ,212, 43, 46}
	BubbleSort(s)
}