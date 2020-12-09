package main

import (
	"fmt"
	"sync"
)

func main() {
	var userMap sync.Map
	userMap.Store("user01", "jack")
	userMap.Store("user02", "rose")
	userMap.Store("user03", "kendo")

	userMap.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})

}
