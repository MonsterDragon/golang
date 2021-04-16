package main

import "fmt"

type CalcSuper interface {
	SetData(data ...interface{})
	CalcOperate() float64
}

type Operation_2 struct {
	numA float64
	numB float64
}

type Add struct {
	Operation_2
}

// 函数，创建加法的实例
func NewAdd() *Add {
	instance := new(Add)
	return instance
}

func (a *Add) SetData(data ...interface{}) {
	if len(data) != 2 {
		fmt.Println("error, Need two parameters")
		return
	}

	if _, ok := data[0].(float64); !ok {
		fmt.Println("error, Need float64 parameters")
		return
	}

	if _, ok := data[1].(float64); !ok {
		fmt.Println("error, Need float64 parameters")
		return
	}

	a.numA, _ = data[0].(float64)
	a.numB, _ = data[1].(float64)
}

func (a *Add) CalcOperate() float64 {
	return a.numA + a.numB
}

type Subtraction struct {
	Operation_2
}

func NewSubtraction() *Subtraction {
	instance := new(Subtraction)
	return instance
}

func (a *Subtraction) SetData(data ...interface{}) {
	if len(data) != 2 {
		fmt.Println("error, Need two parameters")
		return
	}

	if _, ok := data[0].(float64); !ok {
		fmt.Println("error, Need float64 parameters")
		return
	}

	if _, ok := data[1].(float64); !ok {
		fmt.Println("error, Need float64 parameters")
		return
	}

	a.numA, _ = data[0].(float64)
	a.numB, _ = data[1].(float64)
}

func (a *Subtraction) CalcOperate() float64 {
	return a.numA - a.numB
}

type CalcFactory struct {
	
}

func NewCalcFactory() *CalcFactory {
	instance := new(CalcFactory)
	return instance
}

func (f *CalcFactory) CreateOperate(opType string) CalcSuper {
	var op CalcSuper
	switch opType {
	case "+":
		op = NewAdd()
	case "-":
		op = NewSubtraction()
	default:
		panic("error do not has this operation") // 能够改变程序的控制流，调用panic后会立刻停止执行当前函数的剩余程序，并在当前GOroutine中递归执行调用方的defer
	}
	return op
}

func main() {
	factory := NewCalcFactory()
	op := factory.CreateOperate("+")
	op.SetData(1.5, 20.0)
	fmt.Println(op.CalcOperate())
	op = factory.CreateOperate("-")
	op.SetData(2.3, 3.4)
	fmt.Println(op.CalcOperate())
}
