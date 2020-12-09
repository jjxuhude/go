package main

import (
	"fmt"
	"sort"
)

func main() {
	srcSlice := []int{1, 2, 3}
	var tarSlice = make([]int, 3)
	copy(tarSlice, srcSlice)
	fmt.Println(tarSlice)

	users := map[string]string{
		"user03": "kris",
		"user02": "kendo",
		"user01": "jack.xu1",
	}
	//delete(users,"user01")
	var userList []string
	for k := range users {
		userList = append(userList, k)
	}
	sort.Strings(userList)
	fmt.Println(userList)
}
