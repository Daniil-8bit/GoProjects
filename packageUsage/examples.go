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
)

func main() {

	//checkAudiosystems()
	//checkInterfaces()
	//checkErrors()
	//checkRefrigerator()
	//checkDir("C:\\Go\\go1.23.4\\src\\github.com\\Daniil-8bit\\GoProjects\\testDir")
	//checkDirRecursion("C:\\Go\\go1.23.4\\src\\github.com\\Daniil-8bit\\GoProjects\\testDir")
	defer stopPanic()
	//panic("some other panic!")
	checkDirRecursionPanic("C:\\")
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
