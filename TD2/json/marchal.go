package main

import (
	"fmt"
	"encoding/json"
)

type User struct {
	Login string `json:"userName"`
	Password string
}

func main() {
	data := User{
		Login:		"Paul", 
		Password:	"pass123",
	}
	id, err := json.Marshal(data)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(id))
}