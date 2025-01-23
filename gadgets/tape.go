package gadgets

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

func main() {
	list := []string{"Song 1", "Song 2", "Song 3"}
	device := TapePlayer{}
	playSongs(device, list)
}

func playSongs(gadget TapePlayer, songs []string) {
	for _, elem := range songs {
		gadget.Play(elem)
	}
	gadget.Stop()
}
