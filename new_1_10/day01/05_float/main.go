package main

import (
	"fmt"
)

func main()  {
	// math.MaxFloat32
	f1 := 1.2345
	fmt.Printf("%T\n", f1) // Go中默认小数都是float64类型
	f2 := float32(1.2345)
	fmt.Printf("%T\n", f2)
}
