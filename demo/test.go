package main

import "fmt"

func appendPrimeFilterC(inputC <-chan int, prime int) chan int {
	nextC := make(chan int, 10)
	go func() {
		for e := range inputC {
			if e%prime == 0 && e != prime {
				continue
			}
			nextC <- e
		}
		close(nextC)
	}()
	return nextC
}

func seed(minNum, maxNum int) <-chan int {
	numC := make(chan int, 10)
	go func() {
		for i := minNum; i < maxNum; i++ {
			numC <- i
		}
		close(numC)
	}()
	return numC
}

func showPrime(begin, end int) {
	source := seed(begin, end)
	primC := appendPrimeFilterC(source, 2)
	for {
		num, notEmpty := <-primC
		if !notEmpty {
			break
		}
		fmt.Println(num)
		primC = appendPrimeFilterC(primC, num)
	}
}

func main() {
	showPrime(2, 100)
}
