package main

import "fmt"

var intChan1 chan int
var intChan2 chan int
var channels = []chan int{intChan1, intChan2}
var number = []int{1, 2, 3, 4, 5}

func main() {

	select {
	case getChan(0) <- getNumber(0):
		fmt.Println("0 case selected")
	case getChan(1) <- getNumber(1):
		fmt.Println("2 case Selected")
	default:
		fmt.Println("defualt calse")
	}

}

func getNumber(i int) int {
	fmt.Printf("get number: %d \n", i)
	return number[i]
}
func getChan(i int) chan int {
	fmt.Printf("get chan: %d\n", i)
	return channels[i]
}
