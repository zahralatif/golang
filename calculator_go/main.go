package main

import "fmt"

func main() {
	const a = 5
	const b = 10

	var sum = a + b
	var difference = b - a
	var product = a * b
	var division = float64(b) / float64(a)

	fmt.Println("Sum: ", sum)
	fmt.Println("Difference: ", difference)
	fmt.Println("Product: ", product)
	fmt.Println("Division: ", division)

}
