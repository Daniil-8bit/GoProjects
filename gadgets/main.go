package gadgets

import (
	"fmt"

	"github.com/Daniil-8bit/GoProjects/gadgets"
)

func main() {
	list := []string{"Song 1", "Song 2", "Song 3"}
	device := gadgets.TapePlayer
	playSongs(device, list)
}

func playSongs(gadget gadgets.TapePlayer, songs []string) {
	for _, elem := range songs {
		fmt.Println(gadget.Play(elem))
	}
	gadget.Stop()
}
