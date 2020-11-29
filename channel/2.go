package main

import (
	"fmt"
	"time"
)

type Counter struct {
	count int
}

var mapChan1 = make(chan map[string]*Counter, 1)

func main() {
	syncChan := make(chan struct{}, 2)
	go func() {
		for {
			if ele, ok := <-mapChan1; ok {
				counter := ele["count"]
				counter.count++
			} else {
				break
			}
		}
		fmt.Println("停止接受")
		syncChan <- struct{}{}
	}()

	go func() {
		countMap := map[string]*Counter{
			"count": &Counter{},
		}
		for i := 0; i < 5; i++ {
			mapChan1 <- countMap
			time.Sleep(time.Microsecond)
			fmt.Printf("第%d个值是%v\n", i, countMap)
		}
		close(mapChan1)
		syncChan <- struct{}{}
	}()

	<-syncChan
	<-syncChan

}
func (counter *Counter) String() string {
	return fmt.Sprintf("{count %d}", counter.count)
}
