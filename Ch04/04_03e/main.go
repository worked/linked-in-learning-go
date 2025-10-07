package main

import (
	"fmt"
	"math"
)

func main() {
	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	for i := range colors {
		fmt.Println(colors[i])
	}

	for _, color := range colors {
		fmt.Println(color)
	}

	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum:", sum)
		if sum >= 200 {
			goto theEnd
		}
		//sum++
	}

theEnd:
	fmt.Println("End of program")

	var cart []CartItem
	var apples = CartItem{"apple", 2.99, 3}
	var oranges = CartItem{"orange", 1.50, 8}
	var bananas = CartItem{"banana", .49, 12}
	cart = append(cart, apples, oranges, bananas)
	result := calculateTotal(cart)
	fmt.Printf("Cart Total: $%v\n", result)


	sum, count, ave := addAllValues(123, 23, 1, 123, 33)
	fmt.Printf("sum of all values %v\n", sum)
	fmt.Printf("count of all values %v\n", count)
	fmt.Printf("average of all values %v\n", ave)
}

type CartItem struct {
	Name     string
	Price    float64
	Quantity int
}

// calculateTotal() returns the total value of the shopping cart.
func calculateTotal(cart []CartItem) float64 {
	cart_total := 0.00
	//loop over cart, get the price x quanity = cart_total
	for _, item := range cart {
		cart_total += item.Price * float64(item.Quantity) //concatenate cart_total
	}
	return roundFloat(cart_total)
}

func roundFloat(value float64) float64 {
	return math.Round(value*100) / 100
}

// overload function with the same type
// this also return multiple values of different types
func addAllValues(values ...int) (int, int, float64) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	count := len(values)
	ave := float64(sum) / float64(count)
	return sum, count, ave
}
