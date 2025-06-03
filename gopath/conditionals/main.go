package main

import (
	"fmt"
)

func main() {

	/*

		yearsOld := 33

		if yearsOld > 18 {
			fmt.Printf("You are an adult, you are %d years old\n", yearsOld)
		} else {
			fmt.Printf("You are a minor, you are %d years old\n", yearsOld)
		}

		if value := true; value {
			fmt.Println("This is true")
		} else {
			fmt.Println("This is false")
		}
	*/

	number := 5

	switch number {
	case 1:
		fmt.Println("Number is one")
	case 2:
		fmt.Println("Number is two")
	case 3:
		fmt.Println("Number is three")
	default:
		fmt.Println("Number is something else")

	}

}
