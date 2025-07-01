package main

import (
	"fmt"
	"fundamentos-go/functions/function"
)

func main() {
	function.Display(10)

	fmt.Println()

	v := function.Add(5, 10)
	fmt.Println("Sum:", v)

	function.RepeatString(3, "Hello, World!")

	value, err := function.Calc(function.SUM, 10, 5)

	fmt.Println("Value:", value, "- Error:", err)
}
