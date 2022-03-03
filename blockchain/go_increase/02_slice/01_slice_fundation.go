// package main

// import "fmt"

// func main() {
// 	var a []int = []int {1, 2, 3, 4, 5}
// 	b := a[1:3:5]
// 	fmt.Println(b)

// 	var arr []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	s := arr[2:5]
// 	fmt.Println("s=", s, "len(s)=", len(s), "cap(s)=", cap(s))
// 	s_2 := s[2:7]
// 	fmt.Println("s=", s_2, "len(s)=", len(s_2), "cap(s)=", cap(s_2))
// 	// 切片名称[low:high:max]
// 	// low：起始下标位置
// 	// high：结束下标位置 len=high-low
// 	// cap = max - low
// 	// s[:high:max] 从0开始，到high结束
// 	// s[low:] 从0开始，到末尾
// }
