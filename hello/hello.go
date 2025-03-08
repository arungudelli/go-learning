package main

import (
	"fmt"

	"github.com/arungudelli/go-learning/modules/calculator" // Change this to the actual import path
)

func main() {
	a, b := 10.0, 5.0

	fmt.Println("Addition:", calculator.Add(a, b))
	fmt.Println("Subtraction:", calculator.Subtract(a, b))
	fmt.Println("Multiplication:", calculator.Multiply(a, b))

	result, err := calculator.Divide(a, b)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Division:", result)
	}
}
