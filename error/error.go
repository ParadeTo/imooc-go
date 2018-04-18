package main

import (
	"os"
	"bufio"
	"fmt"
	"errors"
)

func fibnaci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	f := fibnaci()
	for i := 0; i < 20; i++ {
		j := f()
		fmt.Println(j)
		fmt.Fprintln(writer, j)
	}
}

func main() {
	writeFile("./fib2.txt")
	// Create a file and use bufio.NewWriter.
	errors.New("this is a custom error")
}

