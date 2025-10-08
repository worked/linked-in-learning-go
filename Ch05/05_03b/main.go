package main

import (
	"fmt"
	"time"
)

const (
	oneSecond = 1 * time.Second
	twoSecond = 2 * time.Second
)

func main() {
	go say("hello from goroutine")
	fmt.Println("hello from main")

	go func(message string) {
		fmt.Println(message)
	}("Hello from anon func")

	time.Sleep(twoSecond)
	fmt.Println("all done")
}

func say(message string) {
	time.Sleep(oneSecond)
	fmt.Println(message)
}
