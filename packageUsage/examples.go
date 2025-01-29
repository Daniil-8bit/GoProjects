package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
	"unicode/utf8"

	"github.com/Daniil-8bit/GoProjects/abstracts"
	"github.com/Daniil-8bit/GoProjects/audiosystems"
	"github.com/Daniil-8bit/GoProjects/hotel"
	"github.com/Daniil-8bit/GoProjects/readFile"
	"github.com/Daniil-8bit/GoProjects/refrigerator"
	"github.com/Daniil-8bit/GoProjects/sentences"
	"github.com/Daniil-8bit/GoProjects/webpage"
	"github.com/Daniil-8bit/GoProjects/webserver"
)

func main() {

	//checkAudiosystems()
	//checkInterfaces()
	//checkErrors()
	//checkRefrigerator()
	//checkDir("C:\\Go\\go1.23.4\\src\\github.com\\Daniil-8bit\\GoProjects\\testDir")
	//checkDirRecursion("C:\\Go\\go1.23.4\\src\\github.com\\Daniil-8bit\\GoProjects\\testDir")
	//defer stopPanic()
	//panic("some other panic!")
	//checkDirRecursionPanic("C:\\")

	//checkWebPage()\
	//checkWebPageCycle()
	//checkWebPageCycleStruct()
	//checkConcat()
	//checkWebServer()
	//checkFirstClassFunc(tryFunc2, tryFunc4)
	//startServer()
	//checkHotelWeb()
	//checkTemplates()
	//checkWorkWithFile()
	//filework.Work()
	//checkIfInit()

	channel := make(chan string, 3)
	fmt.Println(time.Now())
	go showGoroutineLetters(channel)
	time.Sleep(time.Second * 5)
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(<-channel, time.Now())
	fmt.Println(time.Now())
}

func TryInterface(player audiosystems.PlayDevice) {
	player.Play("*Music*")
	player.Stop()
	recordPlayer, ok := player.(audiosystems.TapeRecorder)
	if ok {
		recordPlayer.Record()
	} else {
		fmt.Println("Not ok!")
	}
}

func checkAudiosystems() {
	var newDevice audiosystems.PlayDevice = audiosystems.TapePlayer{Batteries: "line"}
	newDevice.Play("Song")
	newDevice.Stop()

	var device audiosystems.TapeRecorder
	TryInterface(device)

	//var recorderDevice audiosystems.PlayDevice = audiosystems.TapeRecorder{Microphones: 3}
	var tr audiosystems.TapeRecorder //= recorderDevice.(audiosystems.TapeRecorder)

	TryInterface(tr)

	var tp audiosystems.TapePlayer
	TryInterface(tp)
}

func checkInterfaces() {
	//var line abstracts.Example
	var err error
	err = abstracts.Example("Some line")
	fmt.Println(err)

	var err1 error = abstracts.HighTemp(78.123)
	fmt.Println(err1)

	car1 := abstracts.Car{Brand: "Audi", Model: "A7", Year: 2017}
	fmt.Println(car1)
}

func checkErrors() {
	//numbersFloat, err := readFile.GetFloats("C:\\Go\\go1.23.4\\src\\github.com\\Daniil-8bit\\GoProjects\\readFile\\text.txt")
	numbersFloat, err := readFile.GetFloats(os.Args[1])
	if err != nil {
		//fmt.Println(err.Error())
		log.Fatal(err)
	}
	fmt.Println(numbersFloat)
	var sum float64
	for _, val := range numbersFloat {
		sum += val
	}
	fmt.Println("Sum: ", sum)
}

func checkRefrigerator() {
	var newRef refrigerator.Refrigerator = refrigerator.Refrigerator{"apple", "cheese", "milk", "butter", "bread"}
	newRef.FindProduct("milk")
	newRef.FindProduct("borsch")
}

func checkDir(dirName string) error {
	elements, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, value := range elements {
		if value.IsDir() {
			fmt.Println("Directory: ", value.Name())
		} else {
			fmt.Println("File: ", value.Name())
		}
	}
	return nil
}

