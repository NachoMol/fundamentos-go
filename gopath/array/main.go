package main

import (
	"fmt"
	"unsafe"
)

func main() {

	var myIntVar int = 30
	fmt.Printf("type: %T, value: %d, size: %d, bytes %d, bits: %d \n", myIntVar, myIntVar, unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar), unsafe.Sizeof(myIntVar)*8)

	var myArrayVar1 [5]int
	fmt.Println(myArrayVar1)

	fmt.Println(myArrayVar1, "Size: ", len(myArrayVar1))

	myArrayVar2 := [3]string{"Hello", "World", "!"}
	fmt.Println(myArrayVar2, "Size: ", len(myArrayVar2))

	myArrayVar1[0] = 2
	myArrayVar1[1] = 5
	myArrayVar1[2] = 9
	fmt.Println(myArrayVar1, "Size: ", len(myArrayVar1))

	fmt.Printf("type: %T, value: %v, bytes %d, bits: %d \n", myArrayVar2, myArrayVar2, unsafe.Sizeof(myArrayVar2), unsafe.Sizeof(myArrayVar2)*8)

	for i, v := range myArrayVar1 {
		fmt.Printf("Index: %d, Value: %d\n", i, v)
	}

}
