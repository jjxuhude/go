package main

import (
	"fmt"
	"time"
)

var mapChan = make(chan map[string]int, 1)

func main() {
	var syncChan = make(chan struct{}, 2)

	go func() {
		for {
			if ele, ok := <-mapChan; ok {
				ele["count"]++
			} else {
				break
			}
		}
		fmt.Println("stop receive.")
		syncChan <- struct{}{}
	}()

	go func() {
		mapCount := make(map[string]int)
		for i := 0; i < 5; i++ {
			mapChan <- mapCount
			time.Sleep(time.Microsecond)
			fmt.Printf("第%v个数据：%v\n", i, mapCount)
		}
		close(mapChan)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan

}