func checkDirRecursion(dirName string) error {
	fmt.Println("Current dir: ", dirName)
	elements, err := ioutil.ReadDir(dirName)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	for _, value := range elements {
		fullPath := filepath.Join(dirName, value.Name())
		if value.IsDir() {
			fmt.Println("Directory: ", value.Name())
			err = checkDirRecursion(fullPath)
			if err != nil {
				return err
			}
		} else {
			fmt.Println("File: ", value.Name())
		}
	}
	return nil
}

func checkDirRecursionPanic(dirName string) {
	fmt.Println("Current dir: ", dirName)
	elements, err := ioutil.ReadDir(dirName)
	if err != nil {
		panic(err)
	}
	for _, value := range elements {
		fullPath := filepath.Join(dirName, value.Name())
		if value.IsDir() {
			checkDirRecursionPanic(fullPath)
		} else {
			fmt.Println("File: ", value.Name())
		}
	}
}

func stopPanic() {
	res := recover()
	if res == nil {
		return
	}
	err, ok := res.(error)
	if ok {
		fmt.Println(err)
	} else {
		panic(res)
	}
}

func checkWebPage() {

	sizeChan := make(chan int)

	go webpage.ReadWebPage("https://bpmsoft.ru", sizeChan)
	go webpage.ReadWebPage("https://vk.com", sizeChan)
	go webpage.ReadWebPage("https://www.google.com", sizeChan)

	fmt.Println(<-sizeChan)
	fmt.Println(<-sizeChan)
	fmt.Println(<-sizeChan)
}

func checkWebPageCycle() {

	sizeChan := make(chan int)
	urls := []string{"https://bpmsoft.ru", "https://vk.com", "https://www.google.com"}

	for _, url := range urls {
		go webpage.ReadWebPage(url, sizeChan)
	}

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizeChan)
	}
}

func checkWebPageCycleStruct() {

	sizeChan := make(chan webpage.WebPage)
	urls := []string{"https://bpmsoft.ru", "https://vk.com", "https://www.google.com"}

	for _, url := range urls {
		go webpage.ReadWebPageStruct(url, sizeChan)
	}

	for i := 0; i < len(urls); i++ {
		page := <-sizeChan
		fmt.Printf("Web page: %s, size: %d\n", page.Url, page.Size)
	}
}

func checkConcat() {
	sentence := []string{"first", "program", "it's ok!"}
	finalLine := sentences.JoinSomeWords(sentence)
	fmt.Println(finalLine)
	sentence = []string{"second", "it's ok!"}
	finalLine = sentences.JoinSomeWords(sentence)
	fmt.Println(finalLine)
	sentence = []string{"one", "two", "three", "four", "it's ok!"}
	finalLine = sentences.JoinSomeWords(sentence)
	fmt.Println(finalLine)
	sentence = []string{"one"}
	finalLine = "Number " + sentences.JoinSomeWords(sentence)
	fmt.Println(finalLine)
}

func checkWebServer() {
	webserver.ActivateServer()
}

func checkFirstClassFunc(tf func(), tf1 func(int) string) {
	var a func() = tryFunc1
	a()
	tf()
	var b func(int, float64) string = tryFunc3
	fmt.Println(b(13, 3.14))
	fmt.Println(tf1(23))
}

func tryFunc1() {
	fmt.Println("This is the function1 in variable!")
}

func tryFunc2() {
	fmt.Println("This is the function2 in valiable!")
}

func tryFunc3(num int, num1 float64) string {
	return fmt.Sprintf("Line of values: %d, %.2f\n", num, num1)
}

func tryFunc4(num int) string {
	return fmt.Sprintf("Line of values: %d\n", num)
}

/*func startServer() {
	http.HandleFunc("/", pageHandler)
	http.ListenAndServe(":1313", nil)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Some information"))
}*/

func checkHotelWeb() {
	hotel.StartServer()
}

