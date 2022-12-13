package main

import (
	"bufio"
	"log"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("should sum all numbers on numbers.txt", func(t *testing.T) {
		f, err := os.Open("./numbers.txt")
		if err != nil {
			log.Fatalln(err.Error())
		}
		defer f.Close()

		reader := bufio.NewReader(f)
		result := sumAllNumbersFromFile(reader)
		expected := 43907621

		if result != expected {
			t.Errorf("Result: +%v, Expected: %+v", result, expected)
		}
	})
}
