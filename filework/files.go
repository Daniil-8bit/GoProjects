package filework

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func Work() {
	fmt.Printf("\n***Welcome to our working with files app!***\n\n")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()

		fmt.Println("Write number of action you need:")
		scanner.Scan()
		userNum, err := strconv.ParseInt(strings.TrimSpace(scanner.Text()), 10, 64)
		checkError(err)

		//fmt.Println("Your choise: ", userNum)

		switch userNum {
		case 1:
			fmt.Println("Write full path to exists file or new file:")
			scanner.Scan()
			filePath := scanner.Text()
			fmt.Println("Write text that you need to add:")
			scanner.Scan()
			userText := scanner.Text()
			addFileData(filePath, userText)
		case 2:
			fmt.Println("Write full path to your file:")
			scanner.Scan()
			filePath := scanner.Text()
			readFileData(filePath)
		case 3:
			return
		default:
			fmt.Println("You input is incorrect!")
		}
	}
}

func readFileData(fileName string) []string {

	var lines []string

	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("\nData in file %s:\n", fileName)
	for i, val := range lines {
		fmt.Printf("%11d : %s\n", i+1, val)
	}

	checkError(scanner.Err())
	return lines
}

func addFileData(filename string, userLine string) {

	//scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Write your value to file:")
	//userText := scanner.Text()

	options := os.O_WRONLY | os.O_CREATE | os.O_APPEND

	file, err := os.OpenFile(filename, options, os.FileMode(0600))
	checkError(err)
	file.Write([]byte(userLine + "\n"))
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func showMenu() {
	fmt.Printf("\n%27s\n\n%30s\n%28s\n%18s\n\n", "<<<Main menu>>>", "1 - Add data to file", "2 - Show file data", "3 - Exit")
}
