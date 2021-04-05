package main

import "fmt"

func swap(p *[]int) {
	(*p)[0] = 10
	for index, value :=range *p {
		fmt.Printf("index=%d, value=%d\n", index, value)
	}
}

func main() {
	a := []int{1, 2, 4, 3}
	fmt.Println(a)
	swap(&a)
	fmt.Println(a)

	for index, value := range a{
		fmt.Printf("index=%d value=%d\n", index, value)
	}
}
