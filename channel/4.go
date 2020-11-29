package main

import "fmt"

func main() {
	dataChan := make(chan string, 5)
	syncChan := make(chan struct{}, 2)
	data1 := []string{"a", "b", "c", "d", "e"}
	go receive(dataChan, syncChan)
	go send(data1, dataChan, syncChan)
	<-syncChan
	<-syncChan
}

func receive(dataChan <-chan string, syncChan chan struct{}) {
	for ele := range dataChan {
		fmt.Printf("接受数据:%s\n", ele)
	}
	fmt.Println("接受完成")
	syncChan <- struct{}{}
}

func send(data []string, dataChan chan<- string, syncChan chan struct{}) {
	for _, ele := range data {
		dataChan <- ele
	}
	fmt.Println("发送数据完成")
	syncChan <- struct{}{}
	close(dataChan)
}
