package main

import "fmt"

func f1() (r int) {
	defer func() {
		r++
	}()
	r = 0
	return
}

func main() {
	i := f1()
	fmt.Println(i)
}
