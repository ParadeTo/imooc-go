package main

import (
	"time"
	"math/rand"
	"fmt"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker (id int, c chan int) {
	for n := range c {
		time.Sleep(700 * time.Millisecond)
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	n := 0
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: // nil chan will not be selected
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Timeout")
		case <-tick:
			fmt.Println("Queue:")
			fmt.Println(values)
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}
}
