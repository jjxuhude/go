package main

import (
	"fmt"
	"runtime"
)

func ProtectRun(entry func()) {
	defer func() {
		err := recover()
		switch detail := err.(type) {
		case runtime.Error:
			fmt.Println("运行时错误", detail)
		default:
			fmt.Println("其他错误")
		}
	}()
	entry()
}

func main() {
	fmt.Println("运行前")
	ProtectRun(func() {
		fmt.Println("手工宕机前")
		panic("xxxxx")
		fmt.Println("手工宕机后")
	})

	ProtectRun(func() {
		fmt.Println("赋值宕机前")

		var a *int
		*a = 2

		fmt.Println("赋值宕机后")
	})
	fmt.Println("运行后")
}
