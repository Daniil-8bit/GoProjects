// Package keyboard helps to get float value from user
// Use it to ask user write float value and use it in your program
package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// GetFloatValue() can be used for getting user float value
// Use it your code to save user value and use it in program
// Have a good day!
// You are ready to GO!
func GetFloatValue() (float64, error) {

	fmt.Println("Write your value:")

	reader := bufio.NewReader(os.Stdin)
	value, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	value = strings.TrimSpace(value)

	floatNum, err := strconv.ParseFloat(value, 64)

	if err != nil {
		return 0, err
	}

	return floatNum, nil
}
