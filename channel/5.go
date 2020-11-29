package main

import "fmt"

func main() {
	var strChan = make(chan string, 10)
	var intChan = make(chan int, 10)

	select {
	case e1 := <-strChan:
		fmt.Println("1 case selected", e1)
	case e2 := <-intChan:
		fmt.Println("2th case selected", e2)
	default:
		fmt.Println("default case selected")
	}
}
