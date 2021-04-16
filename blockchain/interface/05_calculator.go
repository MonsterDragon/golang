package main

import "fmt"

type Operation struct {

}

func (o *Operation) GetResult(numA float64, numb float64, operate string) float64 {
	var result float64
	switch operate {
	case "+":
		result = numA + numb
	case "-":
		result = numA - numb
	}
	return result
}

func main() {
	var operator Operation
	result := operator.GetResult(10, 13, "+")
	fmt.Println(result)
}
