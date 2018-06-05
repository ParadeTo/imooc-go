package main

import (
	"time"
	"fmt"
)

func worker (id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}
	// or
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo () {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
}

func bufferedChannel () {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}

// if closed, can still receive data, but the value is 0, can use it to
// decide if the channel is closed.
// if not closed, it will block
func channelClose () {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
}

func main () {
	fmt.Println("Channel as first-class citizen.")
	//chanDemo()
	fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
	time.Sleep(time.Second)
}
