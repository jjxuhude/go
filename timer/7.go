package main

import (
	"fmt"
	"time"
)

func main() {
	tricker := time.NewTicker(time.Second)
	syncChan := make(chan struct{}, 2)
	dataChan := make(chan string, 10)
	stop := make(chan struct{}, 1)
	go func() {
		defer close(dataChan)
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
		syncChan <- struct{}{}

	}()

	go func() {
		var i int
		for ele := range dataChan {
			i++
			fmt.Println(ele)
			if i >= 3 {
				stop <- struct{}{}
				break
			}
		}
		fmt.Println("读取完成")
		syncChan <- struct{}{}
	}()
	<-syncChan
	<-syncChan

}
