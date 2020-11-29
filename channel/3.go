package main

import "fmt"

func main() {
	dataChan := make(chan int, 5)
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)

	go func() {
		<-syncChan1
		for {
			if ele, ok := <-dataChan; ok {
				fmt.Println(ele)
			} else {
				break
			}
		}
		fmt.Println("接受完毕")
		syncChan2 <- struct{}{}
	}()

	go func() {
		for i := 0; i < 5; i++ {
			dataChan <- i
			fmt.Printf("发送数据%d", i)
		}
		close(dataChan)
		syncChan1 <- struct{}{}
		syncChan2 <- struct{}{}
		fmt.Printf("完成")
	}()

	<-syncChan2
	<-syncChan2

}
