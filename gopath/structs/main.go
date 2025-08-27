package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name,omitempty"`
	Address  string `json:"address,omitempty"`
	Age      int    `json:"age,omitempty"`
	LastName string `json:"last_name,omitempty"`
}

func main() {
	user := User{
		ID:   1,
		Name: "Nacho",
	}
	fmt.Println(user)
	v, err := json.Marshal(user)
	fmt.Println(err)
	fmt.Println(string(v))

}
