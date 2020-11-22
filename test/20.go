package main

import (
	"fmt"
	"runtime"
	"time"
)

func runing() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)

	}
}
func main() {
	//go runing()
	//var input string
	//fmt.Scanln(&input)
	fmt.Println(runtime.NumCPU())
}
