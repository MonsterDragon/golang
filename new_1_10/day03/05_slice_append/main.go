package main

import "fmt"

func main() {
	s1 := []string{"hh", "aa", "bb"}
	s2 := [...]string{"ccc", "ddd", "eee"}
	// var s4 [6]string
	fmt.Println(s1, s2)
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1) // 0xc000118330
	// s1[3] = "ff"

	s1 = append(s1, "rr") // 调用append函数可以用未初始化的变量通过“:=”来进行赋值
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1) // 0xc000070060
	s1 = append(s1, "bb")
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))
	fmt.Printf("%p\n", s1) // 0xc000070060
	s3 := append(s1, "kk")
	fmt.Println(s3, s1) // s1不变 s3: [hh aa bb rr bb kk] s1: [hh aa bb rr bb]
	fmt.Printf("len(s3)=%d, cap(s3)=%d\n", len(s3), cap(s3))

	//s4 := append(s3, "qq") // append不能给初始化了的数组进行赋值
	//fmt.Println(s4)

	// 首先判断，如果申请容量（cap）大于2倍的就容量（old.cap），最终容量（newcap）就是新申请的容量（cap）
	// 否则判断，如果旧切片的长度小于1024，则最终容量（newcap）就是旧容量（old.cap）的两倍，即（newcap = doublecap）
	// 否则判断，如果旧切片长度大于等于1024，则最终容量（old.cap）开始循环增加原来的1/4，即（newcap=old.cap, for{newcap += newcap/4}）直到最终容量（newcap）大于等于新申请的容量（cap）即（newcap >= cap）
	// 如果最终容量(cap)计算溢出，则最终容量（cap）就是新申请容量（cap）

	s1 = append(s1, "ability", "source")
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))
	s1 = append(s1, s3...) // 将s3拆开
	fmt.Println(s1)
	fmt.Printf("len(s1)=%d, cap(s1)=%d\n", len(s1), cap(s1))
}
