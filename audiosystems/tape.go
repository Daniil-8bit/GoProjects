package audiosystems

import "fmt"

type TapeRecorder struct {
	Microphones int
}

func (tr TapeRecorder) Play(song string) {
	fmt.Println("Playing ", song)
}

func (tr TapeRecorder) Stop() {
	fmt.Println("Music stopped!")
}

func (tr TapeRecorder) Record() {
	fmt.Println("Recording!")
}

type TapePlayer struct {
	Batteries string
}

func (tp TapePlayer) Play(song string) {
	fmt.Println("Playing ", song)
}

func (tp TapePlayer) Stop() {
	fmt.Println("Music stopped!")
}

type PlayDevice interface {
	Play(song string)
	Stop()
}
