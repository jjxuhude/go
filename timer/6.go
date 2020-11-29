package main

import (
	"fmt"
	"time"
)

func main() {
	tricker := time.NewTicker(time.Second)
	syncChan := make(chan bool, 2)
	dataChan := make(chan string, 10)
	stop := make(chan bool, 1)
	go func() {
	loop:
		for {
			select {
			case <-tricker.C:
				dataChan <- time.Now().String()
			case <-stop:
				tricker.Stop()
				break loop
			}
		}
		fmt.Println("写入完成")
		syncChan <- true

	}()

	go func() {
		var i int
		for {
			if ele, ok := <-dataChan; ok {
				i++
				fmt.Println(ele)
				if i >= 3 {
					stop <- true
					break
				}
			}
		}
		fmt.Println("读取完成")
		syncChan <- true
	}()
	<-syncChan
	<-syncChan

}
