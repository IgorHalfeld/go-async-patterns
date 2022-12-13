package main

import "fmt"

func main() {
	ch := generator()

	for i := 0; i < 5; i++ {
		value := <-ch
		fmt.Println("Value:", value)
	}
}

func generator() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}
