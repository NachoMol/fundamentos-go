package main

import (
	"fmt"
)

func main() {

	myArrayVar := [5]int{3, 6, 9, 10, 16}
	fmt.Println("Array:", myArrayVar, "Length:", len(myArrayVar))
	mySliceVar := []int{}

	mySliceVar = append((mySliceVar), 12, 34, 54)

	mySliceVar = append((mySliceVar), 12, 34, 54)
	fmt.Println("Slice:", mySliceVar, "Length:", len(mySliceVar))

	mySliceVar2 := myArrayVar[2:4]

	mySliceVar2[0] = 19

	fmt.Println("Slice 2:", mySliceVar2, "Length:", len(mySliceVar2), &mySliceVar2[0])
	fmt.Println("Array:", myArrayVar, "Length:", len(myArrayVar), &myArrayVar[2])

	mySliceVar3 := mySliceVar[2:]
	fmt.Println("Slice 3:", mySliceVar3, "Length:", len(mySliceVar3), &mySliceVar3[0])

	mySliceVar4 := make([]int, 3)
	fmt.Println("Slice 4:", mySliceVar4, "Length:", len(mySliceVar4), &mySliceVar4[0])

	mySliceVar5 := []int{1, 2, 6, 11, 20, 5, 1, 0}
	fmt.Println("Slice 5:", mySliceVar5, "Length:", len(mySliceVar5), &mySliceVar5[0])

}
