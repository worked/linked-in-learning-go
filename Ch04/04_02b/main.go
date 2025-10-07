package main

import (
	"fmt"
	"time"
)

func main() {

	weekday := time.Now().Weekday()
	fmt.Printf("Today is %v\n", weekday)

	dayNumber := int(weekday)
	fmt.Printf("The day as a number is %v\n", dayNumber)

	var result string
	switch dayNumber {
	case 1:
		result = "Monday"
	case 2:
		result = "Tuesday"
	case 3:
		result = "Wednesday"
	default:
		result = "Weekend"
	}
	fmt.Println(result)
}
