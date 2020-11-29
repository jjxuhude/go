package main

import (
	"fmt"
	"time"
)

var intChan = make(chan int, 10)

func main() {
	go func() {
		time.Sleep(time.Second)
		intChan <- 1
	}()

	select {
	case ele := <-intChan:
		fmt.Println(ele)
	case <-time.After(time.Microsecond * 500):
		//case <-time.NewTimer(time.Microsecond * 500).C:
		fmt.Println("超时")

	}
}
