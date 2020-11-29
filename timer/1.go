package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Printf("开始时间%v\n", time.Now())
	timer := time.NewTimer(10 * time.Second)
	dTime := <-timer.C
	fmt.Printf("结束时间%v\n", dTime)
	fmt.Printf("停止Timer%v\n", timer.Stop())
}
