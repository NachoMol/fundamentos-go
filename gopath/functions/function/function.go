package function

import "fmt"

func Display(myValue int) {
	display(myValue)
}

func display(myValue int) {
	fmt.Println()
	fmt.Printf("Value: %d\n", myValue)
	fmt.Println()
}

func Add(x, y int) int {
	return x + y
}

func RepeatString(increment int, value string) {
	for i := 0; i < increment; i++ {
		fmt.Println(value)
	}
}
