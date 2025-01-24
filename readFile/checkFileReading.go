package readFile

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func OpenFile(filePath string) (*os.File, error) {
	fmt.Println("Opening ", filePath)
	return os.Open(filePath)
}

func Closefile(file *os.File) {
	fmt.Println("Closing file...")
	file.Close()
}

func GetFloats(filePath string) ([]float64, error) {
	var numbers []float64
	file, err := OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	defer Closefile(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		num, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
		if err != nil {
			return nil, err
		}
		numbers = append(numbers, num)
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}
	return numbers, nil
}
