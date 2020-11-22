package main

import "fmt"

func protectRun(proc func()) {

	defer func() {
		err := recover()
		fmt.Println(err)
	}()
	proc()
}

func main() {

	protectRun(func() {
		var aaa *int
		*aaa = 111
		fmt.Println(aaa)
	})
	fmt.Println("完成")
}
