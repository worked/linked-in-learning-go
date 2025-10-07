package main

import (
	"fmt"
)

func main() {
	dog := Dog{"Poodle", "Woof"}
	dog.Speak()
	Speak(dog)
}

type Dog struct {
	Breed string
	Sound string
}

//method with d reciever
func (d Dog) Speak() {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}

//function
func Speak(d Dog) {
	fmt.Printf("The %v says %v!\n", d.Breed, d.Sound)
}
