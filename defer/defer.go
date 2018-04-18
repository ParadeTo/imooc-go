package main

import "fmt"

func f () int {
	a := 1

	for i := 0; i<10; i++  {
		defer func() {
			fmt.Println("defer", i)
		}()
	}
	return a
}

func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

func main() {
	tryDefer()
	fmt.Println(f())
}
