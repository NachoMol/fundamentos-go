package main

import (
	"fmt"
)

func main() {

	var myIntVar int
	myIntVar = -12
	fmt.Println("My variable is: ", myIntVar)

	var myUintVar uint
	myUintVar = 12
	fmt.Println(" My variable is: ", myUintVar)

	var myStringVar string
	myStringVar = "my string var"
	fmt.Println(" My variable is: ", myStringVar)

	var myBoolVar bool
	myBoolVar = true
	fmt.Println(" My variable is: ", myBoolVar)

	fmt.Println("my variable address is: ", &myStringVar)

	myIntVar2 := 12
	fmt.Println("my IntVar2 is: ", myIntVar2)

	myStringVar2 := "my string 2"
	fmt.Println("my variable is: ", myStringVar2)
}
