package main

import (
	"container/list"
	"fmt"
)

func main() {
	userList1 := list.New()
	var userList2 list.List
	userList1.PushBack("user01")
	userList1.PushBack("user02")
	userList2.PushBack("user03")
	userList2.PushBack("user04")
	for i := userList1.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	for i := userList2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	for i := userList2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

}
