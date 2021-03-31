package main

import "fmt"

func main() {
	s := []int{1, 2, 3 ,4 , 5}
	slice := s[2:5]
	fmt.Println(slice)

	slice[0] = 23
	fmt.Println(slice)
	fmt.Println(s)
	fmt.Printf("%p, %p, %p, %p, %p, %p\n", &s, &s[0], &s[1], &s[2], &s[3], &s[4])
	fmt.Printf("%p\n", &slice)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("%p, %p, %p, %p, %p, %p\n", &b, &b[0], &b[1], &b[2], &b[3], &b[4])
}