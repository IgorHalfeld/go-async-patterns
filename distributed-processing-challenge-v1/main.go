package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sumAllNumbersFromFile(bytes []byte) int {
	total := 0
	file := string(bytes)

	lines := strings.Split(file, "\n")

	for _, line := range lines {
		numbers := strings.Split(line, " ")
		for _, numberStr := range numbers {
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				continue
			}

			total += number
		}
	}

	return total
}

func main() {
	f, err := os.ReadFile("./numbers.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}

	total := sumAllNumbersFromFile(f)

	fmt.Println("Total:", total)
}
