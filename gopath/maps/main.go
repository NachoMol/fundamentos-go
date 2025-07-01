package main

import (
	"fmt"
)

func main() {

	map1 := make(map[int]string)

	map1[1] = "A"
	map1[2] = "B"
	map1[3] = "C"
	map1[99] = "Z"
	map1[-5] = "Negative Five"

	fmt.Println("Map 1:", map1, "Length:", len(map1))
	fmt.Println("Value for key 1:", map1[1])

	map2 := make(map[int][]string)
	map2[1] = []string{"A", "B", "C"}
	map2[2] = []string{"D", "E", "F"}

	fmt.Println("Map 2:", map2, "Length:", len(map2))
	fmt.Println("Value for key 1:", map2[1])

	delete(map1, 99)
	fmt.Println("Map 1 after deletion:", map1, "Length:", len(map1))

	v, ok := map1[3]
	fmt.Println("Value for key 3:", v, "Exists:", ok)
	v, ok = map1[8]
	fmt.Println("Value for key 8:", v, "Exists:", ok)

	map4 := map[int]string{
		1: "One",
		2: "Two",
		3: "Three",
	}

	fmt.Println("Map 4:", map4, "Length:", len(map4))

	for key, value := range map4 {
		fmt.Printf("Key: %d, Value: %s\n", key, value)
	}

}
