package main

import "fmt"

func main() {

	//var w int
	//fmt.Scanf("%d", &w)
	//switch w {
	//case 1:
	//	fmt.Println("Monday")
	//case 2:
	//	fmt.Println("Tuseday")
	//case 3:
	//	fmt.Println("Wednesday")
	//case 4:
	//	fmt.Println("Thursday")
	//case 5:
	//	fmt.Println("Friday")
	//case 6:
	//	fmt.Println("Saturday")
	//case 7:
	//	fmt.Println("Sunday")
	//default:
	//	fmt.Println("error input")
	//}

	var score int
	fmt.Scanf("%d", &score)
	switch score/10 {
	case 10:
		//fmt.Println("a")
		fallthrough
	case 9:
		fmt.Println("a")
	case 8:
		fmt.Println("b")
	default:
		fmt.Println("wrong score")
	}
}