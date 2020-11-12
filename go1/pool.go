package main

import (
	"fmt"
	"sync"
)

func main() {
	var pool = &sync.Pool{
		New: func() interface{} {
			return "Hello,World!"
		},
	}
	value := "Hello,学院君!"
	pool.Put(value)
	fmt.Println(pool.Get())
	fmt.Println(pool.Get())
}
