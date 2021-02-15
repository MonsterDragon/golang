package main

import "fmt"

func main() {
	a := 10
	b := 3.14
	//c := float64(a) * b
	//fmt.Println(c)
	//fmt.Printf("%T\n", a)
	c := a * int(b)
	fmt.Println(c)

	var d int32 = 10
	var e int64 = 20

	f := int64(d) + e
	fmt.Println(f)

	var time int = 107653
	fmt.Printf("day: %d, hour: %d, min: %d, sec: %d\n", time/60/60/24%365, time/60/60%24, time/60%60, time%60)
}
