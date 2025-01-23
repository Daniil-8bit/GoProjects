package main

import (
	"fmt"

	"github.com/Daniil-8bit/GoProjects/abstracts"
	"github.com/Daniil-8bit/GoProjects/audiosystems"
	"github.com/Daniil-8bit/GoProjects/greeting"
)

func main() {
	greeting.EngGreeting()
	greeting.EspGreeting()

	var newDevice audiosystems.PlayDevice

	newDevice = audiosystems.TapePlayer{Batteries: "line"}
	newDevice.Play("Song")
	newDevice.Stop()

	var device audiosystems.TapeRecorder
	TryInterface(device)

	//var recorderDevice audiosystems.PlayDevice = audiosystems.TapeRecorder{Microphones: 3}
	var tr audiosystems.TapeRecorder //= recorderDevice.(audiosystems.TapeRecorder)

	TryInterface(tr)

	var tp audiosystems.TapePlayer
	TryInterface(tp)

	//var line abstracts.Example
	var err error
	err = abstracts.Example("Some line")
	fmt.Println(err)

	var err1 error = abstracts.HighTemp(78.123)
	fmt.Println(err1)

	car1 := abstracts.Car{Brand: "Audi", Model: "A7", Year: 2017}
	fmt.Println(car1)
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
