package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var rwm sync.RWMutex
	for i := 0; i < 3; i++ {
		go func(i int) {
			rwm.RLock()
			time.Sleep(time.Second * 2)
			fmt.Println(i)
			rwm.RUnlock()
		}(i)
	}

	time.Sleep(time.Microsecond * 500)
	fmt.Println("try to lock fro waiting...")
	rwm.Lock()
	fmt.Println("lock for waiting.")

}
