package main

import "fmt"

type Student_3 struct {
	name string
	id int
}

func main() {
	i := make([]interface{}, 3)
	i[0] = 1
	i[1] = "hello go"
	i[2] = Student_3{"mike", 666}

	for index, data := range i {
		if value, ok := data.(int); ok == true {
			fmt.Printf("x[%d] type is int, value is %d\n", index, value)
		} else if value, ok := data.(string); ok == true {
			fmt.Printf("x[%d] type is string, value is %s\n", index, value)
		} else if value, ok := data.(Student_3); ok == true {
			fmt.Printf("x[%d] type is Student_3, value is %#v\n", index, value)
		}
	}
}