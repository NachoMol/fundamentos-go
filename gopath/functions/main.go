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
}
