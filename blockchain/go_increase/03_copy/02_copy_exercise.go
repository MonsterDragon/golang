package main

import "fmt"

func dele(a []int, b int) []int {
	s1 := a[b:]
	s2 := a[b+1:]
	copy(s1, s2)
	return a
}

func main() {
	data := []int{5, 6, 7, 8, 9}
	adata := dele(data, 2)
	fmt.Println(adata[:len(adata)-1])
}
