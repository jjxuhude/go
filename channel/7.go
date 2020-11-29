package main

import "fmt"

var intChan = make(chan int, 10)

func main() {

	for i := 0; i < 5; i++ {
		select {
		case intChan <- 1:
		case intChan <- 2:
		case intChan <- 3:
		case intChan <- 4:
		case intChan <- 5:
		}
	}
	close(intChan)

	for ele := range intChan {
		fmt.Println(ele)
	}

	//for i:=0;i<5;i++{
	//	fmt.Println(<-intChan)
	//}
}
