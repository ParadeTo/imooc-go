package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

const (
	a = 3
	b = iota
	c
)

func variable() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func euler() {
	fmt.Println(cmplx.Pow(math.E, 1i*math.Pi) + 1)
}

// func main() {
// 	fmt.Println("hello world")
// 	// variableInit()
// 	fmt.Println(a, b, c)
// 	euler()
// }
