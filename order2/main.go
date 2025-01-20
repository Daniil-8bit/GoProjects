package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func averageNum(numbers ...float64) float64 {
	var sum float64 = 0

	for _, value := range numbers {
		sum += value
	}

	return sum / float64(len(numbers))
}

func main() {

	arguments := os.Args[1:]

	var sum float64

	var numbers []float64

	for _, value := range arguments {
		num, err := strconv.ParseFloat(value, 64)
		if err != nil {
			log.Fatal(err)
		}

		numbers = append(numbers, num)
		sum += num
	}

	fmt.Printf("Sum: %.2f\n", sum)
	fmt.Printf("You need to buy %.2f of meat\n", averageNum(numbers...))

	//fmt.Println(os.Args[1:])
}
