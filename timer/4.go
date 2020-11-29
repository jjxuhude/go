package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)
	dataChan := make(chan string, 1)
	syncChan := make(chan bool, 2)

	go func() {
		var i int
		for ele := range dataChan {
			i++
			fmt.Println(ele)
			if i > 3 {
				syncChan <- true
				break
			}
		}
		syncChan <- true
		ticker.Stop()
		fmt.Println("接受结束")
	}()

	go func() {
		for range ticker.C {
			dataChan <- time.Now().String()
		}
		fmt.Println("发送结束")
	}()
	<-syncChan
	<-syncChan
	fmt.Println("End. [receiver]")
}
