package main

import "fmt"

type Operation_1 struct {
	numA float64
	numB float64
}

type  Getresulter interface {
	GetResult() float64
}

type OperationSum struct {
	Operation_1
}

func (o *OperationSum) GetResult() float64{
	result := o.numA + o.numB
	return result
}

type OperationSub struct {
	Operation_1
}

func (o *OperationSub) GetResult() float64 {
	result := o.numA - o.numB
	return result
}

type OperationFactory struct {

}

func (o *OperationFactory) CreateOption(option string, numA float64, numB float64) float64 {
	var result float64
	switch option {
	case "+":
		add := &OperationSum{Operation_1{numA, numB}}
		result = OperationWho(add)
	case "-":
		sub := &OperationSub{Operation_1{numA, numB}}
		result = OperationWho(sub)
	}
	return result
}

func OperationWho(i Getresulter) float64 {
	return i.GetResult()
}

func main()  {
	var op OperationFactory
	s := op.CreateOption("+", 10, 20)
	fmt.Println(s)
}
