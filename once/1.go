package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	max := rand.Intn(100)
	var once sync.Once
	for i := 0; i < max; i++ {
		once.Do(func() {
			fmt.Println(i)
		})
	}
}
