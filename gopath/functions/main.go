package main

import (
	"fmt"
	"fundamentos-go/functions/function"
)

func main() {
	/*

		function.Display(10)

		fmt.Println()

		v := function.Add(5, 10)
		fmt.Println("Sum:", v)

		function.RepeatString(3, "Hello, World!")

		value, err := function.Calc(function.SUM, 10, 5)

		fmt.Println("Value:", value, "- Error:", err)

		xVal, yVal := function.Split(100)
		fmt.Println("Split values:", xVal, yVal)

		val2 := function.MSum(1, 2, 3, 4, 5, 3, 3, 2, 3)
		fmt.Println("Sum of variable arguments:", val2)
	*/
	val3, err := function.MOperations(function.SUM, 1, 2, 3, 4, 5)
	if err != nil {
		fmt.Println("Error in MOperations:", err)
	} else {
		fmt.Println("Result of MOperations:", val3)
	}
}
