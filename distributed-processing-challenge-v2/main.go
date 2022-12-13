package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Dispatch struct {
	Line         string
	TotalOfLines int
}

type Worker struct {
	Sum          int
	TotalOfLines int
}

type Collector struct {
	Sum            int
	LinesProcessed int
	TotalOfLines   int
}

func dispatch(file string) chan Dispatch {
	outCh := make(chan Dispatch)

	go func(f string) {
		defer close(outCh)

		lines := strings.Split(file, "\n")
		totalOfLines := len(lines)

		for _, line := range lines {
			outCh <- Dispatch{
				Line:         line,
				TotalOfLines: totalOfLines,
			}
		}
	}(file)

	return outCh
}

func worker(inCh chan Dispatch) chan Worker {
	outCh := make(chan Worker)

	go func() {
		defer close(outCh)

		for in := range inCh {
			numbers := strings.Split(in.Line, " ")

			lineTotalSum := 0
			for _, numberStr := range numbers {
				number, err := strconv.Atoi(numberStr)
				if err != nil {
					continue
				}

				lineTotalSum += number
			}

			outCh <- Worker{
				Sum:          lineTotalSum,
				TotalOfLines: in.TotalOfLines,
			}
		}
	}()

	return outCh
}

func collector(inCh chan Worker) chan Collector {
	outCh := make(chan Collector)

	go func() {
		defer close(outCh)

		result := Collector{Sum: 0}
		for processed := range inCh {
			result.Sum += processed.Sum
			result.LinesProcessed += 1
			result.TotalOfLines = processed.TotalOfLines
		}

		outCh <- result
	}()

	return outCh
}

func sumAllNumbersFromFile(bytes []byte) int {
	total := 0
	file := string(bytes)

	dispatchCh := dispatch(file)
	workerCh := worker(dispatchCh)
	collectorCh := collector(workerCh)

	for value := range collectorCh {
		fmt.Printf("Collector result: %+v\n", value)
		total = value.Sum
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
