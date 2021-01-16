package main

import "fmt"

func switchDemo(){
	finger := 3
	switch finger{
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	default:
		fmt.Println("输入错误")
	}
}

func goto_demo(){
	for i:=0;i<10;i++{
		for j:=0;j<10;j++{
			if j == 2{
				goto breakTag
			}
			fmt.Println(i, j)
		}
	}

breakTag:
	fmt.Println("结束for循环")
}

var flag bool
func double_break(){
	flag = false
	for i:=0; i<10; i++{
		for j:=0; j<10; j++{
			if j == 2{
				flag = true
				break
			}
			//fmt.Println()
			fmt.Println(i, j)
		}
		if flag{
			break
		}
	}
}

func  main()  {
	//for i:=0; i<10; i++{
	//	if i == 5{
	//		break // 跳出for循环
	//	}
	//	fmt.Println(i)
	//}
	//fmt.Println("over")

	for i:=0; i<10; i++{
		if i == 5{
			continue
		}else {
			fmt.Println(i)
		}
	}
	switchDemo()
	goto_demo()
	double_break()
}
