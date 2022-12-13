package main

import "fmt"

func main() {
	ch := generator()

	for value := range ch {
		fmt.Println("Value:", value)
	}
}

func generator() chan int {
	ch := make(chan int)

	go func() {
		defer close(ch)

		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	return ch
}
