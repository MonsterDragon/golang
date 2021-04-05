package main

import "fmt"

type Address struct {
	Province string
	City string
}

type User struct {
	Name string
	Gender string
	Address Address
}

func main() {
	user1 := User{"prof", "female", Address{"heilongjiang", "harbin"}}
	fmt.Printf("user1=%#v\n", user1)
}