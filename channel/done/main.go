package main

import (
	"fmt"
	"sync"
)

func doWork (id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{in: make(chan int), done: func() {
		wg.Done()
	}}
	go doWork(id, w)
	return w
}

func chanDemo () {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	wg.Wait()
}

func main () {
	fmt.Println("Channel as first-class citizen.")
	chanDemo()
}