func checkTemplates() {
	text := "Start of template\nAction: {{.}}\nEnd of template\n"
	tmpl, err := template.New("test").Parse(text)
	if err != nil {
		log.Fatal(err)
	}
	err = tmpl.Execute(os.Stdout, "Value 1")
	checkError(err)
	err = tmpl.Execute(os.Stdout, 14)
	checkError(err)
	err = tmpl.Execute(os.Stdout, true)
	checkError(err)

	executeTemplate("start {{if .}} Dot is true! {{end}} finish\n", true)
	executeTemplate("start {{if .}} Dot is true! {{end}} finish\n", false)

	loopTemplate := "Before loop: {{.}}\n{{range .}}In loop: {{.}} \n{{end}}After loop: {{.}}\n"
	executeTemplate(loopTemplate, []string{"apple", "peach", "pineapple"})
	executeTemplate(loopTemplate, []int{13, 23, 77})

	structTemplate := "Elements:\n{{range .}}name: {{.Name}} and amount: {{.Amount}}\n{{end}}\n"
	elems := []Elem{
		{Name: "One", Amount: 300},
		{Name: "Two", Amount: 500},
		{Name: "Three", Amount: 1500},
	}
	executeTemplate(structTemplate, elems)
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func executeTemplate(text string, data interface{}) {
	tmpl, err := template.New("test").Parse(text)
	checkError(err)
	err = tmpl.Execute(os.Stdout, data)
	checkError(err)
}

type Elem struct {
	Name   string
	Amount int
}

func checkWorkWithFile() {
	// check read only file
	readOnlyFile, err := os.OpenFile("testFile.txt", os.O_RDONLY, os.FileMode(0600))
	checkError(err)

	scanner := bufio.NewScanner(readOnlyFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// try to write in file: access is denied
	/*_, err = readOnlyFile.Write([]byte("Try to write text"))
	checkError(err)*/

	err = readOnlyFile.Close()
	checkError(err)
	checkError(scanner.Err())

	//check write only file (Doesn't add new info, rewrite text to file!)
	writeOnlyFile, err := os.OpenFile("testFile.txt", os.O_WRONLY, os.FileMode(0600))
	checkError(err)

	//try to read file: access is denied
	/*scanner = bufio.NewScanner(writeOnlyFile)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}*/

	_, err = writeOnlyFile.Write([]byte("Some new text!\n"))
	checkError(err)
	err = writeOnlyFile.Close()
	checkError(err)
	checkError(scanner.Err())

	//check append text to file
	options := os.O_APPEND | os.O_CREATE

	appendFile, err := os.OpenFile("testCreateFile.txt", options, os.FileMode(0600))
	checkError(err)
	_, err = appendFile.Write([]byte("Appended text!\n"))
	checkError(err)
	err = appendFile.Close()
	checkError(err)
}

func checkIfInit() {
	if err := testIfFunc("one"); err != nil {
		fmt.Println("err != nil;", err)
	}

	if err := testIfFunc("TOO BIG LETTERS!"); err != nil {
		fmt.Println("err != nil!", err)
	}

	switch rand.Intn(3) + 1 {
	case 1:
		fmt.Println("Gold!")
	case 2:
		fmt.Println("Silver!")
	case 3:
		fmt.Println("Bronze!")
	default:
		fmt.Println("Wood...")
	}

	var num8 int8 = 8
	var num16 int16 = 16
	var plusNum uint = 123
	var plusNum8 uint8 = 077
	var numFloat32 float32 = 345.123
	fmt.Println(num8, num16, plusNum, plusNum8, numFloat32)

	line1 := "ABCDE"
	line2 := "АБВГД"

	fmt.Println("Length eng in bytes:", len(line1))
	fmt.Println("Length rus in bytes:", len(line2))

	fmt.Println("Length eng in symbols: ", utf8.RuneCountInString(line1))
	fmt.Println("Length rus in symbols: ", utf8.RuneCountInString(line2))

	line1Bytes := []byte(line1)
	partLine1 := line1Bytes[3:]
	fmt.Println(string(partLine1))
	line2Bytes := []byte(line2)
	partLine2 := line2Bytes[3:]
	fmt.Println(string(partLine2))

	for i, val := range line1 {
		fmt.Println(i, ":", string(val))
	}

	for i, val := range line2 {
		fmt.Println(i, ":", string(val))
	}
}

func testIfFunc(line string) error {
	return errors.New("Error: " + line)
}

func showGoroutineLetters(channel chan string) {
	channel <- "a"
	time.Sleep(time.Second)
	channel <- "b"
	time.Sleep(time.Second)
	channel <- "c"
	time.Sleep(time.Second)
}
