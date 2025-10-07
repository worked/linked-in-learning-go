package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"Poodle", 34}
	fmt.Println(poodle)
	fmt.Printf("%+v\n", poodle)
	fmt.Printf("Breed: %v, Weight: %v\n", poodle.Breed, poodle.Weight)
	poodle.Weight = 29
	fmt.Println(poodle.Weight)
}

type Dog struct {
	Breed  string
	Weight int
}
