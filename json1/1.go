package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}
type Result struct {
	Code    int
	Message string
	Data    []User
}

func main() {

	type User struct {
		Name string
		Age  int
	}
	type ColorGroup struct {
		ID     int
		Name   string
		Colors []string
		Users  []User
	}
	group := &ColorGroup{
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon", "粮食"},
		Name:   "Reds",
		Users:  []User{{"user01", 11}, {"user02", 22}},
	}
	group.Users = append(group.Users, User{"user04", 444})
	jsonData, _ := json.Marshal(group)
	fmt.Println(string(jsonData))

	groupUsers := struct {
		Users []User
	}{}
	json.Unmarshal(jsonData, &groupUsers)
	fmt.Printf("%v", groupUsers)

}
