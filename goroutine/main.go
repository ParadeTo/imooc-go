package main

import (
	"time"
	"fmt"
)

func main() {
	//var a [10]int
	//for i := 0; i < 1000; i++ {
	//	go func(i int) {
	//		for {
	//			fmt.Println("a")
	//			//a[i]++
	//			//runtime.Gosched()
	//		}
	//	}(i)
	//}
	//time.Sleep(time.Minute * 5)
	//fmt.Println(a)

	c := make(chan bool)
	closeChan := make(chan bool)

	go func() {
		for {
			select {
			case <- c:
			case a := <- closeChan:
				fmt.Print("close chan")
				fmt.Print(a)
			}
		}
	}()

	close(closeChan)
	time.Sleep(time.Second)
}
