package main

import (
	"encoding/json"
	"fmt"
)

type Stud struct {
	ID int
	Gender string
	Name string
}

type Class struct {
	Title string
	Students []*Stud
}

func main() {
	c := &Class{Title: "305", Students: make([]*Stud, 0, 200)} // type cap len
	for i:=0; i<10; i++ {
		stu := &Stud{
			Name: fmt.Sprintf("std%02d", i),
			Gender: "male",
			ID: i,
		}
		c.Students = append(c.Students, stu)
	}

	// json 序列化：结构体 --》 json格式的字符串
	data, err := json.Marshal(c)
	if err != nil {
		fmt.Println("json marshal failed!")
		return
	}
	fmt.Printf("json:%s\n", data)
	// json反序列化
	// go中的反引号类似于c语言的'''
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"stu00"},{"ID":1,"Gender":"男","Name":"stu01"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c1 := &Class{}
	err = json.Unmarshal([]byte(str), c1)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c1)
}
