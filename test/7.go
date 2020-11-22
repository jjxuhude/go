package main

import "fmt"

var data3 = []string{"user01", "user02"}

func main() {
	for _, v := range data3 {
		*&v = "3333"
	}
	fmt.Println(data3)

}
