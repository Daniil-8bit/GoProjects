package greeting

import "fmt"

func EngGreeting() {
	fmt.Println("Hello, everyone!")
}

func EspGreeting() {
	fmt.Println("Hola! Buenas dias!")
}

func callOtherFunctions() {
	EngGreeting()
	EspGreeting()
}
