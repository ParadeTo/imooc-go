package main

import (
	"time"
	"fmt"
)

func main() {
	var a [10]int
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				fmt.Println("a")
				//a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Minute * 5)
	fmt.Println(a)
}
