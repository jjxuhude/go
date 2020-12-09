package main

import "fmt"

func main() {
	str := "111"
	f := func() {
		str = "2222"
	}
	f()
	fmt.Println(str)
}
