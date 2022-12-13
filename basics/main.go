package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 2)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Iguin"
	}()

	value := <-ch
	fmt.Println("value", value)
}
