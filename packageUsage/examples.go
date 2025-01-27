package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/Daniil-8bit/GoProjects/abstracts"
	"github.com/Daniil-8bit/GoProjects/audiosystems"
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
	checkFirstClassFunc(tryFunc2, tryFunc4)
	//startServer()
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
