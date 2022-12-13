package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func sumAllNumbersFromFile(reader *bufio.Reader) int {
	total := 0

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err != io.EOF {
				log.Println("End!: ", err.Error())
			}

			break
		}

		for _, numberStr := range line {
			number, err := strconv.Atoi(string(numberStr))
			if err != nil {
				continue
			}

			total += number
			fmt.Println("Total now:", total)
		}
	}

	return total
}

func main() {
	f, err := os.Open("./numbers.txt")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()

	reader := bufio.NewReader(f)

	total := sumAllNumbersFromFile(reader)

	fmt.Println("Total:", total)
}
