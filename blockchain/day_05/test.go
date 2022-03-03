package main

import "fmt"

func main() {
	fmt.Print("a: ", 'a')
	fmt.Print("\nA: ", 'A')
	fmt.Print("\n1: ", '1')
	fmt.Print("\n0: ", '0')
	fmt.Print("\n9: ", '9')

	w := "93python22"

	for _, i := range w{
		if i >= '0' && i <= '9' {
			fmt.Printf("%c\n", i)
		}else {
			continue
		}
	}
}
