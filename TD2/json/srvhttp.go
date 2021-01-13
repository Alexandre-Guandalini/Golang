package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
    "os"
)

type User struct {
	ID string `json:"userID"`
	Login string `json:"userName"`
	Password string
}

var usersMap = make(map[string]User)

func main() {
	jsonFile, err := os.Open("users.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println("Successfully Opened users.json")
    // defer the closing of our jsonFile so that we can parse it later on
    defer jsonFile.Close()

    // read our opened xmlFile as a byte array.
    byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("error:", err)
	}
	
	
    // we initialize our Users array
    var users []User

    // we unmarshal our byteArray which contains our
    // jsonFile's content into 'users' which we defined above
	err = json.Unmarshal(byteValue, &users)
	if err != nil {
		fmt.Println("error:", err)
	}
	//fmt.Printf("%+v", users)
	//boucle line 37 with ID
	for _, user := range users {
		usersMap[user.ID] = user
	}
	fmt.Printf("%+v", users)
}