package main

import (
	"fmt"
	"log"

	"github.com/Daniil-8bit/GoProjects/datafile"
)

func main() {
	meatWeekWeight, err := datafile.GetFloats("C:\\Go\\read files\\meatWeights.txt")

	if err != nil {
		log.Fatal(err)
	}

	var sum float64

	fmt.Println("Write values of 3 last weeks:")

	/*for index, _ := range meatWeekWeight {

		var err error
		meatWeekWeight[index], err = keyboard.GetFloatValue()

		if err != nil {
			log.Fatal(err)
		}
	}*/

	for _, value := range meatWeekWeight {
		sum += value
	}

	fmt.Printf("sum: %.2f\n", sum)

	result := sum / float64(len(meatWeekWeight))

	fmt.Printf("You should buy %.2f killos of meat\n", result)
}
