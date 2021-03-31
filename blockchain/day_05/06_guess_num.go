package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(900) + 100
	var n int

	for {
		fmt.Scanf("%d", &n)
		fmt.Println(n)
		if (n > 100 && n < 999){
			if n > random{
				fmt.Println("bigger")
			} else if n < random {
				fmt.Println("smaller")
			} else if n == random{
				fmt.Println("equal")
				break
			}
		}
		fmt.Println("error num")
	}
}
