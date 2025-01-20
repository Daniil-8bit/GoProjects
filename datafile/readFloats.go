package datafile

import (
	"bufio"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {

	//var numbers [3]float64
	var numbers []float64

	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	//i := 0

	/*for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return numbers, err
		}
		i++
	}*/

	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nil, err
		}

		numbers = append(numbers, number)
	}

	err = file.Close()

	if err != nil {
		return nil, err
	}

	if scanner.Err() != nil {
		return nil, err
	}

	return numbers, nil
}
