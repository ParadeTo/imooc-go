package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("error")
	}
}

func apply(op func(int, int) int, a, b int) int {
	p := reflect.ValueOf(op).Pointer()
	opName := runtime.FuncForPC(p).Name()
	fmt.Println("calling function %s with args", opName, a, b)
	return op(a, b)
}

func div(a, b int) (q, r int) {
	return a / b, a % b
}

func main() {
	fmt.Println(apply(func(a int, b int) int {
		return a + b
	}, 3, 4))
}
