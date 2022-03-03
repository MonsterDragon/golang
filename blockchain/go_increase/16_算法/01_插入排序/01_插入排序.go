package main

import "fmt"

func InsertSort(list []int) {
	var i, j, tmp int
	n := len(list)

	for i = 1; i < n; i++ {
		tmp = list[i]
		if list[i] < list[i-1] {
			for j = i - 1; j >= 0 && tmp < list[j]; j-- {
				list[j+1] = list[j]
			}
			list[j+1] = tmp
		}
	}
}

func change(a *int) {
	*a += 1
}

func main() {
	num := []int{6, 2, 7, 3, 8, 5}
	InsertSort(num)

	a := 1
	change(&a)

	fmt.Printf("%#v\n", num)
	fmt.Printf("%d\n", a)
}
