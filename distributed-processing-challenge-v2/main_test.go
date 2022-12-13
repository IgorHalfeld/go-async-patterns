package main

import (
	"log"
	"os"
	"testing"
)

func TestHelloWorld(t *testing.T) {
	t.Run("should sum all numbers on numbers.txt", func(t *testing.T) {
		f, err := os.ReadFile("./numbers.txt")
		if err != nil {
			log.Fatalln(err.Error())
		}

		result := sumAllNumbersFromFile(f)
		expected := 1024126520

		if result != expected {
			t.Errorf("Result: +%v, Expected: %+v", result, expected)
		}
	})
}
