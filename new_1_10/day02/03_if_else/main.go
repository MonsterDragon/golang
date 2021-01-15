package main

import "fmt"

func main(){
	age := 17
	if age > 18{
		fmt.Println("success")
	}else{
		fmt.Println("failed")
	}

	if age > 35{
		fmt.Println("middle age")
	}else if age > 18{
		fmt.Println("teenager")
	}else{
		fmt.Println("grow up")
	}

	if sit := 20; sit > 18{
		fmt.Println("success")
	}else{
		fmt.Println("failed")
	}
}
