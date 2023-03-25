package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	fmt.Println("testing")
	readJsonFile()
}

func readJsonFile() {
	// Open our jsonFile
	jsonFile, err := os.Open("contents.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("successfully opened json file")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Users
	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type:" + users.Users[i].Type)
		fmt.Println("User Age:" + strconv.Itoa((users.Users[i].Age)))
		fmt.Println("User Name")
	}
}
