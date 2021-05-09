package main

import "fmt"

func main() {
	s := []string{"red", "", "black", "", "", "pink", "blue"}
	s1 := []string{}
	for _, i := range s {
		if i != "" {
			s1 = append(s1, i)
		}
	}
	fmt.Printf("%q\n", s1)
}
