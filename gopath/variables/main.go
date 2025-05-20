package main

import (
	"fmt"
	"unsafe"
)

func main() {

	/*


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
	*/

	/* const MyIntConst int = 12
	fmt.Println("my constant is:  ", MyIntConst)

	const myStringConst = "a12"
	fmt.Println("my constant string is: ", myStringConst)

	myOtherScopeVariable := 50

	{
		myScopeVariable := 40
		fmt.Println("My variable: ", myScopeVariable)
		fmt.Println("My variable: ", myOtherScopeVariable)
	}

	myScopeVariable := 40
	fmt.Println("My variable: ", myScopeVariable)
	fmt.Println("My variable: ", myOtherScopeVariable)

	*/

	var my8BitsUintVar uint8
	fmt.Printf("type: %T, value: %d ,bytes: %d, bits: %d \n", my8BitsUintVar, my8BitsUintVar, unsafe.Sizeof(my8BitsUintVar), unsafe.Sizeof(my8BitsUintVar)*8)

	var my64BitsUintVar uint64 = 90
	fmt.Printf("type: %T, value: %d ,bytes: %d, bits: %d \n", my64BitsUintVar, my64BitsUintVar, unsafe.Sizeof(my64BitsUintVar), unsafe.Sizeof(my64BitsUintVar)*8)

	var myFloat32Var float32
	fmt.Printf("type: %T, value: %f ,bytes: %d, bits: %d \n", myFloat32Var, myFloat32Var, unsafe.Sizeof(myFloat32Var), unsafe.Sizeof(myFloat32Var)*8)

	var myStringVar3 string = "test1234"
	fmt.Printf("type: %T, value: %s ,bytes: %d, bits: %d \n", myStringVar3, myStringVar3, unsafe.Sizeof(myStringVar3), unsafe.Sizeof(myStringVar3)*8)

}
