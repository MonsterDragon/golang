package main

import "fmt"

func main() {
	s := []string{"red", "black", "red", "pink", "blue", "pink", "blue"}
	new_str := []string{}
	for i:=0; i<len(s); i++{
		a := 0
		for j:=0; j<i; j++{
			if s[j] == s[i] {
				a = 1
				break
			}
		}
		if a == 0 {
			new_str = append(new_str, s[i])
		}
	}
	fmt.Println(new_str)
}
